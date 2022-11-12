package server

import (
	"context"
	"log"

	"github.com/megalania/erebus/pkg/proto/agent"
)

type AgentServer struct {
	agent.UnimplementedAgentServer
}

func (a *AgentServer) ReadSingle(ctx context.Context, input *agent.ReadSingleRequest) (*agent.AgentResponse, error) {
	log.Printf("Agent ID: %v", input.Id)
	item := agent.AgentItem{
		Id: input.Id,
	}
	resp := agent.AgentResponse{
		Item: &item,
	}
	return &resp, nil
}
