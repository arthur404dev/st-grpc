package st_grpc

type MessageProvider interface {
	StreamMessages(req *StreamRequest, messageChan chan<- *LiveMessage) error
}
