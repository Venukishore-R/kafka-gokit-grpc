package producer

import (
	"log"

	"github.com/IBM/sarama"
)

func ConnectToProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func PushToQueue(brokersUrl []string, topic string, partition int64, message []byte) error {
	producer, err := ConnectToProducer(brokersUrl)
	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: int32(partition),
		Value:     sarama.StringEncoder(message),
	}

	partiotion, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Messge %s msg send to partition %d , partiotion, in the offset %d", msg, partiotion, offset)
	return nil
}
