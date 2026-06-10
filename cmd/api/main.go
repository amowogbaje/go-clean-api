package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"go_clean_api/internal/delivery/http"
	"go_clean_api/internal/repository"
	"go_clean_api/internal/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// 1. Ensure uploads directory exists on host
	os.MkdirAll("uploads", os.ModePerm)

	// 2. Connect to Postgres
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// 3. Auto-migrate table for boilerplate purposes
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS medias (
		id SERIAL PRIMARY KEY,
		url TEXT NOT NULL,
		type TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	// 4. Initialize Clean Architecture Layers
	mediaRepo := repository.NewPostgresMediaRepository(db)
	mediaUseCase := usecase.NewMediaUsecase(mediaRepo)

	// 5. Setup Gin Router
	r := gin.Default()

	// Serve the uploads folder statically so files can be accessed via URL
	r.Static("/uploads", "./uploads")

	// Map routes
	http.NewMediaHandler(r, mediaUseCase)

	// 6. Start server
	log.Println("Server running on port 8080")
	r.Run(":8080")
}