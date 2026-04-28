package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(broker string) *Producer {
	writer := &kafka.Writer{
		Addr: kafka.TCP(broker),
		//Topic:    "orders",      // default topic (can override per message)
		Balancer: &kafka.Hash{}, // ensures same key → same partition

		RequiredAcks: kafka.RequireAll, // durability
		Async:        false,            // sync writes (simpler for learning)

		WriteTimeout: 5 * time.Second,
	}

	return &Producer{writer: writer}
}

func (p *Producer) Produce(topic string, key, value []byte) error {
	msg := kafka.Message{
		Topic: topic,
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}

	err := p.writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Printf("❌ failed to write message: %v", err)
		return err
	}

	log.Printf("✅ message sent to topic=%s key=%s", topic, key)
	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
