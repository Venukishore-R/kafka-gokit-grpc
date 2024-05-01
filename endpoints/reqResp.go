package endpoints

type MessageReq struct {
	Topic     string `json:"topic"`
	Partition int64  `json:"partition"`
	Value1    int64  `json:"value1"`
	Value2    string `json:"value2"`
}

type MessageResp struct {
	Success     bool   `json:"success"`
	Description string `json:"description"`
}

type ConsumeMsgReq struct {
	Topic     string `json:"topic"`
	Partition int64  `json:"partition"`
}

type ConsumeMsgResp struct {
	Topic     string `json:"topic"`
	Partition int64  `json:"partition"`
	Value1    int64  `json:"value1"`
	Value2    string `json:"value2"`
}

type ConsumeMsgFinalResp struct {
	ConsumeMsgFinalResp []*ConsumeMsgResp `json:"consumeMsgFinalResp"`
}
