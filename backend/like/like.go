package like

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	db "musique/database"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)

	type LikeRequest struct {
		PostID string `json:"post_id"`
	}

	var req LikeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Vérifier si l'utilisateur a déjà liké ce post
	var exists bool
	err := db.DB.QueryRow(ctx,
		`SELECT EXISTS (
			SELECT 1 FROM "like" WHERE user_id = $1 AND post_id = $2
		)`, userID, req.PostID).Scan(&exists)

	if err != nil {
		http.Error(w, "Erreur lors de la vérification", http.StatusInternalServerError)
		return
	}

	if exists {
		http.Error(w, "Déjà liké", http.StatusConflict)
		return
	}

	_, err = db.DB.Exec(ctx,
		`INSERT INTO "like" (user_id, post_id, created_at)
		VALUES ($1, $2, NOW())`, userID, req.PostID)

	if err != nil {
		http.Error(w, "Erreur lors du like", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Post liké",
	})
}
