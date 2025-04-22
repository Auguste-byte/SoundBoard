package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	staticDir := "../frontend/dist" // <- adapte selon ton dossier

	fs := http.FileServer(http.Dir(staticDir))

	http.Handle("/assets/", fs) // les JS, CSS etc.

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go API"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(staticDir, r.URL.Path)
		_, err := os.Stat(path)
		if os.IsNotExist(err) || r.URL.Path == "/" {
			http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
			return
		}
		fs.ServeHTTP(w, r)
	})

	log.Println("Server on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
