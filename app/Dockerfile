FROM golang:1.19 AS builder

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

COPY --from=builder /app/main /main

EXPOSE 8080

CMD ["/main"]