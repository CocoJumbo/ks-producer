package main

import (
	"log"

	_ "ks-producer/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"ks-producer/internal/api"
	"ks-producer/internal/config"
	"ks-producer/internal/kafka"
	"ks-producer/internal/router"
)

func main() {
	cfg := config.Load()

	producer := kafka.NewProducer(cfg.KafkaBootstrap)
	defer producer.Close()

	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	handler := api.NewHandler(producer)

	// ✅ clean routing
	router.RegisterRoutes(r, handler)

	log.Printf("🚀 server running on :%s\n", cfg.ServerPort)

	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
