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

var _ = Describe("GetSchedule", func() {
	Context("NewGetScheduleEndpoint", func() {
		var (
			request  *schedule.GetScheduleRequest
			response *schedule.GetScheduleResponse
			ctx      context.Context
			endpoint handler.GetScheduleEndpoint
			querier  *handlerfakes.FakeDynamoQuerier
			err      error
			t        *timestamp.Timestamp
		)
		BeforeEach(func() {
			err = nil
			t = ptypes.TimestampNow()
			ctx = context.Background()
			request = &schedule.GetScheduleRequest{
				StartDate: t,
				EndDate:   t,
				Station:   "station",
				Direction: "direction",
			}
			response = nil
			querier = &handlerfakes.FakeDynamoQuerier{}
			endpoint = handler.NewGetScheduleEndpoint("tableName", querier)
		})
		JustBeforeEach(func() {
			response, err = endpoint(ctx, request)
		})
		When("everything goes well", func() {
			BeforeEach(func() {
				querier.QueryWithContextReturns(&dynamodb.QueryOutput{}, nil)
			})
			It("should not return an error", func() {
				Expect(len(response.Schedules)).To(BeZero())
				Expect(err).To(BeNil())
			})
		})
	})
	Context("ValidateRequest", func() {
		var (
			t   *timestamp.Timestamp
			in  *schedule.GetScheduleRequest
			err error
		)
		BeforeEach(func() {
			t = ptypes.TimestampNow()
			in = &schedule.GetScheduleRequest{
				StartDate: t,
				EndDate:   t,
				Station:   "North Avenue Station",
				Direction: "North Springs",
			}
			err = nil
		})
		JustBeforeEach(func() {
			err = handler.ValidateRequest(context.Background(), in)
		})
		When("missing direction", func() {
			BeforeEach(func() {
				in = &schedule.GetScheduleRequest{
					StartDate: t,
					EndDate:   t,
					Station:   "North Avenue Station",
				}
			})
			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("missing station", func() {
			BeforeEach(func() {
				in = &schedule.GetScheduleRequest{
					StartDate: t,
					EndDate:   t,
					Direction: "North Springs",
				}
			})
			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("missing start date", func() {
			BeforeEach(func() {
				in = &schedule.GetScheduleRequest{
					EndDate:   t,
					Station:   "North Avenue Station",
					Direction: "North Springs",
				}
			})
			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("missing end date", func() {
			BeforeEach(func() {
				in = &schedule.GetScheduleRequest{
					StartDate: t,
					Station:   "North Avenue Station",
					Direction: "North Springs",
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
	Context("GetScheduleRequestToDynamoQuery", func() {
		var (
			in        *schedule.GetScheduleRequest
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
			in = &schedule.GetScheduleRequest{
				StartDate: t,
				EndDate:   t,
				Station:   "North Avenue Station",
				Direction: "North Springs",
			}
			currTime, err = ptypes.Timestamp(t)
			if err != nil {
				Expect(err).To(BeNil())
			}

			pKey = fmt.Sprintf("%s_%s_%s", in.GetStation(), in.GetDirection(), currTime.Format("2006-01-02"))
		})
		JustBeforeEach(func() {
			queryIn, err = handler.GetScheduleRequestToDynamoQuery(in, tableName)
		})
		When("start date not set", func() {
			BeforeEach(func() {
				in = &schedule.GetScheduleRequest{
					EndDate:   t,
					Station:   "North Avenue Station",
					Direction: "North Springs",
				}
			})
			It("does return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("end date not set", func() {
			BeforeEach(func() {
				in = &schedule.GetScheduleRequest{
					StartDate: t,
					Station:   "North Avenue Station",
					Direction: "North Springs",
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
							"S": PointTo(Equal(currTime.Format(time.RFC3339Nano))),
						})),
						":0": PointTo(MatchFields(IgnoreExtras, Fields{
							"S": PointTo(Equal(pKey)),
						})),
						":1": PointTo(MatchFields(IgnoreExtras, Fields{
							"S": PointTo(Equal(currTime.Format(time.RFC3339Nano))),
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
