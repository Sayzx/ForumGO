#!/bin/bash

# Démarrer l'application Go en arrière-plan
./myapp &

# Remplacer les variables d'environnement dans le fichier de configuration Nginx
envsubst '$SSL_CERT_PATH $SSL_KEY_PATH $PROXY_PASS_URL' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

# Démarrer Nginx
nginx -g 'daemon off;'
