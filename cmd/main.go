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

	defaultBalancerProducer := kafka.DefaultProducer(cfg.KafkaBootstrap, false)
	defer defaultBalancerProducer.Close()

	hashBalancerProducer := kafka.DefaultProducer(cfg.KafkaBootstrap, true)
	defer hashBalancerProducer.Close()

	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	SimpleMessageHandler := api.SimpleMessageHandle(defaultBalancerProducer)
	ThreePartitionTopicHashMessageHandler := api.ThreePartitionTopicHashMessageHandle(hashBalancerProducer)

	// ✅ clean routing
	router.RegisterRoutes(r,
		SimpleMessageHandler,
		ThreePartitionTopicHashMessageHandler)

	log.Printf("🚀 server running on :%s\n", cfg.ServerPort)

	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
