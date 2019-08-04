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
	GetArrivalEstimatesEndpoint
}

func NewWithDefaultEndpoints(
	tableName string,
	querier DynamoQuerier,
	authorizer Authorizer,
) *Server {
	getScheduleEndpoint := NewGetArrivalEstimatesEndpoint(tableName, querier, authorizer)
	return &Server{
		getScheduleEndpoint,
	}
}

type envelope struct {
	Response *schedule.GetArrivalEstimatesResponse
	Err      error
}

func (s *Server) GetArrivalEstimates(ctx context.Context, in *schedule.GetArrivalEstimatesRequest) (*schedule.GetArrivalEstimatesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	respChan := make(chan envelope, 1)
	defer cancel()
	go func() {
		defer close(respChan)
		resp, err := s.GetArrivalEstimatesEndpoint(ctx, in)
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
