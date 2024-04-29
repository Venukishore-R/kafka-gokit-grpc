package transports

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/Venukishore-R/kafka-gokit-grpc/endpoints"
	"github.com/Venukishore-R/kafka-gokit-grpc/protos"
	"github.com/go-kit/kit/transport/grpc"
)

type MyServer struct {
	sendMessage     grpc.Handler
	consumeMessages grpc.Handler
	protos.UnimplementedMessageServiceServer
}

func NewMyServer(endpoints endpoints.Endpoints, logger log.Logger) MyServer {
	return MyServer{
		sendMessage: grpc.NewServer(
			endpoints.SendMessage,
			decodeMessageReq,
			encodeMessageResp,
		),
		consumeMessages: grpc.NewServer(
			endpoints.ConsumeMessages,
			decodeConsumeMsgReq,
			encodeConsumeMsgResp,
		),
	}
}

func (s *MyServer) SendMessage(ctx context.Context, request *protos.ProducerMessageReq) (*protos.ProducerMessageResp, error) {
	_, resp, err := s.sendMessage.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*protos.ProducerMessageResp), nil
}
func decodeMessageReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*protos.ProducerMessageReq)
	return endpoints.MessageReq{
		Topic:     req.Topic,
		Partition: req.Partition,
		Value1:    req.Value1,
		Value2:    req.Value2,
	}, nil
}

func encodeMessageResp(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.MessageResp)

	return &protos.ProducerMessageResp{
		Success:     resp.Success,
		Description: resp.Description,
	}, nil
}

func (s *MyServer) ConsumeMessage(ctx context.Context, request *protos.ConsumeMessageReq) (*protos.ConsumerMessageResp, error) {
	_, resp, err := s.consumeMessages.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return resp.(*protos.ConsumerMessageResp), nil
}

func decodeConsumeMsgReq(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*protos.ConsumeMessageReq)
	return endpoints.ConsumeMsgReq{
		Topic:     req.Topic,
		Partition: req.Partition,
	}, nil
}

func encodeConsumeMsgResp(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ConsumeMsgResp)
	return &protos.ConsumerMessageResp{
		Topic:     resp.Topic,
		Partition: resp.Partition,
		Value1:    resp.Value1,
		Value2:    resp.Value2,
	}, nil
}
