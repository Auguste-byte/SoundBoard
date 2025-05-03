package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	// Assure-toi d'importer les bons packages
	db "musique/database"
	mw "musique/middleware"
	su "musique/signup"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Fonction principale
func main() {
	// Charger les variables d'environnement depuis le fichier .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Fichier .env non trouvé")
	}

	// Initialiser la connexion à la base de données
	db.InitDB()

	// Créer un nouveau routeur avec gorilla/mux
	r := mux.NewRouter()

	// Ajouter les routes statiques pour React
	staticDir := "../frontend/dist"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Fatalf("❌ Le dossier des fichiers statiques '%s' est introuvable.", staticDir)
	}
	fs := http.FileServer(http.Dir(staticDir))

	r.Use(mux.CORSMethodMiddleware(r)) // Optionnel mais utile
	r.Use(mw.CorsMiddleware)

	r.PathPrefix("/assets/").Handler(fs)

	// Ajouter la route d'inscription à l'API
	r.HandleFunc("/api/signup", su.SignupHandler).Methods("POST")

	// Route catch-all pour servir React (pour le routing côté client)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(staticDir, r.URL.Path)

		// Si c'est un fichier existant (ex: .js, .css), on le sert
		if stat, err := os.Stat(requestedPath); err == nil && !stat.IsDir() && filepath.Ext(r.URL.Path) != "" {
			http.ServeFile(w, r, requestedPath)
			return
		}

		// Sinon, on sert index.html pour permettre le routing React
		http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
	})

	// Lancer le serveur HTTP sur le port 3000
	log.Println("✅ Serveur Go en ligne sur http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
