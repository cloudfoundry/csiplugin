// Code generated by counterfeiter. DO NOT EDIT.
package csipluginfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/csiplugin"
	"code.cloudfoundry.org/voldriver"
)

type FakeBackgroundInvoker struct {
	InvokeStub        func(env voldriver.Env, executable string, cmdArgs []string, waitFor string, timeout time.Duration) error
	invokeMutex       sync.RWMutex
	invokeArgsForCall []struct {
		env        voldriver.Env
		executable string
		cmdArgs    []string
		waitFor    string
		timeout    time.Duration
	}
	invokeReturns struct {
		result1 error
	}
	invokeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBackgroundInvoker) Invoke(env voldriver.Env, executable string, cmdArgs []string, waitFor string, timeout time.Duration) error {
	var cmdArgsCopy []string
	if cmdArgs != nil {
		cmdArgsCopy = make([]string, len(cmdArgs))
		copy(cmdArgsCopy, cmdArgs)
	}
	fake.invokeMutex.Lock()
	ret, specificReturn := fake.invokeReturnsOnCall[len(fake.invokeArgsForCall)]
	fake.invokeArgsForCall = append(fake.invokeArgsForCall, struct {
		env        voldriver.Env
		executable string
		cmdArgs    []string
		waitFor    string
		timeout    time.Duration
	}{env, executable, cmdArgsCopy, waitFor, timeout})
	fake.recordInvocation("Invoke", []interface{}{env, executable, cmdArgsCopy, waitFor, timeout})
	fake.invokeMutex.Unlock()
	if fake.InvokeStub != nil {
		return fake.InvokeStub(env, executable, cmdArgs, waitFor, timeout)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.invokeReturns.result1
}

func (fake *FakeBackgroundInvoker) InvokeCallCount() int {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return len(fake.invokeArgsForCall)
}

func (fake *FakeBackgroundInvoker) InvokeArgsForCall(i int) (voldriver.Env, string, []string, string, time.Duration) {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return fake.invokeArgsForCall[i].env, fake.invokeArgsForCall[i].executable, fake.invokeArgsForCall[i].cmdArgs, fake.invokeArgsForCall[i].waitFor, fake.invokeArgsForCall[i].timeout
}

func (fake *FakeBackgroundInvoker) InvokeReturns(result1 error) {
	fake.InvokeStub = nil
	fake.invokeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackgroundInvoker) InvokeReturnsOnCall(i int, result1 error) {
	fake.InvokeStub = nil
	if fake.invokeReturnsOnCall == nil {
		fake.invokeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.invokeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackgroundInvoker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBackgroundInvoker) recordInvocation(key string, args []interface{}) {
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

var _ csiplugin.BackgroundInvoker = new(FakeBackgroundInvoker)
