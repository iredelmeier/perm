// Code generated by counterfeiter. DO NOT EDIT.
package monitorfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/perm/pkg/monitor"
	"github.com/cactus/go-statsd-client/statsd"
)

type FakePermStatter struct {
	IncStub        func(string, int64, float32) error
	incMutex       sync.RWMutex
	incArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 float32
	}
	incReturns struct {
		result1 error
	}
	incReturnsOnCall map[int]struct {
		result1 error
	}
	DecStub        func(string, int64, float32) error
	decMutex       sync.RWMutex
	decArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 float32
	}
	decReturns struct {
		result1 error
	}
	decReturnsOnCall map[int]struct {
		result1 error
	}
	GaugeStub        func(string, int64, float32) error
	gaugeMutex       sync.RWMutex
	gaugeArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 float32
	}
	gaugeReturns struct {
		result1 error
	}
	gaugeReturnsOnCall map[int]struct {
		result1 error
	}
	GaugeDeltaStub        func(string, int64, float32) error
	gaugeDeltaMutex       sync.RWMutex
	gaugeDeltaArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 float32
	}
	gaugeDeltaReturns struct {
		result1 error
	}
	gaugeDeltaReturnsOnCall map[int]struct {
		result1 error
	}
	TimingStub        func(string, int64, float32) error
	timingMutex       sync.RWMutex
	timingArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 float32
	}
	timingReturns struct {
		result1 error
	}
	timingReturnsOnCall map[int]struct {
		result1 error
	}
	TimingDurationStub        func(string, time.Duration, float32) error
	timingDurationMutex       sync.RWMutex
	timingDurationArgsForCall []struct {
		arg1 string
		arg2 time.Duration
		arg3 float32
	}
	timingDurationReturns struct {
		result1 error
	}
	timingDurationReturnsOnCall map[int]struct {
		result1 error
	}
	SetStub        func(string, string, float32) error
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 float32
	}
	setReturns struct {
		result1 error
	}
	setReturnsOnCall map[int]struct {
		result1 error
	}
	SetIntStub        func(string, int64, float32) error
	setIntMutex       sync.RWMutex
	setIntArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 float32
	}
	setIntReturns struct {
		result1 error
	}
	setIntReturnsOnCall map[int]struct {
		result1 error
	}
	RawStub        func(string, string, float32) error
	rawMutex       sync.RWMutex
	rawArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 float32
	}
	rawReturns struct {
		result1 error
	}
	rawReturnsOnCall map[int]struct {
		result1 error
	}
	NewSubStatterStub        func(string) statsd.SubStatter
	newSubStatterMutex       sync.RWMutex
	newSubStatterArgsForCall []struct {
		arg1 string
	}
	newSubStatterReturns struct {
		result1 statsd.SubStatter
	}
	newSubStatterReturnsOnCall map[int]struct {
		result1 statsd.SubStatter
	}
	SetPrefixStub        func(string)
	setPrefixMutex       sync.RWMutex
	setPrefixArgsForCall []struct {
		arg1 string
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	RotateStub                     func()
	rotateMutex                    sync.RWMutex
	rotateArgsForCall              []struct{}
	RecordProbeDurationStub        func(logger lager.Logger, d time.Duration)
	recordProbeDurationMutex       sync.RWMutex
	recordProbeDurationArgsForCall []struct {
		logger lager.Logger
		d      time.Duration
	}
	SendFailedProbeStub        func(logger lager.Logger)
	sendFailedProbeMutex       sync.RWMutex
	sendFailedProbeArgsForCall []struct {
		logger lager.Logger
	}
	SendIncorrectProbeStub        func(logger lager.Logger)
	sendIncorrectProbeMutex       sync.RWMutex
	sendIncorrectProbeArgsForCall []struct {
		logger lager.Logger
	}
	SendCorrectProbeStub        func(logger lager.Logger)
	sendCorrectProbeMutex       sync.RWMutex
	sendCorrectProbeArgsForCall []struct {
		logger lager.Logger
	}
	SendStatsStub        func(logger lager.Logger)
	sendStatsMutex       sync.RWMutex
	sendStatsArgsForCall []struct {
		logger lager.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePermStatter) Inc(arg1 string, arg2 int64, arg3 float32) error {
	fake.incMutex.Lock()
	ret, specificReturn := fake.incReturnsOnCall[len(fake.incArgsForCall)]
	fake.incArgsForCall = append(fake.incArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 float32
	}{arg1, arg2, arg3})
	fake.recordInvocation("Inc", []interface{}{arg1, arg2, arg3})
	fake.incMutex.Unlock()
	if fake.IncStub != nil {
		return fake.IncStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.incReturns.result1
}

func (fake *FakePermStatter) IncCallCount() int {
	fake.incMutex.RLock()
	defer fake.incMutex.RUnlock()
	return len(fake.incArgsForCall)
}

func (fake *FakePermStatter) IncArgsForCall(i int) (string, int64, float32) {
	fake.incMutex.RLock()
	defer fake.incMutex.RUnlock()
	return fake.incArgsForCall[i].arg1, fake.incArgsForCall[i].arg2, fake.incArgsForCall[i].arg3
}

func (fake *FakePermStatter) IncReturns(result1 error) {
	fake.IncStub = nil
	fake.incReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) IncReturnsOnCall(i int, result1 error) {
	fake.IncStub = nil
	if fake.incReturnsOnCall == nil {
		fake.incReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.incReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) Dec(arg1 string, arg2 int64, arg3 float32) error {
	fake.decMutex.Lock()
	ret, specificReturn := fake.decReturnsOnCall[len(fake.decArgsForCall)]
	fake.decArgsForCall = append(fake.decArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 float32
	}{arg1, arg2, arg3})
	fake.recordInvocation("Dec", []interface{}{arg1, arg2, arg3})
	fake.decMutex.Unlock()
	if fake.DecStub != nil {
		return fake.DecStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.decReturns.result1
}

func (fake *FakePermStatter) DecCallCount() int {
	fake.decMutex.RLock()
	defer fake.decMutex.RUnlock()
	return len(fake.decArgsForCall)
}

func (fake *FakePermStatter) DecArgsForCall(i int) (string, int64, float32) {
	fake.decMutex.RLock()
	defer fake.decMutex.RUnlock()
	return fake.decArgsForCall[i].arg1, fake.decArgsForCall[i].arg2, fake.decArgsForCall[i].arg3
}

func (fake *FakePermStatter) DecReturns(result1 error) {
	fake.DecStub = nil
	fake.decReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) DecReturnsOnCall(i int, result1 error) {
	fake.DecStub = nil
	if fake.decReturnsOnCall == nil {
		fake.decReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.decReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) Gauge(arg1 string, arg2 int64, arg3 float32) error {
	fake.gaugeMutex.Lock()
	ret, specificReturn := fake.gaugeReturnsOnCall[len(fake.gaugeArgsForCall)]
	fake.gaugeArgsForCall = append(fake.gaugeArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 float32
	}{arg1, arg2, arg3})
	fake.recordInvocation("Gauge", []interface{}{arg1, arg2, arg3})
	fake.gaugeMutex.Unlock()
	if fake.GaugeStub != nil {
		return fake.GaugeStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.gaugeReturns.result1
}

func (fake *FakePermStatter) GaugeCallCount() int {
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	return len(fake.gaugeArgsForCall)
}

func (fake *FakePermStatter) GaugeArgsForCall(i int) (string, int64, float32) {
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	return fake.gaugeArgsForCall[i].arg1, fake.gaugeArgsForCall[i].arg2, fake.gaugeArgsForCall[i].arg3
}

func (fake *FakePermStatter) GaugeReturns(result1 error) {
	fake.GaugeStub = nil
	fake.gaugeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) GaugeReturnsOnCall(i int, result1 error) {
	fake.GaugeStub = nil
	if fake.gaugeReturnsOnCall == nil {
		fake.gaugeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.gaugeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) GaugeDelta(arg1 string, arg2 int64, arg3 float32) error {
	fake.gaugeDeltaMutex.Lock()
	ret, specificReturn := fake.gaugeDeltaReturnsOnCall[len(fake.gaugeDeltaArgsForCall)]
	fake.gaugeDeltaArgsForCall = append(fake.gaugeDeltaArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 float32
	}{arg1, arg2, arg3})
	fake.recordInvocation("GaugeDelta", []interface{}{arg1, arg2, arg3})
	fake.gaugeDeltaMutex.Unlock()
	if fake.GaugeDeltaStub != nil {
		return fake.GaugeDeltaStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.gaugeDeltaReturns.result1
}

func (fake *FakePermStatter) GaugeDeltaCallCount() int {
	fake.gaugeDeltaMutex.RLock()
	defer fake.gaugeDeltaMutex.RUnlock()
	return len(fake.gaugeDeltaArgsForCall)
}

func (fake *FakePermStatter) GaugeDeltaArgsForCall(i int) (string, int64, float32) {
	fake.gaugeDeltaMutex.RLock()
	defer fake.gaugeDeltaMutex.RUnlock()
	return fake.gaugeDeltaArgsForCall[i].arg1, fake.gaugeDeltaArgsForCall[i].arg2, fake.gaugeDeltaArgsForCall[i].arg3
}

func (fake *FakePermStatter) GaugeDeltaReturns(result1 error) {
	fake.GaugeDeltaStub = nil
	fake.gaugeDeltaReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) GaugeDeltaReturnsOnCall(i int, result1 error) {
	fake.GaugeDeltaStub = nil
	if fake.gaugeDeltaReturnsOnCall == nil {
		fake.gaugeDeltaReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.gaugeDeltaReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) Timing(arg1 string, arg2 int64, arg3 float32) error {
	fake.timingMutex.Lock()
	ret, specificReturn := fake.timingReturnsOnCall[len(fake.timingArgsForCall)]
	fake.timingArgsForCall = append(fake.timingArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 float32
	}{arg1, arg2, arg3})
	fake.recordInvocation("Timing", []interface{}{arg1, arg2, arg3})
	fake.timingMutex.Unlock()
	if fake.TimingStub != nil {
		return fake.TimingStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.timingReturns.result1
}

func (fake *FakePermStatter) TimingCallCount() int {
	fake.timingMutex.RLock()
	defer fake.timingMutex.RUnlock()
	return len(fake.timingArgsForCall)
}

func (fake *FakePermStatter) TimingArgsForCall(i int) (string, int64, float32) {
	fake.timingMutex.RLock()
	defer fake.timingMutex.RUnlock()
	return fake.timingArgsForCall[i].arg1, fake.timingArgsForCall[i].arg2, fake.timingArgsForCall[i].arg3
}

func (fake *FakePermStatter) TimingReturns(result1 error) {
	fake.TimingStub = nil
	fake.timingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) TimingReturnsOnCall(i int, result1 error) {
	fake.TimingStub = nil
	if fake.timingReturnsOnCall == nil {
		fake.timingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.timingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) TimingDuration(arg1 string, arg2 time.Duration, arg3 float32) error {
	fake.timingDurationMutex.Lock()
	ret, specificReturn := fake.timingDurationReturnsOnCall[len(fake.timingDurationArgsForCall)]
	fake.timingDurationArgsForCall = append(fake.timingDurationArgsForCall, struct {
		arg1 string
		arg2 time.Duration
		arg3 float32
	}{arg1, arg2, arg3})
	fake.recordInvocation("TimingDuration", []interface{}{arg1, arg2, arg3})
	fake.timingDurationMutex.Unlock()
	if fake.TimingDurationStub != nil {
		return fake.TimingDurationStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.timingDurationReturns.result1
}

func (fake *FakePermStatter) TimingDurationCallCount() int {
	fake.timingDurationMutex.RLock()
	defer fake.timingDurationMutex.RUnlock()
	return len(fake.timingDurationArgsForCall)
}

func (fake *FakePermStatter) TimingDurationArgsForCall(i int) (string, time.Duration, float32) {
	fake.timingDurationMutex.RLock()
	defer fake.timingDurationMutex.RUnlock()
	return fake.timingDurationArgsForCall[i].arg1, fake.timingDurationArgsForCall[i].arg2, fake.timingDurationArgsForCall[i].arg3
}

func (fake *FakePermStatter) TimingDurationReturns(result1 error) {
	fake.TimingDurationStub = nil
	fake.timingDurationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) TimingDurationReturnsOnCall(i int, result1 error) {
	fake.TimingDurationStub = nil
	if fake.timingDurationReturnsOnCall == nil {
		fake.timingDurationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.timingDurationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) Set(arg1 string, arg2 string, arg3 float32) error {
	fake.setMutex.Lock()
	ret, specificReturn := fake.setReturnsOnCall[len(fake.setArgsForCall)]
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 float32
	}{arg1, arg2, arg3})
	fake.recordInvocation("Set", []interface{}{arg1, arg2, arg3})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		return fake.SetStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setReturns.result1
}

func (fake *FakePermStatter) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakePermStatter) SetArgsForCall(i int) (string, string, float32) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return fake.setArgsForCall[i].arg1, fake.setArgsForCall[i].arg2, fake.setArgsForCall[i].arg3
}

func (fake *FakePermStatter) SetReturns(result1 error) {
	fake.SetStub = nil
	fake.setReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) SetReturnsOnCall(i int, result1 error) {
	fake.SetStub = nil
	if fake.setReturnsOnCall == nil {
		fake.setReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) SetInt(arg1 string, arg2 int64, arg3 float32) error {
	fake.setIntMutex.Lock()
	ret, specificReturn := fake.setIntReturnsOnCall[len(fake.setIntArgsForCall)]
	fake.setIntArgsForCall = append(fake.setIntArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 float32
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetInt", []interface{}{arg1, arg2, arg3})
	fake.setIntMutex.Unlock()
	if fake.SetIntStub != nil {
		return fake.SetIntStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setIntReturns.result1
}

func (fake *FakePermStatter) SetIntCallCount() int {
	fake.setIntMutex.RLock()
	defer fake.setIntMutex.RUnlock()
	return len(fake.setIntArgsForCall)
}

func (fake *FakePermStatter) SetIntArgsForCall(i int) (string, int64, float32) {
	fake.setIntMutex.RLock()
	defer fake.setIntMutex.RUnlock()
	return fake.setIntArgsForCall[i].arg1, fake.setIntArgsForCall[i].arg2, fake.setIntArgsForCall[i].arg3
}

func (fake *FakePermStatter) SetIntReturns(result1 error) {
	fake.SetIntStub = nil
	fake.setIntReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) SetIntReturnsOnCall(i int, result1 error) {
	fake.SetIntStub = nil
	if fake.setIntReturnsOnCall == nil {
		fake.setIntReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setIntReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) Raw(arg1 string, arg2 string, arg3 float32) error {
	fake.rawMutex.Lock()
	ret, specificReturn := fake.rawReturnsOnCall[len(fake.rawArgsForCall)]
	fake.rawArgsForCall = append(fake.rawArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 float32
	}{arg1, arg2, arg3})
	fake.recordInvocation("Raw", []interface{}{arg1, arg2, arg3})
	fake.rawMutex.Unlock()
	if fake.RawStub != nil {
		return fake.RawStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.rawReturns.result1
}

func (fake *FakePermStatter) RawCallCount() int {
	fake.rawMutex.RLock()
	defer fake.rawMutex.RUnlock()
	return len(fake.rawArgsForCall)
}

func (fake *FakePermStatter) RawArgsForCall(i int) (string, string, float32) {
	fake.rawMutex.RLock()
	defer fake.rawMutex.RUnlock()
	return fake.rawArgsForCall[i].arg1, fake.rawArgsForCall[i].arg2, fake.rawArgsForCall[i].arg3
}

func (fake *FakePermStatter) RawReturns(result1 error) {
	fake.RawStub = nil
	fake.rawReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) RawReturnsOnCall(i int, result1 error) {
	fake.RawStub = nil
	if fake.rawReturnsOnCall == nil {
		fake.rawReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rawReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) NewSubStatter(arg1 string) statsd.SubStatter {
	fake.newSubStatterMutex.Lock()
	ret, specificReturn := fake.newSubStatterReturnsOnCall[len(fake.newSubStatterArgsForCall)]
	fake.newSubStatterArgsForCall = append(fake.newSubStatterArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("NewSubStatter", []interface{}{arg1})
	fake.newSubStatterMutex.Unlock()
	if fake.NewSubStatterStub != nil {
		return fake.NewSubStatterStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newSubStatterReturns.result1
}

func (fake *FakePermStatter) NewSubStatterCallCount() int {
	fake.newSubStatterMutex.RLock()
	defer fake.newSubStatterMutex.RUnlock()
	return len(fake.newSubStatterArgsForCall)
}

func (fake *FakePermStatter) NewSubStatterArgsForCall(i int) string {
	fake.newSubStatterMutex.RLock()
	defer fake.newSubStatterMutex.RUnlock()
	return fake.newSubStatterArgsForCall[i].arg1
}

func (fake *FakePermStatter) NewSubStatterReturns(result1 statsd.SubStatter) {
	fake.NewSubStatterStub = nil
	fake.newSubStatterReturns = struct {
		result1 statsd.SubStatter
	}{result1}
}

func (fake *FakePermStatter) NewSubStatterReturnsOnCall(i int, result1 statsd.SubStatter) {
	fake.NewSubStatterStub = nil
	if fake.newSubStatterReturnsOnCall == nil {
		fake.newSubStatterReturnsOnCall = make(map[int]struct {
			result1 statsd.SubStatter
		})
	}
	fake.newSubStatterReturnsOnCall[i] = struct {
		result1 statsd.SubStatter
	}{result1}
}

func (fake *FakePermStatter) SetPrefix(arg1 string) {
	fake.setPrefixMutex.Lock()
	fake.setPrefixArgsForCall = append(fake.setPrefixArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetPrefix", []interface{}{arg1})
	fake.setPrefixMutex.Unlock()
	if fake.SetPrefixStub != nil {
		fake.SetPrefixStub(arg1)
	}
}

func (fake *FakePermStatter) SetPrefixCallCount() int {
	fake.setPrefixMutex.RLock()
	defer fake.setPrefixMutex.RUnlock()
	return len(fake.setPrefixArgsForCall)
}

func (fake *FakePermStatter) SetPrefixArgsForCall(i int) string {
	fake.setPrefixMutex.RLock()
	defer fake.setPrefixMutex.RUnlock()
	return fake.setPrefixArgsForCall[i].arg1
}

func (fake *FakePermStatter) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *FakePermStatter) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakePermStatter) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePermStatter) Rotate() {
	fake.rotateMutex.Lock()
	fake.rotateArgsForCall = append(fake.rotateArgsForCall, struct{}{})
	fake.recordInvocation("Rotate", []interface{}{})
	fake.rotateMutex.Unlock()
	if fake.RotateStub != nil {
		fake.RotateStub()
	}
}

func (fake *FakePermStatter) RotateCallCount() int {
	fake.rotateMutex.RLock()
	defer fake.rotateMutex.RUnlock()
	return len(fake.rotateArgsForCall)
}

func (fake *FakePermStatter) RecordProbeDuration(logger lager.Logger, d time.Duration) {
	fake.recordProbeDurationMutex.Lock()
	fake.recordProbeDurationArgsForCall = append(fake.recordProbeDurationArgsForCall, struct {
		logger lager.Logger
		d      time.Duration
	}{logger, d})
	fake.recordInvocation("RecordProbeDuration", []interface{}{logger, d})
	fake.recordProbeDurationMutex.Unlock()
	if fake.RecordProbeDurationStub != nil {
		fake.RecordProbeDurationStub(logger, d)
	}
}

func (fake *FakePermStatter) RecordProbeDurationCallCount() int {
	fake.recordProbeDurationMutex.RLock()
	defer fake.recordProbeDurationMutex.RUnlock()
	return len(fake.recordProbeDurationArgsForCall)
}

func (fake *FakePermStatter) RecordProbeDurationArgsForCall(i int) (lager.Logger, time.Duration) {
	fake.recordProbeDurationMutex.RLock()
	defer fake.recordProbeDurationMutex.RUnlock()
	return fake.recordProbeDurationArgsForCall[i].logger, fake.recordProbeDurationArgsForCall[i].d
}

func (fake *FakePermStatter) SendFailedProbe(logger lager.Logger) {
	fake.sendFailedProbeMutex.Lock()
	fake.sendFailedProbeArgsForCall = append(fake.sendFailedProbeArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("SendFailedProbe", []interface{}{logger})
	fake.sendFailedProbeMutex.Unlock()
	if fake.SendFailedProbeStub != nil {
		fake.SendFailedProbeStub(logger)
	}
}

func (fake *FakePermStatter) SendFailedProbeCallCount() int {
	fake.sendFailedProbeMutex.RLock()
	defer fake.sendFailedProbeMutex.RUnlock()
	return len(fake.sendFailedProbeArgsForCall)
}

func (fake *FakePermStatter) SendFailedProbeArgsForCall(i int) lager.Logger {
	fake.sendFailedProbeMutex.RLock()
	defer fake.sendFailedProbeMutex.RUnlock()
	return fake.sendFailedProbeArgsForCall[i].logger
}

func (fake *FakePermStatter) SendIncorrectProbe(logger lager.Logger) {
	fake.sendIncorrectProbeMutex.Lock()
	fake.sendIncorrectProbeArgsForCall = append(fake.sendIncorrectProbeArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("SendIncorrectProbe", []interface{}{logger})
	fake.sendIncorrectProbeMutex.Unlock()
	if fake.SendIncorrectProbeStub != nil {
		fake.SendIncorrectProbeStub(logger)
	}
}

func (fake *FakePermStatter) SendIncorrectProbeCallCount() int {
	fake.sendIncorrectProbeMutex.RLock()
	defer fake.sendIncorrectProbeMutex.RUnlock()
	return len(fake.sendIncorrectProbeArgsForCall)
}

func (fake *FakePermStatter) SendIncorrectProbeArgsForCall(i int) lager.Logger {
	fake.sendIncorrectProbeMutex.RLock()
	defer fake.sendIncorrectProbeMutex.RUnlock()
	return fake.sendIncorrectProbeArgsForCall[i].logger
}

func (fake *FakePermStatter) SendCorrectProbe(logger lager.Logger) {
	fake.sendCorrectProbeMutex.Lock()
	fake.sendCorrectProbeArgsForCall = append(fake.sendCorrectProbeArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("SendCorrectProbe", []interface{}{logger})
	fake.sendCorrectProbeMutex.Unlock()
	if fake.SendCorrectProbeStub != nil {
		fake.SendCorrectProbeStub(logger)
	}
}

func (fake *FakePermStatter) SendCorrectProbeCallCount() int {
	fake.sendCorrectProbeMutex.RLock()
	defer fake.sendCorrectProbeMutex.RUnlock()
	return len(fake.sendCorrectProbeArgsForCall)
}

func (fake *FakePermStatter) SendCorrectProbeArgsForCall(i int) lager.Logger {
	fake.sendCorrectProbeMutex.RLock()
	defer fake.sendCorrectProbeMutex.RUnlock()
	return fake.sendCorrectProbeArgsForCall[i].logger
}

func (fake *FakePermStatter) SendStats(logger lager.Logger) {
	fake.sendStatsMutex.Lock()
	fake.sendStatsArgsForCall = append(fake.sendStatsArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("SendStats", []interface{}{logger})
	fake.sendStatsMutex.Unlock()
	if fake.SendStatsStub != nil {
		fake.SendStatsStub(logger)
	}
}

func (fake *FakePermStatter) SendStatsCallCount() int {
	fake.sendStatsMutex.RLock()
	defer fake.sendStatsMutex.RUnlock()
	return len(fake.sendStatsArgsForCall)
}

func (fake *FakePermStatter) SendStatsArgsForCall(i int) lager.Logger {
	fake.sendStatsMutex.RLock()
	defer fake.sendStatsMutex.RUnlock()
	return fake.sendStatsArgsForCall[i].logger
}

func (fake *FakePermStatter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.incMutex.RLock()
	defer fake.incMutex.RUnlock()
	fake.decMutex.RLock()
	defer fake.decMutex.RUnlock()
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	fake.gaugeDeltaMutex.RLock()
	defer fake.gaugeDeltaMutex.RUnlock()
	fake.timingMutex.RLock()
	defer fake.timingMutex.RUnlock()
	fake.timingDurationMutex.RLock()
	defer fake.timingDurationMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	fake.setIntMutex.RLock()
	defer fake.setIntMutex.RUnlock()
	fake.rawMutex.RLock()
	defer fake.rawMutex.RUnlock()
	fake.newSubStatterMutex.RLock()
	defer fake.newSubStatterMutex.RUnlock()
	fake.setPrefixMutex.RLock()
	defer fake.setPrefixMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.rotateMutex.RLock()
	defer fake.rotateMutex.RUnlock()
	fake.recordProbeDurationMutex.RLock()
	defer fake.recordProbeDurationMutex.RUnlock()
	fake.sendFailedProbeMutex.RLock()
	defer fake.sendFailedProbeMutex.RUnlock()
	fake.sendIncorrectProbeMutex.RLock()
	defer fake.sendIncorrectProbeMutex.RUnlock()
	fake.sendCorrectProbeMutex.RLock()
	defer fake.sendCorrectProbeMutex.RUnlock()
	fake.sendStatsMutex.RLock()
	defer fake.sendStatsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePermStatter) recordInvocation(key string, args []interface{}) {
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

var _ monitor.PermStatter = new(FakePermStatter)
