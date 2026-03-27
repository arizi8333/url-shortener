package main

import (
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 🔌 connect DB
	db := repository.NewPostgresConnection()

	// 🚀 run migration (kalau ada)
	dbURL := "postgres://postgres:postgres@my-postgres:5432/url-shortener?sslmode=disable"
	repository.RunMigration(dbURL)

	// 🧱 init layer
	repo := repository.NewURLRepository(db)
	service := service.NewURLService(repo)
	handler := handler.NewURLHandler(service)

	// 🌐 routes
	r.POST("/shorten", handler.Shorten)
	r.GET("/:code", handler.Redirect)
	r.GET("/stats/:code", handler.Stats)

	r.Run(":8080")
}
