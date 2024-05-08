FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./app/cmd/server/main.go

EXPOSE 10080

CMD ["./server", "start"]