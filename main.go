package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Venukishore-R/kafka-gokit-grpc/endpoints"
	"github.com/Venukishore-R/kafka-gokit-grpc/protos"
	"github.com/Venukishore-R/kafka-gokit-grpc/services"
	"github.com/Venukishore-R/kafka-gokit-grpc/transports"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"google.golang.org/grpc"
)

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := services.NewLoggerService(logger)
	makEndpoints := endpoints.MakeEndpoints(service)
	server := transports.NewMyServer(makEndpoints, logger)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":5010")
	if err != nil {
		logger.Log("during", "listen", "err", err)
		os.Exit(1)
	}

	go func() {
		serverRegistarar := grpc.NewServer()
		protos.RegisterMessageServiceServer(serverRegistarar, &server)
		level.Info(logger).Log("msg", "server created successfully")
		serverRegistarar.Serve(grpcListener)
	}()

	logger.Log("exiting with error", <-errs)

}
