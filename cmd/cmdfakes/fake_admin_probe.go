// Code generated by counterfeiter. DO NOT EDIT.
package cmdfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/perm/cmd"
)

type FakeAdminProbe struct {
	CleanupStub        func(context.Context, lager.Logger, string) error
	cleanupMutex       sync.RWMutex
	cleanupArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}
	cleanupReturns struct {
		result1 error
	}
	cleanupReturnsOnCall map[int]struct {
		result1 error
	}
	RunStub        func(context.Context, lager.Logger, string) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}
	runReturns struct {
		result1 error
	}
	runReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAdminProbe) Cleanup(arg1 context.Context, arg2 lager.Logger, arg3 string) error {
	fake.cleanupMutex.Lock()
	ret, specificReturn := fake.cleanupReturnsOnCall[len(fake.cleanupArgsForCall)]
	fake.cleanupArgsForCall = append(fake.cleanupArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Cleanup", []interface{}{arg1, arg2, arg3})
	fake.cleanupMutex.Unlock()
	if fake.CleanupStub != nil {
		return fake.CleanupStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanupReturns.result1
}

func (fake *FakeAdminProbe) CleanupCallCount() int {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return len(fake.cleanupArgsForCall)
}

func (fake *FakeAdminProbe) CleanupArgsForCall(i int) (context.Context, lager.Logger, string) {
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	return fake.cleanupArgsForCall[i].arg1, fake.cleanupArgsForCall[i].arg2, fake.cleanupArgsForCall[i].arg3
}

func (fake *FakeAdminProbe) CleanupReturns(result1 error) {
	fake.CleanupStub = nil
	fake.cleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAdminProbe) CleanupReturnsOnCall(i int, result1 error) {
	fake.CleanupStub = nil
	if fake.cleanupReturnsOnCall == nil {
		fake.cleanupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAdminProbe) Run(arg1 context.Context, arg2 lager.Logger, arg3 string) error {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Run", []interface{}{arg1, arg2, arg3})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runReturns.result1
}

func (fake *FakeAdminProbe) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeAdminProbe) RunArgsForCall(i int) (context.Context, lager.Logger, string) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1, fake.runArgsForCall[i].arg2, fake.runArgsForCall[i].arg3
}

func (fake *FakeAdminProbe) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAdminProbe) RunReturnsOnCall(i int, result1 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAdminProbe) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cleanupMutex.RLock()
	defer fake.cleanupMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAdminProbe) recordInvocation(key string, args []interface{}) {
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

var _ cmd.AdminProbe = new(FakeAdminProbe)