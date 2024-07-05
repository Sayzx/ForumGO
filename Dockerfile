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

# Copier l'exécutable
COPY --from=build /app/main .

# Assurez-vous que le fichier de la base de données est accessible et correct
COPY /internal/sql/forum.db /root/data/forum.db

RUN chmod 777 /root/data/forum.db

# Exposer le port sur lequel votre application est configurée pour écouter
EXPOSE 8080

CMD ["./main"]
