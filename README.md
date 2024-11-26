
# 🚲 **BikeRentBuy**

Une plateforme en ligne pour louer et acheter des vélos.

## 📋 **Description**
Ce projet a pour objectif de créer une application web permettant :
- De louer des vélos pour une durée déterminée.
- D'acheter des vélos en ligne, neufs ou d'occasion.
- De suivre les commandes et réservations via un espace utilisateur.

## 🛠️ **Technologies utilisées**
- **Langage backend :** Go (Golang)
- **Framework web :** Echo ou Gin
- **Base de données :** PostgreSQL ou SQLite
- **Versionnement :** Git/GitHub
- **Frontend :** HTML, CSS, JavaScript (intégration minimale côté Go avec templates)
- **Paiement :** Intégration avec Stripe ou PayPal (si nécessaire)

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

## 🌟 **User Stories**
1. **En tant qu'utilisateur non inscrit,**
   Je veux pouvoir consulter les vélos disponibles  
   Afin de décider si je veux m'inscrire ou louer/acheter un vélo.

2. **En tant qu'utilisateur inscrit,**
   Je veux pouvoir réserver un vélo pour une durée déterminée  
   Afin de m'assurer qu'il soit disponible à la date choisie.

3. **En tant qu'administrateur,**
   Je veux pouvoir ajouter ou modifier les vélos dans le catalogue  
   Afin de maintenir un inventaire précis et à jour.

4. **En tant qu'utilisateur inscrit,**
   Je veux pouvoir accéder à mon espace personnel  
   Afin de consulter l'historique de mes commandes ou réservations.

---

## 🚀 **Workflow Agile**

### **Sprint 1 : Analyse et Conception**
- **User Stories :**
  - Explorer les vélos sans se connecter.
  - Page d’accueil avec mise en avant des offres spéciales.
- **Tâches :**
  - Rédiger toutes les User Stories.
  - Créer le schéma UML pour la base de données.
  - Réaliser les wireframes des principales pages.

### **Sprint 2 : Développement Backend**
- **User Stories :**
  - Enregistrer les utilisateurs dans une base de données.
  - Réserver un vélo.
- **Tâches :**
  - Créer les modèles (users, bikes, orders).
  - Configurer la base de données avec Go.
  - Développer les API pour les vélos et les commandes.

### **Sprint 3 : Développement Frontend**
- **User Stories :**
  - Afficher une liste de vélos avec filtres (prix, disponibilité).
  - Intégrer un panier d’achat.
- **Tâches :**
  - Créer les templates HTML pour chaque page.
  - Ajouter de l'interactivité avec JavaScript.

### **Sprint 4 : Tests et Intégration**
- **User Stories :**
  - Vérifier que les réservations sont correctement enregistrées.
  - S'assurer que les paiements fonctionnent.
- **Tâches :**
  - Effectuer des tests unitaires sur le backend.
  - Tester l’interface utilisateur.

### **Sprint 5 : Déploiement et Documentation**
- Préparer une documentation claire pour les développeurs et les utilisateurs.
- Déployer l’application sur un serveur (ex. : Heroku, AWS).

---

## 🛠️ **Installation et Lancement**
### 1. Clone du projet
```bash
git clone https://github.com/<ton-username>/bikerentbuy.git
cd bikerentbuy
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
Ajoutez les migrations SQL (via un outil comme GORM ou SQL directement). Exemple :
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

---

## 📖 **Documentation technique**
- **Framework choisi :** Gin/Echo pour les routes web.
- **Modèle de données :** Voir la section UML ci-dessous.

---

## 🖼️ **Diagramme UML (Architecture de la base de données)**

Voici un diagramme de la structure relationnelle de ta base de données :

- **Table `users` :**
  - `id` (PK)
  - `name`
  - `email`
  - `password`
  - `created_at`

- **Table `bikes` :**
  - `id` (PK)
  - `name`
  - `type` (route, VTT, électrique...)
  - `price` (location/achat)
  - `status` (disponible, loué, vendu)
  - `created_at`

- **Table `orders` :**
  - `id` (PK)
  - `user_id` (FK)
  - `bike_id` (FK)
  - `type` (location/achat)
  - `start_date` (pour location)
  - `end_date` (pour location)
  - `total_price`
  - `status` (en attente, validé)
  - `created_at`

---

## 🤝 **Contribuer**
Si vous souhaitez contribuer :
1. Forkez le projet.
2. Créez une branche pour votre fonctionnalité (`git checkout -b feature/AmazingFeature`).
3. Poussez vos modifications (`git push origin feature/AmazingFeature`).
4. Soumettez une pull request.
