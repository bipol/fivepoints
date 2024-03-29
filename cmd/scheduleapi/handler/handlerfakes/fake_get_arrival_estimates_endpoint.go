// Code generated by counterfeiter. DO NOT EDIT.
package handlerfakes

import (
	"context"
	"sync"

	"github.com/smartatransit/fivepoints/api/v1/schedule"
	"github.com/smartatransit/fivepoints/cmd/scheduleapi/handler"
)

type FakeGetArrivalEstimatesEndpoint struct {
	Stub        func(context.Context, *schedule.GetArrivalEstimatesRequest) (*schedule.GetArrivalEstimatesResponse, error)
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 context.Context
		arg2 *schedule.GetArrivalEstimatesRequest
	}
	returns struct {
		result1 *schedule.GetArrivalEstimatesResponse
		result2 error
	}
	returnsOnCall map[int]struct {
		result1 *schedule.GetArrivalEstimatesResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGetArrivalEstimatesEndpoint) Spy(arg1 context.Context, arg2 *schedule.GetArrivalEstimatesRequest) (*schedule.GetArrivalEstimatesResponse, error) {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 context.Context
		arg2 *schedule.GetArrivalEstimatesRequest
	}{arg1, arg2})
	fake.recordInvocation("GetArrivalEstimatesEndpoint", []interface{}{arg1, arg2})
	fake.mutex.Unlock()
	if fake.Stub != nil {
		return fake.Stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.returns.result1, fake.returns.result2
}

func (fake *FakeGetArrivalEstimatesEndpoint) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FakeGetArrivalEstimatesEndpoint) Calls(stub func(context.Context, *schedule.GetArrivalEstimatesRequest) (*schedule.GetArrivalEstimatesResponse, error)) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *FakeGetArrivalEstimatesEndpoint) ArgsForCall(i int) (context.Context, *schedule.GetArrivalEstimatesRequest) {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1, fake.argsForCall[i].arg2
}

func (fake *FakeGetArrivalEstimatesEndpoint) Returns(result1 *schedule.GetArrivalEstimatesResponse, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 *schedule.GetArrivalEstimatesResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeGetArrivalEstimatesEndpoint) ReturnsOnCall(i int, result1 *schedule.GetArrivalEstimatesResponse, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 *schedule.GetArrivalEstimatesResponse
			result2 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 *schedule.GetArrivalEstimatesResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeGetArrivalEstimatesEndpoint) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGetArrivalEstimatesEndpoint) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ handler.GetArrivalEstimatesEndpoint = new(FakeGetArrivalEstimatesEndpoint).Spy
