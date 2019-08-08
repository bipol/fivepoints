package handler_test

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/smartatransit/fivepoints/api/v1/schedule"
	"github.com/smartatransit/fivepoints/cmd/scheduleapi/handler"
	"github.com/smartatransit/fivepoints/cmd/scheduleapi/handler/handlerfakes"
)

var _ = Describe("GetArrivalEstimates", func() {
	Context("NewGetArrivalEstimatesEndpoint", func() {
		var (
			request    *schedule.GetArrivalEstimatesRequest
			response   *schedule.GetArrivalEstimatesResponse
			ctx        context.Context
			endpoint   handler.GetArrivalEstimatesEndpoint
			querier    *handlerfakes.FakeDynamoQuerier
			authorizer *handlerfakes.FakeAuthorizer
			err        error
			t          *timestamp.Timestamp
		)
		BeforeEach(func() {
			err = nil
			t = ptypes.TimestampNow()
			ctx = context.Background()
			request = &schedule.GetArrivalEstimatesRequest{
				StartDate:   t,
				EndDate:     t,
				Station:     "station",
				Destination: "direction",
			}
			response = nil
			querier = &handlerfakes.FakeDynamoQuerier{}
			authorizer = &handlerfakes.FakeAuthorizer{}
			authorizer.IsAuthorizedReturns(true, nil)
			endpoint = handler.NewGetArrivalEstimatesEndpoint("tableName", querier, authorizer)
		})
		JustBeforeEach(func() {
			response, err = endpoint(ctx, request)
		})
		When("everything goes well", func() {
			BeforeEach(func() {
				querier.QueryWithContextReturns(&handler.DynamoJSON, nil)
			})
			It("should not return an error", func() {
				Expect(len(response.ArrivalEstimates)).To(Equal(3))
				Expect(err).To(BeNil())
			})
		})
	})
	Context("ValidateRequest", func() {
		var (
			t   *timestamp.Timestamp
			n   *timestamp.Timestamp
			in  *schedule.GetArrivalEstimatesRequest
			err error
		)
		BeforeEach(func() {
			t = ptypes.TimestampNow()
			n, _ = ptypes.TimestampProto(time.Now().Add(24 * 60 * time.Minute))
			in = &schedule.GetArrivalEstimatesRequest{
				StartDate:   t,
				EndDate:     n,
				Station:     "North Avenue Station",
				Destination: "North Springs",
			}
			err = nil
		})
		JustBeforeEach(func() {
			err = handler.ValidateRequest(context.Background(), in)
		})
		When("missing direction", func() {
			BeforeEach(func() {
				in = &schedule.GetArrivalEstimatesRequest{
					StartDate: t,
					EndDate:   n,
					Station:   "North Avenue Station",
				}
			})
			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("missing station", func() {
			BeforeEach(func() {
				in = &schedule.GetArrivalEstimatesRequest{
					StartDate:   t,
					EndDate:     n,
					Destination: "North Springs",
				}
			})
			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("missing start date", func() {
			BeforeEach(func() {
				in = &schedule.GetArrivalEstimatesRequest{
					EndDate:     n,
					Station:     "North Avenue Station",
					Destination: "North Springs",
				}
			})
			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("missing end date", func() {
			BeforeEach(func() {
				in = &schedule.GetArrivalEstimatesRequest{
					StartDate:   t,
					Station:     "North Avenue Station",
					Destination: "North Springs",
				}
			})
			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("everything is great", func() {
			It("should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})
	Context("GetArrivalEstimatesRequestToDynamoQuery", func() {
		var (
			in        *schedule.GetArrivalEstimatesRequest
			tableName string
			queryIn   *dynamodb.QueryInput
			err       error
			t         *timestamp.Timestamp
			pKey      string
			currTime  time.Time
		)
		BeforeEach(func() {
			tableName = "table"
			t = ptypes.TimestampNow()
			in = &schedule.GetArrivalEstimatesRequest{
				StartDate:        t,
				EndDate:          t,
				Station:          "North Avenue Station",
				Destination:      "North Springs",
				LastEvaluatedKey: "bnVsbA==",
			}
			currTime, err = ptypes.Timestamp(t)
			if err != nil {
				Expect(err).To(BeNil())
			}

			pKey = fmt.Sprintf("%s_%s", in.GetStation(), in.GetDestination())
		})
		JustBeforeEach(func() {
			queryIn, err = handler.GetArrivalEstimatesRequestToDynamoQuery(in, tableName)
		})
		When("start date not set", func() {
			BeforeEach(func() {
				in = &schedule.GetArrivalEstimatesRequest{
					EndDate:     t,
					Station:     "North Avenue Station",
					Destination: "North Springs",
				}
			})
			It("does return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("end date not set", func() {
			BeforeEach(func() {
				in = &schedule.GetArrivalEstimatesRequest{
					StartDate:   t,
					Station:     "North Avenue Station",
					Destination: "North Springs",
				}
			})
			It("does return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("everything works", func() {
			It("returns the correct query input", func() {
				Expect(queryIn).To(PointTo(MatchFields(IgnoreExtras, Fields{
					"ExpressionAttributeNames": MatchAllKeys(Keys{
						"#0": PointTo(Equal("PrimaryKey")),
						"#1": PointTo(Equal("SortKey")),
					}),
					"ExpressionAttributeValues": MatchAllKeys(Keys{
						":2": PointTo(MatchFields(IgnoreExtras, Fields{
							"S": PointTo(Equal(currTime.Format(time.RFC3339))),
						})),
						":0": PointTo(MatchFields(IgnoreExtras, Fields{
							"S": PointTo(Equal(pKey)),
						})),
						":1": PointTo(MatchFields(IgnoreExtras, Fields{
							"S": PointTo(Equal(currTime.Format(time.RFC3339))),
						})),
					}),
				})))
			})
			It("does not return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

})
