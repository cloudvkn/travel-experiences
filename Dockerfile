FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o travelapp cmd/travelapp/main.go

EXPOSE 8080

CMD ["./travelapp"]
