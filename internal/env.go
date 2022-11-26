package internal

import (
	"log"
	"os"
)

func MustGetEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		log.Fatalf("%s env is required.", name)
	}

	return v
}

func GetEnvDefault(name, fallback string) string {
	v := os.Getenv(name)
	if v == "" {
		return fallback
	}

	return v
}
