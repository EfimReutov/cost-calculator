package grpc

import (
	"cost-calculator/gen/proto/base_service"
	"cost-calculator/store"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	store store.StorePostgres
	base_service.UnimplementedTestServiceServer
}

// NewServer returns *server
func NewServer(store store.StorePostgres) *server {
	return &server{
		store: store,
	}
}

// Run runs the server.
func Run(store store.StorePostgres, address string) error {
	s := grpc.NewServer()
	srv := NewServer(store)
	base_service.RegisterTestServiceServer(s, srv)

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Println("gRPC server is running")
	return s.Serve(l)
}
