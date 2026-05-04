// internal/router/router.go
package router

import (
	"ks-producer/internal/api"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *api.SimpleMessageHandler) {

	// group endpoints (like @RequestMapping)
	/*messages := r.Group("/simple-messages")
	{
		messages.POST("", handler.SimpleMessage)
	}*/

	r.POST("/simple-messages", handler.SimpleMessage)
}
