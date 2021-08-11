## Per-Diem Backend Challenge

### Introduction

Per-Diem (Latin for "per day" or "for each day") - This is a is a substitute for using an actual expense reimbursement method.

The project in this repository is the backend api for the per-diem backend challenge. Which - as is related to its name - involves a scheduled way for customers to schedule orders they would like fulfilled over a period of time.

The features of the API satisy the given requirements, they include:
- 4 API endpoints.
- JWT Authentication.
- Dockerized.
- Unit Tests.
- Inline and external documentation of the repo (This readme file).

### Architecture

Here is a representation of the architecture.

![alt text](https://github.com/jaysonmulwa/per-diem/blob/main/arch.png?raw=true)

### Installation

1. Clone this repositry.
```bash
 https://github.com/jaysonmulwa/per-diem.git
```

2. Install docker to your machine.


3. Run the following commands to start up the instance.
```bash
$ docker build -t per .
```
```bash
$ docker run -p 3000:3000 per
```

4. Request for the jwt token from the endpoint.
```bash
http://localhost:3000/jwt
```

5. Add the token your request header for other requests.
```bash
Authorization : Bearer <token>
```

6. Access other endpoints. eg:
```bash
http://localhost:3000/order
```


### Endpoints

1. Get All Orders
```bash
GET http://localhost:3000/order
```



2. Get a single order
```bash
GET http://localhost:3000/order/{{orderID}}
```
Parameters: 
- orderID (int)



3. Update the order
```bash
PUT http://localhost:3000/order/{{orderID}}
```
Parameters: 
- orderID (int)

Body: 
- userId (int)
- storeId (int)
- products (Array)
- fulfillmentDate (Date)

Example Request Body:
```bash
{
    "userId": 121,
    "storeId": 120,
    "products": [
        "Eggs",
        "Milk",
        "Bread"
    ],
    "fulfillmentDate": ""
}
```


4. Create orders from the cart
```bash
POST http://localhost:3000/jwt
```
Body: 
- frequency (string) - weekly, bi-weekly, monthly
- userId (int)
- storeId (int)
- products (Array)
- scheduledDate (Array) - days of the week
- duration (int) - Number of times order is expected

Example Request Body:
```bash
{
    "frequency": "4",
    "userId": 44,
    "storeId": 44,
    "products": [
        "Eggs",
        "Milk",
        "Bread"
    ],
    "scheduledDate": [
    	"Tuesday"
    ],
    "duration": 2
}
```


### Tests
Run go test
```bash
$ go test
```

### Future Improvements
- Add a clean file structure


