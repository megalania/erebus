package server

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/megalania/erebus/pkg/database"
	"github.com/megalania/erebus/pkg/proto/agent"
)

type AgentServer struct {
	agent.UnimplementedAgentServer
	Service *database.AgentService
}

func (a *AgentServer) ReadSingle(ctx context.Context, input *agent.ReadSingleRequest) (*agent.AgentResponse, error) {
	agt, err := a.Service.SelectSingle(
		ctx,
		input.Id,
	)
	if err != nil {
		return nil, err
	}

	c := timestamppb.New(agt.CreatedAt)
	u := timestamppb.New(agt.UpdatedAt)

	item := agent.AgentItem{
		Id:        agt.ID,
		CreatedAt: c,
		UpdatedAt: u,
	}
	resp := agent.AgentResponse{
		Agent: &item,
	}
	return &resp, nil
}
