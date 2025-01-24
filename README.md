
# 🚲 **BikeRentBuy**

Une plateforme en ligne pour louer et acheter des vélos.

## 📋 **Description**
Ce projet a pour objectif de créer une application web permettant :
- De louer des vélos pour une durée déterminée.
- D'acheter des vélos en ligne, neufs ou d'occasion.
- De suivre les commandes et réservations via un espace utilisateur.

## 🛠️ **Technologies utilisées**
- **Langage backend :** Go (Golang)
- **Framework web :** Not yet
- **Base de données :** SQLite
- **Versionnement :** Git/GitHub
- **Frontend :** HTML, CSS, JavaScript (intégration minimale côté Go avec templates)
- **Paiement :** Intégration avec Stripe ou PayPal

---

## 🏗️ **Fonctionnalités**
### **Utilisateurs :**
- Création de compte et connexion.
- Accès à un espace utilisateur (historique des commandes et réservations).

### **Catalogue de vélos :**
- Liste des vélos disponibles à louer ou acheter.
- Filtres par prix, type et disponibilité.

### **Gestion des commandes :**
- Réservation des vélos pour une durée spécifique.
- Ajout de vélos au panier et validation d'achat.

---

## 🛠️ **Installation et Lancement**
### 1. Clone du projet
```bash
git clone https://github.com/zanakan12/zk12ebike.git
cd zk12ebike
```

### 2. Configuration des dépendances
Installez les dépendances nécessaires :
```bash
go mod tidy
```

### 3. Configuration de la base de données
Ajoutez un fichier `.env` pour la configuration :
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=bikerentbuy
```

### 4. Exécutez les migrations
Ajoutez les migrations SQLITE. Exemple :
```bash
go run main.go migrate
```

### 5. Lancez le serveur
```bash
go run main.go
```

Accédez à l'application à l'adresse : [http://localhost:8080](http://localhost:8080)

---

## 📋 **Roadmap**
- [ ] Création des modèles pour la base de données.
- [ ] Développement des routes backend (CRUD des vélos, commandes).
- [ ] Développement de l'interface utilisateur avec des templates HTML.
- [ ] Intégration du paiement en ligne.
- [ ] Tests et déploiement.


## 🖼️ **Diagramme UML [Architecture de la base de données](https://dbdiagram.io/d/dbZK12EBIKE-6745d6fae9daa85acac4d8a6)**


## 🤝 **Contribuer**
Si vous souhaitez contribuer :
1. Forkez le projet.
2. Créez une branche pour votre fonctionnalité (`git checkout -b feature/AmazingFeature`).
3. Poussez vos modifications (`git push origin feature/AmazingFeature`).
4. Soumettez une pull request.
