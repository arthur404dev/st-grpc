package server

import (
	"context"
	"errors"
	"testing"

	"github.com/arthur404dev/st-grpc/api/st_grpc"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type mockStream struct {
	grpc.ServerStream
	context  context.Context
	messages []*st_grpc.LiveMessage
	header   metadata.MD
}

func (ms *mockStream) Context() context.Context {
	return ms.context
}

func (ms *mockStream) Send(msg *st_grpc.LiveMessage) error {
	ms.messages = append(ms.messages, msg)
	return nil
}

func (ms *mockStream) SetHeader(md metadata.MD) error {
	ms.header = md
	return nil
}

type mockProvider struct {
	err error
}

func (mp *mockProvider) StreamMessages(req *st_grpc.StreamRequest, messageChan chan<- *st_grpc.LiveMessage) error {
	if mp.err != nil {
		return mp.err
	}

	for i := 0; i < 5; i++ {
		messageChan <- &st_grpc.LiveMessage{
			Author:    "test_author",
			Timestamp: 279272197299,
			Content:   "test_message",
		}
	}
	return nil
}

func TestGRPCServer_StreamMessages(t *testing.T) {
	tests := []struct {
		name          string
		providerError error
		expectedError error
	}{
		{
			name:          "provider error",
			providerError: errors.New("provider error"),
			expectedError: nil,
		},
		{
			name:          "success",
			providerError: nil,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider := &mockProvider{err: test.providerError}
			server := NewGRPCServer(provider)

			stream := &mockStream{context: context.Background()}
			err := server.StreamMessages(&st_grpc.StreamRequest{ServiceName: "test_service"}, stream)
			assert.Equal(t, test.expectedError, err)

			if test.providerError != nil {
				assert.Equal(t, "provider error", stream.header["error"][0])
			} else {
				assert.Len(t, stream.messages, 5)
			}
		})
	}
}
