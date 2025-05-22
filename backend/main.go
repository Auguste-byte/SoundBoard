package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	// Assure-toi d'importer les bons packages
	auth "musique/auth"
	db "musique/database"
	lk "musique/like"
	mw "musique/middleware"
	pt "musique/post"
	su "musique/registration"
	st "musique/style"
	ws "musique/ws"

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

	r.Handle("/api/registration", mw.RateLimiter(http.HandlerFunc(su.RegistrationHandler))).Methods("POST")
	r.Handle("/api/login", mw.RateLimiter(http.HandlerFunc(auth.LoginHandler))).Methods("POST")
	r.Handle("/api/profile", mw.AuthMiddleware(http.HandlerFunc(auth.GetProfileHandler))).Methods("GET")
	r.Handle("/api/profile", mw.AuthMiddleware(http.HandlerFunc(auth.UpdateProfileHandler))).Methods("PUT")
	r.HandleFunc("/api/posts", pt.GetAllPostsHandler).Methods("GET")
	r.Handle("/api/posts", mw.AuthMiddleware(http.HandlerFunc(pt.CreatePostHandler))).Methods("POST")
	r.HandleFunc("/api/style", st.GetStylesHandler).Methods("GET")
	r.Handle("/api/posts/like", mw.AuthMiddleware(http.HandlerFunc(lk.LikePostHandler))).Methods("POST")

	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	go ws.StartBroadcast()
	r.HandleFunc("/ws", ws.HandleConnections)

	// Route catch-all pour servir React (pour le routing côté client)
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(staticDir, r.URL.Path)

		// Si le fichier demandé existe, on le sert
		if stat, err := os.Stat(requestedPath); err == nil && !stat.IsDir() && filepath.Ext(r.URL.Path) != "" {
			http.ServeFile(w, r, requestedPath)
			return
		}

		// Sinon on sert index.html (React prendra en charge le routing)
		http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
	})

	// Lancer le serveur HTTP sur le port 8080
	log.Println("✅ Serveur Go en ligne sur http://localhost:8080")
	//log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", r))
	log.Fatal(http.ListenAndServe(":8080", r)) // HTTP simple

}
