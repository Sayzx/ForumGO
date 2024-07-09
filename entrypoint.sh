#!/bin/sh

# Remplacer les variables d'environnement dans le fichier de configuration Nginx
envsubst '$SSL_CERT_PATH $SSL_KEY_PATH $PROXY_PASS_URL' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

# Vérifier la configuration de Nginx et démarrer les services
nginx -t && service nginx start && ./myapp
