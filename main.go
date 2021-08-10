package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Order struct {
	ID              string   `json:"id,omitempty" bson:"_id,omitempty"`
	UserID          int      `json:"userId" bson:"userId"`
	StoreID         int      `json:"storeId" bson:"storeId"`
	Products        []string `json:"products" bson:"products"`
	FulfillmentDate string   `json:"fulfillmentDate" bson:"fulfillmentDate"`
}

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

var (
	// Obviously, this is just a test example. Do not do this in production.
	// In production, you would have the private key and public key pair generated
	// in advance. NEVER add a private key to any GitHub repo.
	privateKey *rsa.PrivateKey
)

// Database settings (database name and connection URI)
const dbName = "per-diem"
const mongoURI = "mongodb+srv://jay:jayjay123@per.wxzqv.mongodb.net/" + dbName + "?retryWrites=true&w=majority"

func main() {

	// Connect to the database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// generate a new private/public key pair on each run. See note above.
	rng := rand.Reader
	var err error
	privateKey, err = rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	app := fiber.New()

	// cross-origin middleware
	app.Use(cors.New())

	setupRoutes(app)

	app.Listen(":3000")
}

// Connect configures the MongoDB client and initializes the database connection.
func Connect() error {
	client, _ := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func setupRoutes(app *fiber.App) {

	app.Get("/jwt", generateJWT)

	//Fetch specific order
	app.Get("/order/*", func(c *fiber.Ctx) error {

		id := c.Params("*")

		if id != "" {
			return singleOrder(c, id)
		}

		return multipleOrders(c)

	})

	//Update order
	app.Put("/order/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		var err error

		orderId, err := primitive.ObjectIDFromHex(id)

		// the provided ID might be invalid ObjectID
		if err != nil {
			return c.SendStatus(400)
		}

		order := new(Order)

		// Parse body into struct
		if err := c.BodyParser(order); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		update := bson.M{
			"$set": bson.M{
				"userId":          order.UserID,
				"storeId":         order.StoreID,
				"products":        order.Products,
				"fulfillmentDate": order.FulfillmentDate,
			},
		}
		err = mg.Db.Collection("order").FindOneAndUpdate(c.Context(), bson.M{"_id": orderId}, update).Err()

		if err != nil {
			// ErrNoDocuments means that the filter did not match any documents in the collection
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(404)
			}
			return c.SendStatus(500)
		}

		// return the updated order
		order.ID = id
		return c.Status(200).JSON(order)

	})

	//Add shopping cart and generate orders
	app.Post("/cart", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    privateKey.Public(),
	}))
}

func singleOrder(c *fiber.Ctx, id string) error {

	result := Order{}
	orderId, err := primitive.ObjectIDFromHex(id)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.SendStatus(400)
	}

	// get all record as a cursor
	query := bson.D{{Key: "_id", Value: orderId}}
	err = mg.Db.Collection("order").FindOne(c.Context(), query).Decode(&result)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// return order in JSON format
	return c.JSON(result)
}

func multipleOrders(c *fiber.Ctx) error {

	// get all records as a cursor
	query := bson.D{{}}
	cursor, err := mg.Db.Collection("order").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var orders []Order = make([]Order, 0)

	// iterate the cursor and decode each item into an Order
	if err := cursor.All(c.Context(), &orders); err != nil {
		return c.Status(500).SendString(err.Error())

	}

	// return orders list in JSON format
	return c.JSON(orders)
}

func generateJWT(c *fiber.Ctx) error {

	user := "john"
	pass := "123"

	// Throws Unauthorized error
	if user != "john" || pass != "123" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodRS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString(privateKey)
	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})

}

func Calculate(x int) (result int) {
	result = x + 2
	return result
}
