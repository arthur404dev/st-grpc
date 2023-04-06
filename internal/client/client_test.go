package client

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/arthur404dev/st-grpc/api/st_grpc"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	go func() {
		s := grpc.NewServer()
		st_grpc.RegisterLiveMessageServiceServer(s, &mockServer{})
		if err := s.Serve(lis); err != nil {
			panic(fmt.Sprintf("Failed to serve: %v", err))
		}
	}()
}

type mockServer struct {
	st_grpc.UnimplementedLiveMessageServiceServer
}

func (s *mockServer) StreamMessages(_ *st_grpc.StreamRequest, _ st_grpc.LiveMessageService_StreamMessagesServer) error {
	return nil
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGRPCClient_StreamMessages(t *testing.T) {
	tests := []struct {
		name     string
		req      *st_grpc.StreamRequest
		expected error
	}{
		{
			name:     "success",
			req:      &st_grpc.StreamRequest{ServiceName: "test_service"},
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
			if err != nil {
				t.Fatalf("Failed to create connection: %v", err)
			}
			defer conn.Close()

			client := NewGRPCClientWithConnection(conn)
			defer client.Close()

			_, err = client.StreamMessages(context.Background(), test.req)
			assert.Equal(t, test.expected, err)
		})
	}
}
