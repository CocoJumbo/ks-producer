package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Producer interface {
	Produce(topic string, key, value []byte) error
}

type Handler struct {
	producer Producer
}

func NewHandler(p Producer) *Handler {
	return &Handler{producer: p}
}

func (h *Handler) Register(r *gin.Engine) {
	r.POST("/simple-messages", h.createMessage) // ✅ updated route
}

// SimpleMessage now contains only Data
type SimpleMessage struct {
	Data string `json:"data" example:"hello kafka"`
}

// @Summary Send simple message
// @Description Send message to simple Kafka topic
// @Tags messages
// @Accept json
// @Produce json
// @Param request body SimpleMessage true "Message payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /simple-messages [post]
func (h *Handler) createMessage(c *gin.Context) {
	var req SimpleMessage

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Since no ID → use nil or generate key
	err := h.producer.Produce(
		"simple-topic",
		nil, // ⚠️ no key → Kafka will distribute randomly
		[]byte(req.Data),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to produce"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "sent"})
}
