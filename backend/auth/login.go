package Auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	db "musique/database" // si tu veux une fonction jsonError

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Identifier string `json:"identifier"` // email ou username
	Password   string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Recherche de l'utilisateur par email OU username
	query := `
		SELECT id, user_name, password
		FROM users
		WHERE email = $1 OR user_name = $1
	`
	var userID, username, hashedPassword string
	err := db.DB.QueryRow(ctx, query, strings.ToLower(loginData.Identifier)).Scan(&userID, &username, &hashedPassword)
	if err != nil {
		http.Error(w, "Utilisateur introuvable", http.StatusUnauthorized)
		return
	}

	// Vérification du mot de passe
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginData.Password)); err != nil {
		http.Error(w, "Mot de passe invalide", http.StatusUnauthorized)
		return
	}

	// Génération du token
	token, err := GenerateJWT(userID, username)
	if err != nil {
		http.Error(w, "Erreur lors de la génération du token", http.StatusInternalServerError)
		return
	}

	// Réponse JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Connexion réussie",
		"token":   token,
	})
}
