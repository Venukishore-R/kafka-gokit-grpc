package services

import (
	"context"
	"encoding/json"

	"github.com/Venukishore-R/kafka-gokit-grpc/kafka/producer"
	workers "github.com/Venukishore-R/kafka-gokit-grpc/kafka/worker"
	"github.com/Venukishore-R/kafka-gokit-grpc/models"
	"github.com/go-kit/log"
)

type LoggerService struct {
	logger log.Logger
}

type Service interface {
	SendMessage(ctx context.Context, topic string, partition int64, value1 int64, value2 string) (bool, string, error)
	ConsumeMessage(ctx context.Context, topic string, partition int64) (string, int64, int64, string, error)
}

func NewLoggerService(logger log.Logger) *LoggerService {
	return &LoggerService{
		logger: logger,
	}
}
func (s LoggerService) SendMessage(ctx context.Context, topic string, partition int64, value1 int64, value2 string) (bool, string, error) {
	newMessage := models.Message{
		Topic:     topic,
		Partition: partition,
		Value1:    value1,
		Value2:    value2,
	}

	msgInBytes, err := json.Marshal(newMessage)
	if err != nil {
		return false, "error while marshalling message", err
	}

	err = producer.PushToQueue(models.BrokersUrl, topic, partition, msgInBytes)
	if err != nil {
		return false, "unable to push message to queue", err
	}

	return true, "message pushed to queue", nil
}

func (s LoggerService) ConsumeMessage(ctx context.Context, topic string, partition int64) (string, int64, int64, string, error) {
	var messages *models.Message
	var err error

	messages, err = workers.ConsumeFromQueue(topic, partition)
	if err != nil {
		return "", 0, 0, "", err
	}

	return messages.Topic, messages.Partition, messages.Value1, messages.Value2, nil

}
