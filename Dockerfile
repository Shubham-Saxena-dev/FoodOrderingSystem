FROM golang:1.14.13-alpine3.12
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build
EXPOSE 8080
CMD ["./FoodOrderingSystem"]