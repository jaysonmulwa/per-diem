FROM golang:1.16.7 as development
# Add a work directory
WORKDIR /per

RUN go clean --modcache 
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .

# Expose port
EXPOSE 3000
# Start app
CMD go run main.go --start-service