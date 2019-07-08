package handler

import (
	"context"
	"time"

	"github.com/smartatransit/fivepoints/api/v1/schedule"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	authorizer Authorizer,
) *Server {
	getScheduleEndpoint := NewGetScheduleEndpoint(tableName, querier, authorizer)
	return &Server{
		getScheduleEndpoint,
	}
}

type envelope struct {
	Response *schedule.GetScheduleResponse
	Err      error
}

func (s *Server) GetSchedule(ctx context.Context, in *schedule.GetScheduleRequest) (*schedule.GetScheduleResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	var respChan chan envelope
	defer cancel()
	go func() {
		defer close(respChan)
		resp, err := s.GetScheduleEndpoint(ctx, in)
		respChan <- envelope{
			resp,
			err,
		}
	}()
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "Deadline Exceeded")
	case r := <-respChan:
		return r.Response, r.Err
	}
}
