package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer interface {
	Produce(topic string, key, value []byte) error
	Close() error
}

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewProducer(broker string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr: kafka.TCP(broker),
		//Topic:    "orders",      // default topic (can override per message)
		Balancer: &kafka.Hash{}, // ensures same key → same partition

		RequiredAcks: kafka.RequireAll, // durability
		Async:        false,            // sync writes (simpler for learning)

		WriteTimeout: 5 * time.Second,
	}

	return &KafkaProducer{writer: writer}
}

func (p *KafkaProducer) Produce(topic string, key, value []byte) error {
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

func (p *KafkaProducer) Close() error {
	return p.writer.Close()
}
