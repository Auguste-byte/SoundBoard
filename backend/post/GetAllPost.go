package post

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	db "musique/database"
)

type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	MusicFile string `json:"music_file"`
	Author    string `json:"author"`
	StyleID   string `json:"style_id"`
	CreatedAt string `json:"created_at"`
}

func GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(context.Background(), `
		SELECT 
			p.id::text, 
			p.title, 
			COALESCE(p.content, ''), 
			p.music_file, 
			u.user_name, 
			p.style_id::text, 
			p.created_at::text
		FROM post p
		JOIN users u ON u.id = p.user_id
		ORDER BY p.created_at DESC
	`)
	if err != nil {
		log.Println("❌ Erreur requête SQL :", err)
		http.Error(w, "Erreur de récupération des posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.MusicFile,
			&post.Author,
			&post.StyleID,
			&post.CreatedAt,
		)
		if err != nil {
			log.Println("❌ Erreur lecture d'une ligne :", err)
			http.Error(w, "Erreur de lecture des posts", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	log.Printf("✅ %d post(s) récupéré(s)", len(posts))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
