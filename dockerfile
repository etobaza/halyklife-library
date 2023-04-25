FROM golang:1.20.3-bullseye

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
