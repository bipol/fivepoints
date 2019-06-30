package handler

import (
	"context"

	"github.com/smartatransit/fivepoints/api/v1/schedule"
)

//go:generate counterfeiter . API
type API interface {
	schedule.ScheduleServiceServer
}

type Server struct {
	GetScheduleEndpoint
}

func NewWithDefaultEndpoints(
	tableName string,
	querier DynamoQuerier,
) *Server {
	getScheduleEndpoint := NewGetScheduleEndpoint(tableName, querier)
	return &Server{
		getScheduleEndpoint,
	}
}

func (s *Server) GetSchedule(ctx context.Context, in *schedule.GetScheduleRequest) (*schedule.GetScheduleResponse, error) {
	return s.GetScheduleEndpoint(ctx, in)
}
