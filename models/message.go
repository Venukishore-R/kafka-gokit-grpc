package models

type Message struct {
	Topic     string `json:"topic"`
	Partition int64  `json:"partition"`
	Value1    int64  `json:"value1"`
	Value2    string `json:"value2"`
}

var BrokersUrl = []string{"localhost:29092"}
