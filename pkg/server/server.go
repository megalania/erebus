package server

import (
	"github.com/megalania/erebus/pkg/database"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/megalania/erebus/pkg/proto/agent"
)

var agentServer AgentServer

func init() {
	pool, err := database.NewPool(
		"postgres://erebus:am2U1obbvyMVDnIT0drX@localhost:5432/erebus",
	)
	if err != nil {
		log.Fatal(err)
	}
	AgentService := database.AgentService{
		Pool: pool,
	}
	agentServer = AgentServer{
		Service: &AgentService,
	}
}

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
	//ag := AgentServer{}
	agent.RegisterAgentServer(server, &agentServer)

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
