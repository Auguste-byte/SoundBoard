package signup

import (
	"context"
	"encoding/json"
	db "musique/database" // Assure-toi d'importer ta propre package DB
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Structure pour la requête d'inscription
type SignupRequest struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Fonction pour hacher un mot de passe
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signupData SignupRequest

	err := json.NewDecoder(r.Body).Decode(&signupData)
	if err != nil {
		jsonError(w, "Données invalides", http.StatusBadRequest)
		return
	}

	hashedPassword, err := hashPassword(signupData.Password)
	if err != nil {
		jsonError(w, "Erreur lors du hachage du mot de passe", http.StatusInternalServerError)
		return
	}

	conn := db.DB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = conn.Exec(ctx, `
		INSERT INTO users (email, user_name, password, created_at, updated_at) 
		VALUES ($1, $2, $3, NOW(), NOW())`,
		signupData.Email, signupData.UserName, hashedPassword)
	if err != nil {
		jsonError(w, "Erreur lors de la création du compte", http.StatusInternalServerError)
		return
	}

	// Réponse de succès au format JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Inscription réussie"})
}

func jsonError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
