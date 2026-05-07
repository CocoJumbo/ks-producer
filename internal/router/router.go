// internal/router/router.go
package router

import (
	"ks-producer/internal/api"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine,
	simpleMessageHandler *api.SimpleMessageHandler,
	threePartitionTopicHashMessageHandler *api.ThreePartitionTopicHashMessageHandler) {

	// group endpoints (like @RequestMapping)
	/*messages := r.Group("/simple-messages")
	{
		messages.POST("", handler.SimpleMessage)
	}*/

	r.POST("/simple-messages", simpleMessageHandler.SimpleMessage)
	r.POST("/3-partitions-hash-balanced", threePartitionTopicHashMessageHandler.ThreePartitionTopicHashMessage)
	r.POST("/3-partitions-round-robin-balanced", simpleMessageHandler.ThreePartitionRoundRobinTopicMessage)
}
