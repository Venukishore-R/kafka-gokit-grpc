package workers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/Venukishore-R/kafka-gokit-grpc/models"
)

func ConnectToConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ConsumeFromQueue(topic string, partition int64) (*models.Message, error) {

	worker, err := ConnectToConsumer(models.BrokersUrl)
	if err != nil {
		return nil, err
	}

	partitions, err := worker.Partitions(topic)
	if err != nil {
		return nil, err
	}

	topics, err := worker.Topics()
	if err != nil {
		return nil, err
	}

	log.Println("topics ", topics)
	
	msgCount := 0
	var message models.Message

	for _, partition := range partitions {
		pc, err := worker.ConsumePartition(topic, partition, sarama.OffsetOldest)
		if err != nil {
			return nil, err
		}
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				msgCount++

				json.Unmarshal(msg.Value, &message)
				fmt.Printf("Received message on topic %s || Partition %d || message: %s || %d || count %d\n", message.Topic, message.Partition, message.Value2, message.Value1, msgCount)
			}
		}(pc)
	}

	log.Println("msg", message)
	return &message, nil
}
