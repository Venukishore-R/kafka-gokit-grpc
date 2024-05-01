package endpoints

import (
	"context"

	"github.com/Venukishore-R/kafka-gokit-grpc/services"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SendMessage     endpoint.Endpoint
	ConsumeMessages endpoint.Endpoint
}

func MakeEndpoints(s services.Service) Endpoints {
	return Endpoints{
		SendMessage:     makeSendMessageEndpoint(s),
		ConsumeMessages: makeConsumeMessagesEndpoint(s),
	}
}

func makeSendMessageEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MessageReq)
		success, desc, err := s.SendMessage(ctx, req.Topic, req.Partition, req.Value1, req.Value2)
		return MessageResp{
			Success:     success,
			Description: desc,
		}, err
	}
}

func makeConsumeMessagesEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var finalMsg []*ConsumeMsgResp
		req := request.(ConsumeMsgReq)

		messages, err := s.ConsumeMessage(ctx, req.Topic, req.Partition)

		for _, message := range messages {
			msg := &ConsumeMsgResp{
				Topic:     message.Topic,
				Partition: message.Partition,
				Value1:    message.Value1,
				Value2:    message.Value2,
			}

			finalMsg = append(finalMsg, msg)
		}
		return ConsumeMsgFinalResp{
			ConsumeMsgFinalResp: finalMsg,
		}, err
	}
}
