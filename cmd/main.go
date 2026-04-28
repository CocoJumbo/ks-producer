package main

// @title Kafka Producer API
// @version 1.0
// @description Go Kafka producer with Gin
// @host localhost:8081
// @BasePath /

import (
	"log"

	_ "ks-producer/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"ks-producer/internal/api"
	"ks-producer/internal/config"
	"ks-producer/internal/kafka"
)

func main() {
	// Load configuration (env or defaults)
	cfg := config.Load()

	// Create Kafka producer (no error returned in pure Go version)
	producer := kafka.NewProducer(cfg.KafkaBootstrap)
	defer producer.Close()

	// Setup HTTP server
	r := gin.Default()

	// ✅ ADD SWAGGER ROUTE HERE
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register API handlers
	handler := api.NewHandler(producer)
	handler.Register(r)

	log.Printf("🚀 server running on :%s\n", cfg.ServerPort)

	// Start server
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
