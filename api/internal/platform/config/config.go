// Package config reads settings from environment variables.
// Locally they come from api/.env (see .env.example); on AWS from the SAM template.
package config

import (
	"os"
	"strings"
)

// Config holds runtime settings.
type Config struct {
	Port           string   // HTTP port for local runs (Lambda ignores it)
	Env            string   // "local", "dev" or "prod"
	AllowedOrigins []string // CORS origins allowed to call the API
	StoreBackend   string   // "memory" (default) or "dynamodb" (Sprint 2)
	TableName      string   // DynamoDB table name (Sprint 2)
}

// Load reads the environment with sensible local defaults.
func Load() Config {
	return Config{
		Port:           get("PORT", "8080"),
		Env:            get("APP_ENV", "local"),
		AllowedOrigins: split(get("ALLOWED_ORIGINS", "http://localhost:3000")),
		StoreBackend:   get("STORE_BACKEND", "memory"),
		TableName:      get("TABLE_NAME", "cakeshop-dev"),
	}
}

func get(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

func split(s string) []string {
	parts := strings.Split(s, ",")
	out := parts[:0]
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}
