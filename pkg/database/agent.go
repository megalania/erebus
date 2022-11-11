package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/megalania/erebus/pkg/model"
)

type AgentService struct {
	Pool *pgxpool.Pool
}

func (a *AgentService) SelectSingle(ctx context.Context, id string) (*model.Agent, error) {
	sql := "SELECT * FROM public.agent WHERE id=$1;"
	row := a.Pool.QueryRow(
		ctx,
		sql,
		id,
	)
	var agent model.Agent
	err := row.Scan(
		&agent.ID,
		&agent.CreatedAt,
		&agent.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (a *AgentService) DeleteSingle(ctx context.Context, id string) (*model.Agent, error) {
	sql := "DELETE FROM public.agent WHERE id=$1;"
	row := a.Pool.QueryRow(
		ctx,
		sql,
		id,
	)
	var agent model.Agent
	err := row.Scan(
		&agent.ID,
		&agent.CreatedAt,
		&agent.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}
