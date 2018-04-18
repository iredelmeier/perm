// Code generated by counterfeiter. DO NOT EDIT.
package genfakes

import (
	"sync"

	protos "code.cloudfoundry.org/perm/protos/gen"
	_ "github.com/gogo/protobuf/gogoproto"
	"golang.org/x/net/context"
)

type FakePermissionServiceServer struct {
	HasPermissionStub        func(context.Context, *protos.HasPermissionRequest) (*protos.HasPermissionResponse, error)
	hasPermissionMutex       sync.RWMutex
	hasPermissionArgsForCall []struct {
		arg1 context.Context
		arg2 *protos.HasPermissionRequest
	}
	hasPermissionReturns struct {
		result1 *protos.HasPermissionResponse
		result2 error
	}
	hasPermissionReturnsOnCall map[int]struct {
		result1 *protos.HasPermissionResponse
		result2 error
	}
	ListResourcePatternsStub        func(context.Context, *protos.ListResourcePatternsRequest) (*protos.ListResourcePatternsResponse, error)
	listResourcePatternsMutex       sync.RWMutex
	listResourcePatternsArgsForCall []struct {
		arg1 context.Context
		arg2 *protos.ListResourcePatternsRequest
	}
	listResourcePatternsReturns struct {
		result1 *protos.ListResourcePatternsResponse
		result2 error
	}
	listResourcePatternsReturnsOnCall map[int]struct {
		result1 *protos.ListResourcePatternsResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePermissionServiceServer) HasPermission(arg1 context.Context, arg2 *protos.HasPermissionRequest) (*protos.HasPermissionResponse, error) {
	fake.hasPermissionMutex.Lock()
	ret, specificReturn := fake.hasPermissionReturnsOnCall[len(fake.hasPermissionArgsForCall)]
	fake.hasPermissionArgsForCall = append(fake.hasPermissionArgsForCall, struct {
		arg1 context.Context
		arg2 *protos.HasPermissionRequest
	}{arg1, arg2})
	fake.recordInvocation("HasPermission", []interface{}{arg1, arg2})
	fake.hasPermissionMutex.Unlock()
	if fake.HasPermissionStub != nil {
		return fake.HasPermissionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.hasPermissionReturns.result1, fake.hasPermissionReturns.result2
}

func (fake *FakePermissionServiceServer) HasPermissionCallCount() int {
	fake.hasPermissionMutex.RLock()
	defer fake.hasPermissionMutex.RUnlock()
	return len(fake.hasPermissionArgsForCall)
}

func (fake *FakePermissionServiceServer) HasPermissionArgsForCall(i int) (context.Context, *protos.HasPermissionRequest) {
	fake.hasPermissionMutex.RLock()
	defer fake.hasPermissionMutex.RUnlock()
	return fake.hasPermissionArgsForCall[i].arg1, fake.hasPermissionArgsForCall[i].arg2
}

func (fake *FakePermissionServiceServer) HasPermissionReturns(result1 *protos.HasPermissionResponse, result2 error) {
	fake.HasPermissionStub = nil
	fake.hasPermissionReturns = struct {
		result1 *protos.HasPermissionResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePermissionServiceServer) HasPermissionReturnsOnCall(i int, result1 *protos.HasPermissionResponse, result2 error) {
	fake.HasPermissionStub = nil
	if fake.hasPermissionReturnsOnCall == nil {
		fake.hasPermissionReturnsOnCall = make(map[int]struct {
			result1 *protos.HasPermissionResponse
			result2 error
		})
	}
	fake.hasPermissionReturnsOnCall[i] = struct {
		result1 *protos.HasPermissionResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePermissionServiceServer) ListResourcePatterns(arg1 context.Context, arg2 *protos.ListResourcePatternsRequest) (*protos.ListResourcePatternsResponse, error) {
	fake.listResourcePatternsMutex.Lock()
	ret, specificReturn := fake.listResourcePatternsReturnsOnCall[len(fake.listResourcePatternsArgsForCall)]
	fake.listResourcePatternsArgsForCall = append(fake.listResourcePatternsArgsForCall, struct {
		arg1 context.Context
		arg2 *protos.ListResourcePatternsRequest
	}{arg1, arg2})
	fake.recordInvocation("ListResourcePatterns", []interface{}{arg1, arg2})
	fake.listResourcePatternsMutex.Unlock()
	if fake.ListResourcePatternsStub != nil {
		return fake.ListResourcePatternsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listResourcePatternsReturns.result1, fake.listResourcePatternsReturns.result2
}

func (fake *FakePermissionServiceServer) ListResourcePatternsCallCount() int {
	fake.listResourcePatternsMutex.RLock()
	defer fake.listResourcePatternsMutex.RUnlock()
	return len(fake.listResourcePatternsArgsForCall)
}

func (fake *FakePermissionServiceServer) ListResourcePatternsArgsForCall(i int) (context.Context, *protos.ListResourcePatternsRequest) {
	fake.listResourcePatternsMutex.RLock()
	defer fake.listResourcePatternsMutex.RUnlock()
	return fake.listResourcePatternsArgsForCall[i].arg1, fake.listResourcePatternsArgsForCall[i].arg2
}

func (fake *FakePermissionServiceServer) ListResourcePatternsReturns(result1 *protos.ListResourcePatternsResponse, result2 error) {
	fake.ListResourcePatternsStub = nil
	fake.listResourcePatternsReturns = struct {
		result1 *protos.ListResourcePatternsResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePermissionServiceServer) ListResourcePatternsReturnsOnCall(i int, result1 *protos.ListResourcePatternsResponse, result2 error) {
	fake.ListResourcePatternsStub = nil
	if fake.listResourcePatternsReturnsOnCall == nil {
		fake.listResourcePatternsReturnsOnCall = make(map[int]struct {
			result1 *protos.ListResourcePatternsResponse
			result2 error
		})
	}
	fake.listResourcePatternsReturnsOnCall[i] = struct {
		result1 *protos.ListResourcePatternsResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePermissionServiceServer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.hasPermissionMutex.RLock()
	defer fake.hasPermissionMutex.RUnlock()
	fake.listResourcePatternsMutex.RLock()
	defer fake.listResourcePatternsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePermissionServiceServer) recordInvocation(key string, args []interface{}) {
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

var _ protos.PermissionServiceServer = new(FakePermissionServiceServer)
