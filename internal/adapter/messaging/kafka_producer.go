package messaging

import (
	"context"
	"time"

	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/config"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/interfaces"
	"github.com/IsaiasGC/poc-ports-adapters-scaffold/internal/domain/models"
	"github.com/segmentio/kafka-go"
)

const RESOURCE_NAME = "kafka"

type kafkaRepository struct {
	config *config.KafkaConfig
}

type kafkaProducer struct {
	kafkaRepository
	writer *kafka.Writer
}

func NewMessageProducer(c *config.Configuration) interfaces.MessageProducer {
	return &kafkaProducer{
		kafkaRepository: kafkaRepository{
			config: c.KafkaConfig,
		},
		writer: getKafkaWriter(c.KafkaConfig),
	}
}

func (r *kafkaRepository) HealthCheck(ctx context.Context, check chan<- *models.ComponentCheck) {
	st := time.Now()
	health := &models.ComponentCheck{
		Name:   RESOURCE_NAME,
		Type:   models.TypeBroker,
		Status: models.StatusPass,
	}

	conn, err := kafka.DialContext(ctx, "tcp", r.config.Brokers[0])
	if conn != nil {
		defer conn.Close()

		_, err = conn.Controller()
	}
	if err != nil {
		health.Status = models.StatusFail
		health.Output = err.Error()
	}

	health.Time = time.Since(st)

	check <- health
}

func (r *kafkaProducer) Publish(ctx context.Context, topic string, message []byte) error {
	err := r.writer.WriteMessages(ctx,
		kafka.Message{
			Topic: topic,
			Value: message,
		},
	)

	return err
}

func getKafkaWriter(config *config.KafkaConfig) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:      config.Brokers,
		MaxAttempts:  3,
		Balancer:     &kafka.CRC32Balancer{},
		BatchSize:    10,
		BatchTimeout: 1 * time.Millisecond,
		RequiredAcks: 1,
	})
}
