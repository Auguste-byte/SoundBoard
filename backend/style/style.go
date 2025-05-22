package style

import (
	"context"
	"encoding/json"
	"net/http"

	db "musique/database"
)

type Style struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetStylesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(context.Background(), `SELECT id, name FROM style`)
	if err != nil {
		http.Error(w, "Erreur de récupération des styles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var styles []Style
	for rows.Next() {
		var s Style
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			http.Error(w, "Erreur de lecture des données", http.StatusInternalServerError)
			return
		}
		styles = append(styles, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(styles)
}
