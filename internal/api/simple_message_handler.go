// internal/api/message_handler.go
package api

import (
	"encoding/json"
	"ks-producer/internal/kafka"
	"ks-producer/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*type Producer interface {
	Produce(topic string, key, value []byte) error
}*/

type SimpleMessageHandler struct {
	producer kafka.Producer
}

func NewHandler(p kafka.Producer) *SimpleMessageHandler {
	return &SimpleMessageHandler{producer: p}
}

// @Summary Send simple message
// @Description Send message to simple Kafka topic
// @Tags messages
// @Accept json
// @Produce json
// @Param request body models.SimpleMessage true "Message payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /simple-messages [post]
func (h *SimpleMessageHandler) SimpleMessage(c *gin.Context) {
	var req models.SimpleMessage

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.ID == "" {
		req.ID = "msg-" + time.Now().Format("20060102150405")
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to serialize"})
		return
	}

	err = h.producer.Produce("simple-topic", nil, jsonBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to produce"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "sent"})
}
