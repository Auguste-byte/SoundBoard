package corsMiddleware

import (
	"context"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type clientData struct {
	requests int
	lastSeen time.Time
}

var rateLimiter = make(map[string]*clientData)
var mu sync.Mutex

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		mu.Lock()
		data, exists := rateLimiter[ip]
		if !exists || time.Since(data.lastSeen) > 30*time.Second {
			rateLimiter[ip] = &clientData{requests: 1, lastSeen: time.Now()}
		} else {
			data.requests++
			data.lastSeen = time.Now()
			if data.requests > 3 {
				mu.Unlock()
				http.Error(w, "Trop de requêtes", http.StatusTooManyRequests)
				return
			}
		}
		mu.Unlock()

		next.ServeHTTP(w, r)
	})
}

var secretKey = []byte(os.Getenv("SECRET_KEY"))

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Lire le header Authorization
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token manquant", http.StatusUnauthorized)
			return
		}

		// Extraire le token
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Valider le token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token invalide", http.StatusUnauthorized)
			return
		}

		// Extraire les claims (infos dans le token)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["user_id"] == nil {
			http.Error(w, "Token non conforme", http.StatusUnauthorized)
			return
		}

		// Injecter user_id dans le contexte pour les handlers
		ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
