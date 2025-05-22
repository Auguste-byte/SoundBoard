package Auth

import (
	"context"
	"encoding/json"
	"net/http"

	db "musique/database"
)

type UpdateProfileRequest struct {
	Email  string `json:"email"`
	Pseudo string `json:"pseudo"`
}

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)

	var req UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	// Vérifier si un autre utilisateur utilise déjà cet email
	var existingUserID string
	err := db.DB.QueryRow(context.Background(),
		"SELECT id FROM users WHERE email = $1 AND id != $2", req.Email, userID).Scan(&existingUserID)

	if err == nil {
		// Un autre utilisateur avec cet email existe
		http.Error(w, "Email déjà utilisé", http.StatusConflict)
		return
	}

	// Continuer la mise à jour
	_, err = db.DB.Exec(context.Background(),
		"UPDATE users SET email = $1, user_name = $2, updated_at = NOW() WHERE id = $3",
		req.Email, req.Pseudo, userID)

	if err != nil {
		http.Error(w, "Échec de mise à jour", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Profil mis à jour avec succès",
	})
}
