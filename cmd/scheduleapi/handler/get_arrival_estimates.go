package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smartatransit/fivepoints/api/v1/schedule"
	"github.com/smartatransit/fivepoints/pkg/martaapi"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DynamoQuerier
type DynamoQuerier interface {
	QueryPagesWithContext(ctx aws.Context, input *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool, opts ...request.Option) error
	QueryWithContext(ctx aws.Context, input *dynamodb.QueryInput, opts ...request.Option) (*dynamodb.QueryOutput, error)
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . GetArrivalEstimatesEndpoint
type GetArrivalEstimatesEndpoint func(context.Context, *schedule.GetArrivalEstimatesRequest) (*schedule.GetArrivalEstimatesResponse, error)

func ValidateRequest(ctx context.Context, in *schedule.GetArrivalEstimatesRequest) error {
	var errStrings []string
	if in == nil {
		errStrings = append(errStrings, "request body is nil")
	}
	if strings.TrimSpace(in.GetDestination()) == "" {
		errStrings = append(errStrings, "destination is nil")
	}
	if strings.TrimSpace(in.GetStation()) == "" {
		errStrings = append(errStrings, "station is nil")
	}
	if in.GetStartDate() == nil {
		errStrings = append(errStrings, "start date is nil")
	}
	if in.GetEndDate() == nil {
		errStrings = append(errStrings, "end date is nil")
	}
	if in.GetStartDate().GetSeconds() > in.GetEndDate().GetSeconds() {
		errStrings = append(errStrings, "start date must be before end date")
	}
	s, err := ptypes.Timestamp(in.GetStartDate())
	if err != nil {
		errStrings = append(errStrings, "start must be RFC3339 encoded")
	}
	e, err := ptypes.Timestamp(in.GetEndDate())
	if err != nil {
		errStrings = append(errStrings, "end must be RFC3339 encoded")
	}
	if e.Format("2006-01-02") != s.Format("2006-01-02") {
		errStrings = append(errStrings, "start and end must be on the same day")
	}
	if len(errStrings) != 0 {
		return errors.New(fmt.Sprintf("validation errors: %s", strings.Join(errStrings, ", ")))
	}
	return nil
}

func GetArrivalEstimatesRequestToDynamoQuery(in *schedule.GetArrivalEstimatesRequest, tableName string) (*dynamodb.QueryInput, error) {
	s, err := ptypes.Timestamp(in.GetStartDate())
	if err != nil {
		return nil, err
	}
	e, err := ptypes.Timestamp(in.GetEndDate())
	if err != nil {
		return nil, err
	}
	primaryKey := fmt.Sprintf("%s_%s_%s", in.GetStation(), in.GetDestination(), s.Format("2006-01-02"))
	keyCondition := expression.
		Key("PrimaryKey").
		Equal(expression.Value(primaryKey)).
		And(expression.Key("SortKey").
			Between(expression.Value(s.Format(time.RFC3339)), expression.Value(e.Format(time.RFC3339))))
	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}
	if in.GetLastEvaluatedKey() != "" {
		lastEval := make(map[string]*dynamodb.AttributeValue)
		decoded, err := base64.StdEncoding.DecodeString(in.GetLastEvaluatedKey())
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(decoded, &lastEval); err != nil {
			return nil, err
		}
		input.ExclusiveStartKey = lastEval
	}

	return input, nil
}

//go:generate counterfeiter . Authorizer
type Authorizer interface {
	IsAuthorized(ctx context.Context) (bool, error)
}

func NewGetArrivalEstimatesEndpoint(
	tableName string,
	querier DynamoQuerier,
	authorizer Authorizer,
) func(context.Context, *schedule.GetArrivalEstimatesRequest) (*schedule.GetArrivalEstimatesResponse, error) {
	return func(ctx context.Context, in *schedule.GetArrivalEstimatesRequest) (*schedule.GetArrivalEstimatesResponse, error) {
		var schedules []martaapi.ArrivalEstimate

		if ok, err := authorizer.IsAuthorized(ctx); !ok {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		err := ValidateRequest(ctx, in)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		queryInput, err := GetArrivalEstimatesRequestToDynamoQuery(in, tableName)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		output, err := querier.QueryWithContext(ctx, queryInput)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		if len(output.Items) == 0 {
			return &schedule.GetArrivalEstimatesResponse{}, nil
		}
		err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &schedules)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		var encodedBody string
		if output.LastEvaluatedKey != nil {
			marshal, err := json.Marshal(output.LastEvaluatedKey)
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			encodedBody = base64.StdEncoding.EncodeToString(marshal)
		}

		return &schedule.GetArrivalEstimatesResponse{
			LastEvaluatedKey: encodedBody,
			ResultLength:     int32(len(schedules)),
			ArrivalEstimates: MartaArrivalEstimatesToProtoArrivalEstimates(schedules),
		}, nil
	}
}

func MartaArrivalEstimatesToProtoArrivalEstimates(martaScheds []martaapi.ArrivalEstimate) []*schedule.ArrivalEstimate {
	var protoArrivalEstimate []*schedule.ArrivalEstimate
	for _, sched := range martaScheds {
		x := schedule.ArrivalEstimate{
			PrimaryKey:     sched.PrimaryKey,
			SortKey:        sched.SortKey,
			Destination:    sched.Destination,
			Direction:      sched.Direction,
			EventTime:      sched.EventTime,
			Line:           sched.Line,
			NextArrival:    sched.NextArrival,
			Station:        sched.Station,
			TrainID:        sched.TrainID,
			WaitingSeconds: sched.WaitingSeconds,
			WaitingTime:    sched.WaitingTime,
			TTL:            sched.TTL,
		}
		protoArrivalEstimate = append(protoArrivalEstimate, &x)
	}
	return protoArrivalEstimate
}
