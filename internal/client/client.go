package client

import (
	"context"

	"github.com/arthur404dev/st-grpc/api/st_grpc"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	conn   *grpc.ClientConn
	client st_grpc.LiveMessageServiceClient
}

func NewGRPCClient(target string) (*GRPCClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	client := st_grpc.NewLiveMessageServiceClient(conn)

	return &GRPCClient{
		conn:   conn,
		client: client,
	}, nil
}

func NewGRPCClientWithConnection(conn *grpc.ClientConn) *GRPCClient {
	return &GRPCClient{
		conn:   conn,
		client: st_grpc.NewLiveMessageServiceClient(conn),
	}
}

func (c *GRPCClient) StreamMessages(ctx context.Context, req *st_grpc.StreamRequest) (st_grpc.LiveMessageService_StreamMessagesClient, error) {
	return c.client.StreamMessages(ctx, req)
}

func (c *GRPCClient) Close() error {
	return c.conn.Close()
}
