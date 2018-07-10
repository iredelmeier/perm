// Code generated by counterfeiter. DO NOT EDIT.
package logxfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/perm/pkg/logx"
)

type FakeSecurityLogger struct {
	LogStub        func(ctx context.Context, signature, name string, args ...logx.SecurityData)
	logMutex       sync.RWMutex
	logArgsForCall []struct {
		ctx       context.Context
		signature string
		name      string
		args      []logx.SecurityData
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecurityLogger) Log(ctx context.Context, signature string, name string, args ...logx.SecurityData) {
	fake.logMutex.Lock()
	fake.logArgsForCall = append(fake.logArgsForCall, struct {
		ctx       context.Context
		signature string
		name      string
		args      []logx.SecurityData
	}{ctx, signature, name, args})
	fake.recordInvocation("Log", []interface{}{ctx, signature, name, args})
	fake.logMutex.Unlock()
	if fake.LogStub != nil {
		fake.LogStub(ctx, signature, name, args...)
	}
}

func (fake *FakeSecurityLogger) LogCallCount() int {
	fake.logMutex.RLock()
	defer fake.logMutex.RUnlock()
	return len(fake.logArgsForCall)
}

func (fake *FakeSecurityLogger) LogArgsForCall(i int) (context.Context, string, string, []logx.SecurityData) {
	fake.logMutex.RLock()
	defer fake.logMutex.RUnlock()
	return fake.logArgsForCall[i].ctx, fake.logArgsForCall[i].signature, fake.logArgsForCall[i].name, fake.logArgsForCall[i].args
}

func (fake *FakeSecurityLogger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.logMutex.RLock()
	defer fake.logMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecurityLogger) recordInvocation(key string, args []interface{}) {
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

var _ logx.SecurityLogger = new(FakeSecurityLogger)
