package auth

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
	"zk12ebike/internal/cookies"
	"zk12ebike/internal/database"

	"golang.org/x/crypto/bcrypt"
)

// Fonction pour charger les variables d'environnement à partir du fichier .env
func LoadEnvFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'ouverture du fichier .env : %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Ignorer les lignes vides ou les commentaires
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		// Séparer la clé et la valeur (clé=valeur)
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("Ligne mal formée dans le fichier .env : %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Définir la variable d'environnement
		err := os.Setenv(key, value)
		if err != nil {
			return fmt.Errorf("Erreur lors de la définition de la variable d'environnement : %s", key)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Erreur lors de la lecture du fichier .env : %v", err)
	}

	return nil
}

// Gestionnaire de la route /login
func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	redirectURI := os.Getenv("REDIRECT_URI")

	authURL := "https://accounts.google.com/o/oauth2/v2/auth"
	params := fmt.Sprintf(
		"?client_id=%s&redirect_uri=%s&response_type=code&scope=openid%%20email%%20profile&state=random_csrf_token",
		clientID, redirectURI,
	)

	http.Redirect(w, r, authURL+params, http.StatusFound)
}

// Gestionnaire de la route /callback pour récupérer le token et les données utilisateur
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code d'autorisation manquant", http.StatusBadRequest)
		return
	}

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURI := os.Getenv("REDIRECT_URI")

	tokenURL := "https://oauth2.googleapis.com/token"
	data := fmt.Sprintf(
		"code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code",
		code, clientID, clientSecret, redirectURI,
	)

	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		http.Error(w, "Erreur lors de la demande de token", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture de la réponse", http.StatusInternalServerError)
		return
	}

	var tokenResponse map[string]interface{}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse de la réponse JSON", http.StatusInternalServerError)
		return
	}

	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		http.Error(w, "Access token manquant dans la réponse", http.StatusInternalServerError)
		return
	}

	userInfoURL := "https://www.googleapis.com/oauth2/v2/userinfo"
	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		http.Error(w, "Erreur lors de la création de la requête", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	userInfoResp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des informations utilisateur", http.StatusInternalServerError)
		return
	}
	defer userInfoResp.Body.Close()

	userInfoBody, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture de la réponse", http.StatusInternalServerError)
		return
	}

	// Définition d'une structure pour analyser les informations de l'utilisateur
	var userInfo struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		Picture   string `json:"picture"`
		FirstName string `json:"given_name"`
		LastName  string `json:"family_name"`
	}

	err = json.Unmarshal(userInfoBody, &userInfo)
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse des informations utilisateur", http.StatusInternalServerError)
		return
	}
	role := "user"
	if userInfo.FirstName == "Zanakan" {
		role = "admin"
	}

	if database.GetUserId(userInfo.FirstName) == 0 {
		password, err := generateRandomPassword(12)
		if err != nil {
			fmt.Println("Erreur lors de la génération du mot de passe :", err)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Erreur lors du hachage du mot de passe", http.StatusInternalServerError)
			return
		}
		if err := database.SaveUserToDB(userInfo.FirstName, userInfo.Email, string(hashedPassword), role, userInfo.Picture); err != nil {
			fmt.Println("error to add new user in ")
			return
		}
	}
	cookies.CreateSession(w, database.GetUserId(userInfo.FirstName), userInfo.FirstName, role)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func generateRandomPassword(length int) (string, error) {
	// Définir les caractères possibles pour le mot de passe
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"
	var password strings.Builder

	for i := 0; i < length; i++ {
		// Générer un index aléatoire pour un caractère dans charSet
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", err
		}

		// Ajouter le caractère sélectionné au mot de passe
		password.WriteByte(charSet[index.Int64()])
	}

	return password.String(), nil
}
