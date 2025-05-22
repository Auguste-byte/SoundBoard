package post

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	db "musique/database"
	ws "musique/ws" // 👈 importe ton package websocket
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)

	// 🔽 Parse le formulaire
	err := r.ParseMultipartForm(10 << 20) // 10 Mo
	if err != nil {
		http.Error(w, "Requête invalide (multipart)", http.StatusBadRequest)
		log.Println("ParseMultipartForm error:", err)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	styleID := r.FormValue("style_id")

	if styleID == "" || title == "" {
		http.Error(w, "Champs requis manquants", http.StatusBadRequest)
		log.Println("Champ requis manquant : styleID ou title")
		return
	}

	// 🔽 Récupération du fichier audio
	file, header, err := r.FormFile("audio")
	if err != nil {
		http.Error(w, "Fichier audio manquant", http.StatusBadRequest)
		log.Println("Fichier audio manquant:", err)
		return
	}
	defer file.Close()

	// 🔽 Création du nom et du chemin de sauvegarde
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), header.Filename)
	uploadPath := filepath.Join("./uploads", filename)

	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		http.Error(w, "Erreur de création du dossier upload", http.StatusInternalServerError)
		log.Println("Erreur dossier upload:", err)
		return
	}

	dst, err := os.Create(uploadPath)
	if err != nil {
		http.Error(w, "Erreur de création de fichier", http.StatusInternalServerError)
		log.Println("Erreur création fichier:", err)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Erreur d'enregistrement du fichier", http.StatusInternalServerError)
		log.Println("Erreur io.Copy:", err)
		return
	}

	// 🔽 Insertion dans la base de données
	var postID string
	query := `
		INSERT INTO post (user_id, title, content, music_file, style_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id
	`
	err = db.DB.QueryRow(context.Background(), query,
		userID, title, content, filename, styleID,
	).Scan(&postID)
	if err != nil {
		http.Error(w, "Erreur d'insertion dans la base", http.StatusInternalServerError)
		log.Println("Erreur DB:", err)
		return
	}

	// 🔽 Récupération du nom de l'auteur
	var author string
	err = db.DB.QueryRow(context.Background(),
		`SELECT user_name FROM users WHERE id = $1`, userID).Scan(&author)
	if err != nil {
		author = "Inconnu"
	}

	createdAt := time.Now().Format(time.RFC3339)

	// 🔽 Envoi via WebSocket aux clients connectés
	newPost := map[string]interface{}{
		"id":         postID,
		"title":      title,
		"content":    content,
		"music_file": filename,
		"author":     author,
		"style_id":   styleID,
		"created_at": createdAt,
	}
	msg, _ := json.Marshal(newPost)
	ws.Send(msg) // ✅ envoie le post en temps réel

	// 🔽 Réponse HTTP classique
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    "Post publié ✅",
		"id":         postID,
		"title":      title,
		"content":    content,
		"music_file": filename,
		"author":     author,
		"style_id":   styleID,
		"created_at": createdAt,
	})
}
