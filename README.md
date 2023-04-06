# Streamer Toolkit gRPC

This is the gRPC service for Streamer Toolkit. It provides a central server for real-time messages from various providers, such as YouTube and Twitch. The service communicates using gRPC and exposes a WebSocket server for frontend consumption.

## Technologies

- [Go](https://golang.org/doc/) - The programming language used for the implementation of the services
- [gRPC](https://grpc.io/docs/languages/go/) - A high-performance, open-source universal RPC framework for communication between services
- [Protocol Buffers](https://developers.google.com/protocol-buffers) - A language- and platform-neutral mechanism for serializing structured data

## Modules

- `api/st_grpc`: Contains the Protocol Buffers service definition and generated Go code
- `internal/client`: Handles gRPC client creation and management for provider services
- `internal/server`: Implements the gRPC server that handles incoming connections from provider services

## Usage

### Server

Start the gRPC server by calling the `StartGRPCServer` function:

```golang
import (
"github.com/arthur404dev/st-grpc/internal/server"
)

func main() {
address := "localhost:50051"
provider := NewCustomProvider() // Replace with your custom provider implementing the `st_grpc.MessageProvider` interface
err := server.StartGRPCServer(address, provider)
if err != nil {
log.Fatalf("Failed to start gRPC server: %v", err)
}
}
```

### Client

Create a gRPC client by calling the `client.New` function:

```golang
import (
"github.com/arthur404dev/st-grpc/internal/client"
)

func main() {
address := "localhost:50051"
grpcClient, err := client.New(address)
if err != nil {
log.Fatalf("Failed to create gRPC client: %v", err)
}
defer grpcClient.Close()
}
```

## Contributing

Feel free to open an issue or submit a pull request if you have any improvements or suggestions for the project. Make sure to follow the existing code style and provide tests and documentation where necessary.

## License

Streamer Toolkit gRPC is released under the MIT License. See [LICENSE](./LICENSE) for details.
