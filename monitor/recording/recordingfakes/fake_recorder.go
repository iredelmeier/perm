// Code generated by counterfeiter. DO NOT EDIT.
package recordingfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/perm/monitor/recording"
)

type FakeRecorder struct {
	ObserveStub        func(duration time.Duration) error
	observeMutex       sync.RWMutex
	observeArgsForCall []struct {
		duration time.Duration
	}
	observeReturns struct {
		result1 error
	}
	observeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRecorder) Observe(duration time.Duration) error {
	fake.observeMutex.Lock()
	ret, specificReturn := fake.observeReturnsOnCall[len(fake.observeArgsForCall)]
	fake.observeArgsForCall = append(fake.observeArgsForCall, struct {
		duration time.Duration
	}{duration})
	fake.recordInvocation("Observe", []interface{}{duration})
	fake.observeMutex.Unlock()
	if fake.ObserveStub != nil {
		return fake.ObserveStub(duration)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.observeReturns.result1
}

func (fake *FakeRecorder) ObserveCallCount() int {
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
	return len(fake.observeArgsForCall)
}

func (fake *FakeRecorder) ObserveArgsForCall(i int) time.Duration {
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
	return fake.observeArgsForCall[i].duration
}

func (fake *FakeRecorder) ObserveReturns(result1 error) {
	fake.ObserveStub = nil
	fake.observeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRecorder) ObserveReturnsOnCall(i int, result1 error) {
	fake.ObserveStub = nil
	if fake.observeReturnsOnCall == nil {
		fake.observeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.observeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRecorder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRecorder) recordInvocation(key string, args []interface{}) {
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

var _ recording.Recorder = new(FakeRecorder)
