package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() *pgxpool.Pool {
	return DB
}

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL non défini")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	DB, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("❌ Erreur connexion DB: %v", err)
	}

	if err = DB.Ping(ctx); err != nil {
		log.Fatalf("❌ DB ping failed: %v", err)
	}

	log.Println("✅ Connexion à la base de données réussie")

	initSchema()
}

func initSchema() {
	ctx := context.Background()

	schema := `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users" (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	email VARCHAR(50) NOT NULL UNIQUE,
	user_name VARCHAR(50) NOT NULL UNIQUE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Ajoute la colonne password si elle n'existe pas
ALTER TABLE users ADD COLUMN IF NOT EXISTS password VARCHAR(255);

CREATE OR REPLACE FUNCTION update_user_updated_at()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = now();
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_update_user_updated_at ON "users";
CREATE TRIGGER trigger_update_user_updated_at
BEFORE UPDATE ON "users"
FOR EACH ROW
EXECUTE PROCEDURE update_user_updated_at();

CREATE TABLE IF NOT EXISTS "style" (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	name TEXT NOT NULL UNIQUE
);

INSERT INTO "style" (name) VALUES
  ('Hip-Hop'),
  ('Jazz'),
  ('Rock'),
  ('Électro'),
  ('Classique'),
  ('Pop')
ON CONFLICT (name) DO NOTHING;

CREATE TABLE IF NOT EXISTS "post" (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id UUID NOT NULL REFERENCES "users" (id) ON DELETE CASCADE,
	title VARCHAR(50) NOT NULL,
	content VARCHAR(10000),
	music_file VARCHAR(50),
	style_id UUID NOT NULL REFERENCES "style" (id) ON DELETE CASCADE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	CONSTRAINT idx_post_title UNIQUE (title)
);

CREATE OR REPLACE FUNCTION update_post_updated_at()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = now();
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_update_post_updated_at ON "post";
CREATE TRIGGER trigger_update_post_updated_at
BEFORE UPDATE ON "post"
FOR EACH ROW
EXECUTE PROCEDURE update_post_updated_at();

CREATE TABLE IF NOT EXISTS "comment" (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id UUID NOT NULL REFERENCES "users" (id) ON DELETE CASCADE,
	post_id UUID NULL REFERENCES "post" (id) ON DELETE CASCADE,
	comment_id UUID NULL REFERENCES "comment" (id) ON DELETE CASCADE,
	content TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	CONSTRAINT check_comment CHECK (
		(post_id IS NOT NULL AND comment_id IS NULL) OR 
		(post_id IS NULL AND comment_id IS NOT NULL)
	)
);

CREATE OR REPLACE FUNCTION update_comment_updated_at()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = now();
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_update_comment_updated_at ON "comment";
CREATE TRIGGER trigger_update_comment_updated_at
BEFORE UPDATE ON "comment"
FOR EACH ROW
EXECUTE PROCEDURE update_comment_updated_at();

CREATE TABLE IF NOT EXISTS "like" (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id UUID NOT NULL REFERENCES "users" (id) ON DELETE CASCADE,
	post_id UUID NULL REFERENCES "post" (id) ON DELETE CASCADE,
	comment_id UUID NULL REFERENCES "comment" (id) ON DELETE CASCADE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	CONSTRAINT check_like CHECK (
		(post_id IS NOT NULL AND comment_id IS NULL) OR 
		(post_id IS NULL AND comment_id IS NOT NULL)
	)
);

CREATE TABLE IF NOT EXISTS "sessions" (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	user_id UUID NOT NULL REFERENCES "users" (id) ON DELETE CASCADE,
	token TEXT NOT NULL,
	expiration TIMESTAMPTZ NOT NULL,
	CONSTRAINT idx_session_token UNIQUE (token)
);`

	_, err := DB.Exec(ctx, schema)
	if err != nil {
		log.Fatalf("❌ Erreur lors de l'initialisation du schéma : %v", err)
	}

	log.Println("✅ Schéma SQL initialisé")
}
