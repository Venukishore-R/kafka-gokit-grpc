package workers

import (
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
	"github.com/Venukishore-R/kafka-gokit-grpc/models"
)

type Result struct {
	allCaps string
	length  int
}

func ConnectToConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ConsumeFromQueue(topic string, partition int64) ([]*models.Message, error) {
	c := make(chan *models.Message)

	var totalMsg []*models.Message

	worker, err := ConnectToConsumer(models.BrokersUrl)
	if err != nil {
		return nil, err
	}

	pc, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		return nil, err
	}

	go storeMsgInArray(pc, c)

	go func() {
		for res := range c {
			totalMsg = append(totalMsg, res)
		}
	}()

	time.Sleep(1 * time.Second)
	return totalMsg, nil
}

func storeMsgInArray(pc sarama.PartitionConsumer, c chan *models.Message) {

	defer close(c)
	pcMesssages := pc.Messages()

	for msg := range pcMesssages {
		var messsage *models.Message

		json.Unmarshal(msg.Value, &messsage)

		c <- messsage
	}
}
