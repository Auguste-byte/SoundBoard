package Auth

import (
	"context"
	"encoding/json"
	"net/http"

	db "musique/database"
)

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)

	var email, pseudo string
	err := db.DB.QueryRow(context.Background(),
		"SELECT email, user_name FROM users WHERE id = $1", userID).Scan(&email, &pseudo)

	if err != nil {
		http.Error(w, "Utilisateur introuvable", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // ✅ ajoute ce header
	json.NewEncoder(w).Encode(map[string]string{
		"email":  email,
		"pseudo": pseudo,
	})
}
