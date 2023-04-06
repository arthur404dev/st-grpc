package server

import (
	"net"

	"github.com/arthur404dev/st-grpc/api/st_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GRPCServer struct {
	st_grpc.UnimplementedLiveMessageServiceServer
	provider st_grpc.MessageProvider
}

func NewGRPCServer(provider st_grpc.MessageProvider) *GRPCServer {
	return &GRPCServer{provider: provider}
}

func (s *GRPCServer) StreamMessages(req *st_grpc.StreamRequest, stream st_grpc.LiveMessageService_StreamMessagesServer) error {
	messageChan := make(chan *st_grpc.LiveMessage)

	go func() {
		defer close(messageChan)
		err := s.provider.StreamMessages(req, messageChan)
		if err != nil {
			stream.SetHeader(metadata.Pairs("error", err.Error()))
		}
	}()

	for msg := range messageChan {
		err := stream.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func StartGRPCServer(address string, provider st_grpc.MessageProvider) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	st_grpc.RegisterLiveMessageServiceServer(grpcServer, NewGRPCServer(provider))

	return grpcServer.Serve(lis)
}
