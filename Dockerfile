# Étape de build
FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

# Étape finale
FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/main .
COPY --from=build /app/web ./web
COPY --from=build /app/internal/sql/forum.db /root/data/forum.db

EXPOSE 8080

CMD ["./main"]
