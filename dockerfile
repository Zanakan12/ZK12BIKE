# Utilisation de l'image Go officielle comme base
FROM golang:1.23.3

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier le fichier go.mod et go.sum (pour tirer parti du cache Docker)
COPY go.mod go.sum ./

# Exécuter la commande go mod tidy pour résoudre les dépendances
RUN go mod tidy

# Copier le reste de votre code source dans l'image Docker
COPY . .

# Construire l'application Go
RUN go build -o /app/bin .

# Exposer le port 8080 pour l'application
EXPOSE 8080

# Lancer l'application Go
CMD ["/app/bin"]
