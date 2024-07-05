# Utiliser l'image officielle de Golang 1.21.3 comme image de base pour la construction
FROM golang:1.21.3 AS builder

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Copier les fichiers go.mod et go.sum dans le répertoire de travail
COPY go.mod go.sum ./

# Télécharger les dépendances
RUN go mod download

# Copier le reste des fichiers du projet dans le répertoire de travail
COPY . .

# Compiler l'application
RUN go build -o myapp ./cmd

# Utiliser une image Ubuntu 22.04 pour l'exécution
FROM ubuntu:22.04

# Installer les dépendances nécessaires (sqlite3), Nginx et les certificats CA
RUN apt-get update && apt-get install -y \
    sqlite3 \
    nginx \
    ca-certificates \
    gettext-base \
    && rm -rf /var/lib/apt/lists/*

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Copier l'exécutable depuis l'étape de construction
COPY --from=builder /app/myapp .

# Copier le reste des fichiers du projet, y compris les templates et la base de données
COPY . .

# Copier les certificats dans l'image
COPY certs/localhost.crt /etc/nginx/certs/localhost.crt
COPY certs/localhost.key /etc/nginx/certs/localhost.key

# Copier la configuration Nginx template
COPY nginx/nginx.conf.template /etc/nginx/nginx.conf.template

# Copier le script d'entrée
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Définir les variables d'environnement pour Nginx
ENV SSL_CERT_PATH=/etc/nginx/certs/localhost.crt
ENV SSL_KEY_PATH=/etc/nginx/certs/localhost.key
ENV PROXY_PASS_URL=http://localhost:8080

# Exposer les ports utilisés par l'application et Nginx
EXPOSE 8080 80 443

# Utiliser le script d'entrée pour démarrer les services
CMD ["/entrypoint.sh"]
