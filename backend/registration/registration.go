package signup

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	auth "musique/auth"
	db "musique/database"
	utl "musique/utils"

	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var signupData SignupRequest

	err := json.NewDecoder(r.Body).Decode(&signupData)
	if err != nil {
		jsonError(w, "Données invalides", http.StatusBadRequest)
		return
	}

	exists, err := utl.UserExists(signupData.Email, signupData.UserName)
	if err != nil {
		jsonError(w, "Erreur lors de la vérification de l'utilisateur", http.StatusInternalServerError)
		return
	}
	if exists {
		jsonError(w, "Email ou nom d'utilisateur déjà utilisé", http.StatusConflict)
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

	var userID string
	err = conn.QueryRow(ctx, `
		INSERT INTO users (email, user_name, password, created_at, updated_at) 
		VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`,
		signupData.Email, signupData.UserName, hashedPassword).Scan(&userID)

	if err != nil {
		jsonError(w, "Erreur lors de la création du compte", http.StatusInternalServerError)
		return
	}

	token, err := auth.GenerateJWT(userID, signupData.UserName)
	if err != nil {
		jsonError(w, "Erreur génération token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Inscription réussie",
		"token":   token,
	})
}

// ✅ Correction ici : champ "error" au lieu de "message"
func jsonError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
