package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func Run(addr string) error {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	lis, err := net.Listen(
		"tcp",
		addr,
	)
	if err != nil {
		return nil
	}
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	cha := make(chan error)
	go func() {
		cha <- server.Serve(lis)
	}()
	select {
	case <-shutdown:
		log.Println(
			"SIGINT has been received",
		)
		server.GracefulStop()
	}
	return nil
}
