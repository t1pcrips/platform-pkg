// Code generated by http://github.com/gojuno/minimock (v3.4.3). DO NOT EDIT.

package mocks

//go:generate minimock -i platform-pkg/pkg/database.NamedExecer -o named_execer_minimock.go -n NamedExecerMock -p mocks

import (
	"context"
	mm_database "platform-pkg/pkg/database"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// NamedExecerMock implements mm_database.NamedExecer
type NamedExecerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcScanAllContext          func(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{}) (err error)
	funcScanAllContextOrigin    string
	inspectFuncScanAllContext   func(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{})
	afterScanAllContextCounter  uint64
	beforeScanAllContextCounter uint64
	ScanAllContextMock          mNamedExecerMockScanAllContext

	funcScanOneContext          func(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{}) (err error)
	funcScanOneContextOrigin    string
	inspectFuncScanOneContext   func(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{})
	afterScanOneContextCounter  uint64
	beforeScanOneContextCounter uint64
	ScanOneContextMock          mNamedExecerMockScanOneContext
}

// NewNamedExecerMock returns a mock for mm_database.NamedExecer
func NewNamedExecerMock(t minimock.Tester) *NamedExecerMock {
	m := &NamedExecerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ScanAllContextMock = mNamedExecerMockScanAllContext{mock: m}
	m.ScanAllContextMock.callArgs = []*NamedExecerMockScanAllContextParams{}

	m.ScanOneContextMock = mNamedExecerMockScanOneContext{mock: m}
	m.ScanOneContextMock.callArgs = []*NamedExecerMockScanOneContextParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mNamedExecerMockScanAllContext struct {
	optional           bool
	mock               *NamedExecerMock
	defaultExpectation *NamedExecerMockScanAllContextExpectation
	expectations       []*NamedExecerMockScanAllContextExpectation

	callArgs []*NamedExecerMockScanAllContextParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// NamedExecerMockScanAllContextExpectation specifies expectation struct of the NamedExecer.ScanAllContext
type NamedExecerMockScanAllContextExpectation struct {
	mock               *NamedExecerMock
	params             *NamedExecerMockScanAllContextParams
	paramPtrs          *NamedExecerMockScanAllContextParamPtrs
	expectationOrigins NamedExecerMockScanAllContextExpectationOrigins
	results            *NamedExecerMockScanAllContextResults
	returnOrigin       string
	Counter            uint64
}

// NamedExecerMockScanAllContextParams contains parameters of the NamedExecer.ScanAllContext
type NamedExecerMockScanAllContextParams struct {
	ctx  context.Context
	deps interface{}
	q    mm_database.Query
	args []interface{}
}

// NamedExecerMockScanAllContextParamPtrs contains pointers to parameters of the NamedExecer.ScanAllContext
type NamedExecerMockScanAllContextParamPtrs struct {
	ctx  *context.Context
	deps *interface{}
	q    *mm_database.Query
	args *[]interface{}
}

// NamedExecerMockScanAllContextResults contains results of the NamedExecer.ScanAllContext
type NamedExecerMockScanAllContextResults struct {
	err error
}

// NamedExecerMockScanAllContextOrigins contains origins of expectations of the NamedExecer.ScanAllContext
type NamedExecerMockScanAllContextExpectationOrigins struct {
	origin     string
	originCtx  string
	originDeps string
	originQ    string
	originArgs string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmScanAllContext *mNamedExecerMockScanAllContext) Optional() *mNamedExecerMockScanAllContext {
	mmScanAllContext.optional = true
	return mmScanAllContext
}

// Expect sets up expected params for NamedExecer.ScanAllContext
func (mmScanAllContext *mNamedExecerMockScanAllContext) Expect(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{}) *mNamedExecerMockScanAllContext {
	if mmScanAllContext.mock.funcScanAllContext != nil {
		mmScanAllContext.mock.t.Fatalf("NamedExecerMock.ScanAllContext mock is already set by Set")
	}

	if mmScanAllContext.defaultExpectation == nil {
		mmScanAllContext.defaultExpectation = &NamedExecerMockScanAllContextExpectation{}
	}

	if mmScanAllContext.defaultExpectation.paramPtrs != nil {
		mmScanAllContext.mock.t.Fatalf("NamedExecerMock.ScanAllContext mock is already set by ExpectParams functions")
	}

	mmScanAllContext.defaultExpectation.params = &NamedExecerMockScanAllContextParams{ctx, deps, q, args}
	mmScanAllContext.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmScanAllContext.expectations {
		if minimock.Equal(e.params, mmScanAllContext.defaultExpectation.params) {
			mmScanAllContext.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmScanAllContext.defaultExpectation.params)
		}
	}

	return mmScanAllContext
}

// ExpectCtxParam1 sets up expected param ctx for NamedExecer.ScanAllContext
func (mmScanAllContext *mNamedExecerMockScanAllContext) ExpectCtxParam1(ctx context.Context) *mNamedExecerMockScanAllContext {
	if mmScanAllContext.mock.funcScanAllContext != nil {
		mmScanAllContext.mock.t.Fatalf("NamedExecerMock.ScanAllContext mock is already set by Set")
	}

	if mmScanAllContext.defaultExpectation == nil {
		mmScanAllContext.defaultExpectation = &NamedExecerMockScanAllContextExpectation{}
	}

	if mmScanAllContext.defaultExpectation.params != nil {
		mmScanAllContext.mock.t.Fatalf("NamedExecerMock.ScanAllContext mock is already set by Expect")
	}

	if mmScanAllContext.defaultExpectation.paramPtrs == nil {
		mmScanAllContext.defaultExpectation.paramPtrs = &NamedExecerMockScanAllContextParamPtrs{}
	}
	mmScanAllContext.defaultExpectation.paramPtrs.ctx = &ctx
	mmScanAllContext.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmScanAllContext
}

// ExpectDepsParam2 sets up expected param deps for NamedExecer.ScanAllContext
func (mmScanAllContext *mNamedExecerMockScanAllContext) ExpectDepsParam2(deps interface{}) *mNamedExecerMockScanAllContext {
	if mmScanAllContext.mock.funcScanAllContext != nil {
		mmScanAllContext.mock.t.Fatalf("NamedExecerMock.ScanAllContext mock is already set by Set")
	}

	if mmScanAllContext.defaultExpectation == nil {
		mmScanAllContext.defaultExpectation = &NamedExecerMockScanAllContextExpectation{}
	}

	if mmScanAllContext.defaultExpectation.params != nil {
		mmScanAllContext.mock.t.Fatalf("NamedExecerMock.ScanAllContext mock is already set by Expect")
	}

	if mmScanAllContext.defaultExpectation.paramPtrs == nil {
		mmScanAllContext.defaultExpectation.paramPtrs = &NamedExecerMockScanAllContextParamPtrs{}
	}
	mmScanAllContext.defaultExpectation.paramPtrs.deps = &deps
	mmScanAllContext.defaultExpectation.expectationOrigins.originDeps = minimock.CallerInfo(1)

	return mmScanAllContext
}

// ExpectQParam3 sets up expected param q for NamedExecer.ScanAllContext
func (mmScanAllContext *mNamedExecerMockScanAllContext) ExpectQParam3(q mm_database.Query) *mNamedExecerMockScanAllContext {
	if mmScanAllContext.mock.funcScanAllContext != nil {
		mmScanAllContext.mock.t.Fatalf("NamedExecerMock.ScanAllContext mock is already set by Set")
	}

	if mmScanAllContext.defaultExpectation == nil {
		mmScanAllContext.defaultExpectation = &NamedExecerMockScanAllContextExpectation{}
	}

	if mmScanAllContext.defaultExpectation.params != nil {
		mmScanAllContext.mock.t.Fatalf("NamedExecerMock.ScanAllContext mock is already set by Expect")
	}

	if mmScanAllContext.defaultExpectation.paramPtrs == nil {
		mmScanAllContext.defaultExpectation.paramPtrs = &NamedExecerMockScanAllContextParamPtrs{}
	}
	mmScanAllContext.defaultExpectation.paramPtrs.q = &q
	mmScanAllContext.defaultExpectation.expectationOrigins.originQ = minimock.CallerInfo(1)

	return mmScanAllContext
}

// ExpectArgsParam4 sets up expected param args for NamedExecer.ScanAllContext
func (mmScanAllContext *mNamedExecerMockScanAllContext) ExpectArgsParam4(args ...interface{}) *mNamedExecerMockScanAllContext {
	if mmScanAllContext.mock.funcScanAllContext != nil {
		mmScanAllContext.mock.t.Fatalf("NamedExecerMock.ScanAllContext mock is already set by Set")
	}

	if mmScanAllContext.defaultExpectation == nil {
		mmScanAllContext.defaultExpectation = &NamedExecerMockScanAllContextExpectation{}
	}

	if mmScanAllContext.defaultExpectation.params != nil {
		mmScanAllContext.mock.t.Fatalf("NamedExecerMock.ScanAllContext mock is already set by Expect")
	}

	if mmScanAllContext.defaultExpectation.paramPtrs == nil {
		mmScanAllContext.defaultExpectation.paramPtrs = &NamedExecerMockScanAllContextParamPtrs{}
	}
	mmScanAllContext.defaultExpectation.paramPtrs.args = &args
	mmScanAllContext.defaultExpectation.expectationOrigins.originArgs = minimock.CallerInfo(1)

	return mmScanAllContext
}

// Inspect accepts an inspector function that has same arguments as the NamedExecer.ScanAllContext
func (mmScanAllContext *mNamedExecerMockScanAllContext) Inspect(f func(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{})) *mNamedExecerMockScanAllContext {
	if mmScanAllContext.mock.inspectFuncScanAllContext != nil {
		mmScanAllContext.mock.t.Fatalf("Inspect function is already set for NamedExecerMock.ScanAllContext")
	}

	mmScanAllContext.mock.inspectFuncScanAllContext = f

	return mmScanAllContext
}

// Return sets up results that will be returned by NamedExecer.ScanAllContext
func (mmScanAllContext *mNamedExecerMockScanAllContext) Return(err error) *NamedExecerMock {
	if mmScanAllContext.mock.funcScanAllContext != nil {
		mmScanAllContext.mock.t.Fatalf("NamedExecerMock.ScanAllContext mock is already set by Set")
	}

	if mmScanAllContext.defaultExpectation == nil {
		mmScanAllContext.defaultExpectation = &NamedExecerMockScanAllContextExpectation{mock: mmScanAllContext.mock}
	}
	mmScanAllContext.defaultExpectation.results = &NamedExecerMockScanAllContextResults{err}
	mmScanAllContext.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmScanAllContext.mock
}

// Set uses given function f to mock the NamedExecer.ScanAllContext method
func (mmScanAllContext *mNamedExecerMockScanAllContext) Set(f func(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{}) (err error)) *NamedExecerMock {
	if mmScanAllContext.defaultExpectation != nil {
		mmScanAllContext.mock.t.Fatalf("Default expectation is already set for the NamedExecer.ScanAllContext method")
	}

	if len(mmScanAllContext.expectations) > 0 {
		mmScanAllContext.mock.t.Fatalf("Some expectations are already set for the NamedExecer.ScanAllContext method")
	}

	mmScanAllContext.mock.funcScanAllContext = f
	mmScanAllContext.mock.funcScanAllContextOrigin = minimock.CallerInfo(1)
	return mmScanAllContext.mock
}

// When sets expectation for the NamedExecer.ScanAllContext which will trigger the result defined by the following
// Then helper
func (mmScanAllContext *mNamedExecerMockScanAllContext) When(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{}) *NamedExecerMockScanAllContextExpectation {
	if mmScanAllContext.mock.funcScanAllContext != nil {
		mmScanAllContext.mock.t.Fatalf("NamedExecerMock.ScanAllContext mock is already set by Set")
	}

	expectation := &NamedExecerMockScanAllContextExpectation{
		mock:               mmScanAllContext.mock,
		params:             &NamedExecerMockScanAllContextParams{ctx, deps, q, args},
		expectationOrigins: NamedExecerMockScanAllContextExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmScanAllContext.expectations = append(mmScanAllContext.expectations, expectation)
	return expectation
}

// Then sets up NamedExecer.ScanAllContext return parameters for the expectation previously defined by the When method
func (e *NamedExecerMockScanAllContextExpectation) Then(err error) *NamedExecerMock {
	e.results = &NamedExecerMockScanAllContextResults{err}
	return e.mock
}

// Times sets number of times NamedExecer.ScanAllContext should be invoked
func (mmScanAllContext *mNamedExecerMockScanAllContext) Times(n uint64) *mNamedExecerMockScanAllContext {
	if n == 0 {
		mmScanAllContext.mock.t.Fatalf("Times of NamedExecerMock.ScanAllContext mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmScanAllContext.expectedInvocations, n)
	mmScanAllContext.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmScanAllContext
}

func (mmScanAllContext *mNamedExecerMockScanAllContext) invocationsDone() bool {
	if len(mmScanAllContext.expectations) == 0 && mmScanAllContext.defaultExpectation == nil && mmScanAllContext.mock.funcScanAllContext == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmScanAllContext.mock.afterScanAllContextCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmScanAllContext.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// ScanAllContext implements mm_database.NamedExecer
func (mmScanAllContext *NamedExecerMock) ScanAllContext(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{}) (err error) {
	mm_atomic.AddUint64(&mmScanAllContext.beforeScanAllContextCounter, 1)
	defer mm_atomic.AddUint64(&mmScanAllContext.afterScanAllContextCounter, 1)

	mmScanAllContext.t.Helper()

	if mmScanAllContext.inspectFuncScanAllContext != nil {
		mmScanAllContext.inspectFuncScanAllContext(ctx, deps, q, args...)
	}

	mm_params := NamedExecerMockScanAllContextParams{ctx, deps, q, args}

	// Record call args
	mmScanAllContext.ScanAllContextMock.mutex.Lock()
	mmScanAllContext.ScanAllContextMock.callArgs = append(mmScanAllContext.ScanAllContextMock.callArgs, &mm_params)
	mmScanAllContext.ScanAllContextMock.mutex.Unlock()

	for _, e := range mmScanAllContext.ScanAllContextMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmScanAllContext.ScanAllContextMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmScanAllContext.ScanAllContextMock.defaultExpectation.Counter, 1)
		mm_want := mmScanAllContext.ScanAllContextMock.defaultExpectation.params
		mm_want_ptrs := mmScanAllContext.ScanAllContextMock.defaultExpectation.paramPtrs

		mm_got := NamedExecerMockScanAllContextParams{ctx, deps, q, args}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmScanAllContext.t.Errorf("NamedExecerMock.ScanAllContext got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmScanAllContext.ScanAllContextMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.deps != nil && !minimock.Equal(*mm_want_ptrs.deps, mm_got.deps) {
				mmScanAllContext.t.Errorf("NamedExecerMock.ScanAllContext got unexpected parameter deps, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmScanAllContext.ScanAllContextMock.defaultExpectation.expectationOrigins.originDeps, *mm_want_ptrs.deps, mm_got.deps, minimock.Diff(*mm_want_ptrs.deps, mm_got.deps))
			}

			if mm_want_ptrs.q != nil && !minimock.Equal(*mm_want_ptrs.q, mm_got.q) {
				mmScanAllContext.t.Errorf("NamedExecerMock.ScanAllContext got unexpected parameter q, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmScanAllContext.ScanAllContextMock.defaultExpectation.expectationOrigins.originQ, *mm_want_ptrs.q, mm_got.q, minimock.Diff(*mm_want_ptrs.q, mm_got.q))
			}

			if mm_want_ptrs.args != nil && !minimock.Equal(*mm_want_ptrs.args, mm_got.args) {
				mmScanAllContext.t.Errorf("NamedExecerMock.ScanAllContext got unexpected parameter args, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmScanAllContext.ScanAllContextMock.defaultExpectation.expectationOrigins.originArgs, *mm_want_ptrs.args, mm_got.args, minimock.Diff(*mm_want_ptrs.args, mm_got.args))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmScanAllContext.t.Errorf("NamedExecerMock.ScanAllContext got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmScanAllContext.ScanAllContextMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmScanAllContext.ScanAllContextMock.defaultExpectation.results
		if mm_results == nil {
			mmScanAllContext.t.Fatal("No results are set for the NamedExecerMock.ScanAllContext")
		}
		return (*mm_results).err
	}
	if mmScanAllContext.funcScanAllContext != nil {
		return mmScanAllContext.funcScanAllContext(ctx, deps, q, args...)
	}
	mmScanAllContext.t.Fatalf("Unexpected call to NamedExecerMock.ScanAllContext. %v %v %v %v", ctx, deps, q, args)
	return
}

// ScanAllContextAfterCounter returns a count of finished NamedExecerMock.ScanAllContext invocations
func (mmScanAllContext *NamedExecerMock) ScanAllContextAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmScanAllContext.afterScanAllContextCounter)
}

// ScanAllContextBeforeCounter returns a count of NamedExecerMock.ScanAllContext invocations
func (mmScanAllContext *NamedExecerMock) ScanAllContextBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmScanAllContext.beforeScanAllContextCounter)
}

// Calls returns a list of arguments used in each call to NamedExecerMock.ScanAllContext.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmScanAllContext *mNamedExecerMockScanAllContext) Calls() []*NamedExecerMockScanAllContextParams {
	mmScanAllContext.mutex.RLock()

	argCopy := make([]*NamedExecerMockScanAllContextParams, len(mmScanAllContext.callArgs))
	copy(argCopy, mmScanAllContext.callArgs)

	mmScanAllContext.mutex.RUnlock()

	return argCopy
}

// MinimockScanAllContextDone returns true if the count of the ScanAllContext invocations corresponds
// the number of defined expectations
func (m *NamedExecerMock) MinimockScanAllContextDone() bool {
	if m.ScanAllContextMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.ScanAllContextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.ScanAllContextMock.invocationsDone()
}

// MinimockScanAllContextInspect logs each unmet expectation
func (m *NamedExecerMock) MinimockScanAllContextInspect() {
	for _, e := range m.ScanAllContextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to NamedExecerMock.ScanAllContext at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterScanAllContextCounter := mm_atomic.LoadUint64(&m.afterScanAllContextCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.ScanAllContextMock.defaultExpectation != nil && afterScanAllContextCounter < 1 {
		if m.ScanAllContextMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to NamedExecerMock.ScanAllContext at\n%s", m.ScanAllContextMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to NamedExecerMock.ScanAllContext at\n%s with params: %#v", m.ScanAllContextMock.defaultExpectation.expectationOrigins.origin, *m.ScanAllContextMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcScanAllContext != nil && afterScanAllContextCounter < 1 {
		m.t.Errorf("Expected call to NamedExecerMock.ScanAllContext at\n%s", m.funcScanAllContextOrigin)
	}

	if !m.ScanAllContextMock.invocationsDone() && afterScanAllContextCounter > 0 {
		m.t.Errorf("Expected %d calls to NamedExecerMock.ScanAllContext at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.ScanAllContextMock.expectedInvocations), m.ScanAllContextMock.expectedInvocationsOrigin, afterScanAllContextCounter)
	}
}

type mNamedExecerMockScanOneContext struct {
	optional           bool
	mock               *NamedExecerMock
	defaultExpectation *NamedExecerMockScanOneContextExpectation
	expectations       []*NamedExecerMockScanOneContextExpectation

	callArgs []*NamedExecerMockScanOneContextParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// NamedExecerMockScanOneContextExpectation specifies expectation struct of the NamedExecer.ScanOneContext
type NamedExecerMockScanOneContextExpectation struct {
	mock               *NamedExecerMock
	params             *NamedExecerMockScanOneContextParams
	paramPtrs          *NamedExecerMockScanOneContextParamPtrs
	expectationOrigins NamedExecerMockScanOneContextExpectationOrigins
	results            *NamedExecerMockScanOneContextResults
	returnOrigin       string
	Counter            uint64
}

// NamedExecerMockScanOneContextParams contains parameters of the NamedExecer.ScanOneContext
type NamedExecerMockScanOneContextParams struct {
	ctx  context.Context
	deps interface{}
	q    mm_database.Query
	args []interface{}
}

// NamedExecerMockScanOneContextParamPtrs contains pointers to parameters of the NamedExecer.ScanOneContext
type NamedExecerMockScanOneContextParamPtrs struct {
	ctx  *context.Context
	deps *interface{}
	q    *mm_database.Query
	args *[]interface{}
}

// NamedExecerMockScanOneContextResults contains results of the NamedExecer.ScanOneContext
type NamedExecerMockScanOneContextResults struct {
	err error
}

// NamedExecerMockScanOneContextOrigins contains origins of expectations of the NamedExecer.ScanOneContext
type NamedExecerMockScanOneContextExpectationOrigins struct {
	origin     string
	originCtx  string
	originDeps string
	originQ    string
	originArgs string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmScanOneContext *mNamedExecerMockScanOneContext) Optional() *mNamedExecerMockScanOneContext {
	mmScanOneContext.optional = true
	return mmScanOneContext
}

// Expect sets up expected params for NamedExecer.ScanOneContext
func (mmScanOneContext *mNamedExecerMockScanOneContext) Expect(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{}) *mNamedExecerMockScanOneContext {
	if mmScanOneContext.mock.funcScanOneContext != nil {
		mmScanOneContext.mock.t.Fatalf("NamedExecerMock.ScanOneContext mock is already set by Set")
	}

	if mmScanOneContext.defaultExpectation == nil {
		mmScanOneContext.defaultExpectation = &NamedExecerMockScanOneContextExpectation{}
	}

	if mmScanOneContext.defaultExpectation.paramPtrs != nil {
		mmScanOneContext.mock.t.Fatalf("NamedExecerMock.ScanOneContext mock is already set by ExpectParams functions")
	}

	mmScanOneContext.defaultExpectation.params = &NamedExecerMockScanOneContextParams{ctx, deps, q, args}
	mmScanOneContext.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmScanOneContext.expectations {
		if minimock.Equal(e.params, mmScanOneContext.defaultExpectation.params) {
			mmScanOneContext.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmScanOneContext.defaultExpectation.params)
		}
	}

	return mmScanOneContext
}

// ExpectCtxParam1 sets up expected param ctx for NamedExecer.ScanOneContext
func (mmScanOneContext *mNamedExecerMockScanOneContext) ExpectCtxParam1(ctx context.Context) *mNamedExecerMockScanOneContext {
	if mmScanOneContext.mock.funcScanOneContext != nil {
		mmScanOneContext.mock.t.Fatalf("NamedExecerMock.ScanOneContext mock is already set by Set")
	}

	if mmScanOneContext.defaultExpectation == nil {
		mmScanOneContext.defaultExpectation = &NamedExecerMockScanOneContextExpectation{}
	}

	if mmScanOneContext.defaultExpectation.params != nil {
		mmScanOneContext.mock.t.Fatalf("NamedExecerMock.ScanOneContext mock is already set by Expect")
	}

	if mmScanOneContext.defaultExpectation.paramPtrs == nil {
		mmScanOneContext.defaultExpectation.paramPtrs = &NamedExecerMockScanOneContextParamPtrs{}
	}
	mmScanOneContext.defaultExpectation.paramPtrs.ctx = &ctx
	mmScanOneContext.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmScanOneContext
}

// ExpectDepsParam2 sets up expected param deps for NamedExecer.ScanOneContext
func (mmScanOneContext *mNamedExecerMockScanOneContext) ExpectDepsParam2(deps interface{}) *mNamedExecerMockScanOneContext {
	if mmScanOneContext.mock.funcScanOneContext != nil {
		mmScanOneContext.mock.t.Fatalf("NamedExecerMock.ScanOneContext mock is already set by Set")
	}

	if mmScanOneContext.defaultExpectation == nil {
		mmScanOneContext.defaultExpectation = &NamedExecerMockScanOneContextExpectation{}
	}

	if mmScanOneContext.defaultExpectation.params != nil {
		mmScanOneContext.mock.t.Fatalf("NamedExecerMock.ScanOneContext mock is already set by Expect")
	}

	if mmScanOneContext.defaultExpectation.paramPtrs == nil {
		mmScanOneContext.defaultExpectation.paramPtrs = &NamedExecerMockScanOneContextParamPtrs{}
	}
	mmScanOneContext.defaultExpectation.paramPtrs.deps = &deps
	mmScanOneContext.defaultExpectation.expectationOrigins.originDeps = minimock.CallerInfo(1)

	return mmScanOneContext
}

// ExpectQParam3 sets up expected param q for NamedExecer.ScanOneContext
func (mmScanOneContext *mNamedExecerMockScanOneContext) ExpectQParam3(q mm_database.Query) *mNamedExecerMockScanOneContext {
	if mmScanOneContext.mock.funcScanOneContext != nil {
		mmScanOneContext.mock.t.Fatalf("NamedExecerMock.ScanOneContext mock is already set by Set")
	}

	if mmScanOneContext.defaultExpectation == nil {
		mmScanOneContext.defaultExpectation = &NamedExecerMockScanOneContextExpectation{}
	}

	if mmScanOneContext.defaultExpectation.params != nil {
		mmScanOneContext.mock.t.Fatalf("NamedExecerMock.ScanOneContext mock is already set by Expect")
	}

	if mmScanOneContext.defaultExpectation.paramPtrs == nil {
		mmScanOneContext.defaultExpectation.paramPtrs = &NamedExecerMockScanOneContextParamPtrs{}
	}
	mmScanOneContext.defaultExpectation.paramPtrs.q = &q
	mmScanOneContext.defaultExpectation.expectationOrigins.originQ = minimock.CallerInfo(1)

	return mmScanOneContext
}

// ExpectArgsParam4 sets up expected param args for NamedExecer.ScanOneContext
func (mmScanOneContext *mNamedExecerMockScanOneContext) ExpectArgsParam4(args ...interface{}) *mNamedExecerMockScanOneContext {
	if mmScanOneContext.mock.funcScanOneContext != nil {
		mmScanOneContext.mock.t.Fatalf("NamedExecerMock.ScanOneContext mock is already set by Set")
	}

	if mmScanOneContext.defaultExpectation == nil {
		mmScanOneContext.defaultExpectation = &NamedExecerMockScanOneContextExpectation{}
	}

	if mmScanOneContext.defaultExpectation.params != nil {
		mmScanOneContext.mock.t.Fatalf("NamedExecerMock.ScanOneContext mock is already set by Expect")
	}

	if mmScanOneContext.defaultExpectation.paramPtrs == nil {
		mmScanOneContext.defaultExpectation.paramPtrs = &NamedExecerMockScanOneContextParamPtrs{}
	}
	mmScanOneContext.defaultExpectation.paramPtrs.args = &args
	mmScanOneContext.defaultExpectation.expectationOrigins.originArgs = minimock.CallerInfo(1)

	return mmScanOneContext
}

// Inspect accepts an inspector function that has same arguments as the NamedExecer.ScanOneContext
func (mmScanOneContext *mNamedExecerMockScanOneContext) Inspect(f func(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{})) *mNamedExecerMockScanOneContext {
	if mmScanOneContext.mock.inspectFuncScanOneContext != nil {
		mmScanOneContext.mock.t.Fatalf("Inspect function is already set for NamedExecerMock.ScanOneContext")
	}

	mmScanOneContext.mock.inspectFuncScanOneContext = f

	return mmScanOneContext
}

// Return sets up results that will be returned by NamedExecer.ScanOneContext
func (mmScanOneContext *mNamedExecerMockScanOneContext) Return(err error) *NamedExecerMock {
	if mmScanOneContext.mock.funcScanOneContext != nil {
		mmScanOneContext.mock.t.Fatalf("NamedExecerMock.ScanOneContext mock is already set by Set")
	}

	if mmScanOneContext.defaultExpectation == nil {
		mmScanOneContext.defaultExpectation = &NamedExecerMockScanOneContextExpectation{mock: mmScanOneContext.mock}
	}
	mmScanOneContext.defaultExpectation.results = &NamedExecerMockScanOneContextResults{err}
	mmScanOneContext.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmScanOneContext.mock
}

// Set uses given function f to mock the NamedExecer.ScanOneContext method
func (mmScanOneContext *mNamedExecerMockScanOneContext) Set(f func(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{}) (err error)) *NamedExecerMock {
	if mmScanOneContext.defaultExpectation != nil {
		mmScanOneContext.mock.t.Fatalf("Default expectation is already set for the NamedExecer.ScanOneContext method")
	}

	if len(mmScanOneContext.expectations) > 0 {
		mmScanOneContext.mock.t.Fatalf("Some expectations are already set for the NamedExecer.ScanOneContext method")
	}

	mmScanOneContext.mock.funcScanOneContext = f
	mmScanOneContext.mock.funcScanOneContextOrigin = minimock.CallerInfo(1)
	return mmScanOneContext.mock
}

// When sets expectation for the NamedExecer.ScanOneContext which will trigger the result defined by the following
// Then helper
func (mmScanOneContext *mNamedExecerMockScanOneContext) When(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{}) *NamedExecerMockScanOneContextExpectation {
	if mmScanOneContext.mock.funcScanOneContext != nil {
		mmScanOneContext.mock.t.Fatalf("NamedExecerMock.ScanOneContext mock is already set by Set")
	}

	expectation := &NamedExecerMockScanOneContextExpectation{
		mock:               mmScanOneContext.mock,
		params:             &NamedExecerMockScanOneContextParams{ctx, deps, q, args},
		expectationOrigins: NamedExecerMockScanOneContextExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmScanOneContext.expectations = append(mmScanOneContext.expectations, expectation)
	return expectation
}

// Then sets up NamedExecer.ScanOneContext return parameters for the expectation previously defined by the When method
func (e *NamedExecerMockScanOneContextExpectation) Then(err error) *NamedExecerMock {
	e.results = &NamedExecerMockScanOneContextResults{err}
	return e.mock
}

// Times sets number of times NamedExecer.ScanOneContext should be invoked
func (mmScanOneContext *mNamedExecerMockScanOneContext) Times(n uint64) *mNamedExecerMockScanOneContext {
	if n == 0 {
		mmScanOneContext.mock.t.Fatalf("Times of NamedExecerMock.ScanOneContext mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmScanOneContext.expectedInvocations, n)
	mmScanOneContext.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmScanOneContext
}

func (mmScanOneContext *mNamedExecerMockScanOneContext) invocationsDone() bool {
	if len(mmScanOneContext.expectations) == 0 && mmScanOneContext.defaultExpectation == nil && mmScanOneContext.mock.funcScanOneContext == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmScanOneContext.mock.afterScanOneContextCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmScanOneContext.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// ScanOneContext implements mm_database.NamedExecer
func (mmScanOneContext *NamedExecerMock) ScanOneContext(ctx context.Context, deps interface{}, q mm_database.Query, args ...interface{}) (err error) {
	mm_atomic.AddUint64(&mmScanOneContext.beforeScanOneContextCounter, 1)
	defer mm_atomic.AddUint64(&mmScanOneContext.afterScanOneContextCounter, 1)

	mmScanOneContext.t.Helper()

	if mmScanOneContext.inspectFuncScanOneContext != nil {
		mmScanOneContext.inspectFuncScanOneContext(ctx, deps, q, args...)
	}

	mm_params := NamedExecerMockScanOneContextParams{ctx, deps, q, args}

	// Record call args
	mmScanOneContext.ScanOneContextMock.mutex.Lock()
	mmScanOneContext.ScanOneContextMock.callArgs = append(mmScanOneContext.ScanOneContextMock.callArgs, &mm_params)
	mmScanOneContext.ScanOneContextMock.mutex.Unlock()

	for _, e := range mmScanOneContext.ScanOneContextMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmScanOneContext.ScanOneContextMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmScanOneContext.ScanOneContextMock.defaultExpectation.Counter, 1)
		mm_want := mmScanOneContext.ScanOneContextMock.defaultExpectation.params
		mm_want_ptrs := mmScanOneContext.ScanOneContextMock.defaultExpectation.paramPtrs

		mm_got := NamedExecerMockScanOneContextParams{ctx, deps, q, args}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmScanOneContext.t.Errorf("NamedExecerMock.ScanOneContext got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmScanOneContext.ScanOneContextMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.deps != nil && !minimock.Equal(*mm_want_ptrs.deps, mm_got.deps) {
				mmScanOneContext.t.Errorf("NamedExecerMock.ScanOneContext got unexpected parameter deps, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmScanOneContext.ScanOneContextMock.defaultExpectation.expectationOrigins.originDeps, *mm_want_ptrs.deps, mm_got.deps, minimock.Diff(*mm_want_ptrs.deps, mm_got.deps))
			}

			if mm_want_ptrs.q != nil && !minimock.Equal(*mm_want_ptrs.q, mm_got.q) {
				mmScanOneContext.t.Errorf("NamedExecerMock.ScanOneContext got unexpected parameter q, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmScanOneContext.ScanOneContextMock.defaultExpectation.expectationOrigins.originQ, *mm_want_ptrs.q, mm_got.q, minimock.Diff(*mm_want_ptrs.q, mm_got.q))
			}

			if mm_want_ptrs.args != nil && !minimock.Equal(*mm_want_ptrs.args, mm_got.args) {
				mmScanOneContext.t.Errorf("NamedExecerMock.ScanOneContext got unexpected parameter args, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmScanOneContext.ScanOneContextMock.defaultExpectation.expectationOrigins.originArgs, *mm_want_ptrs.args, mm_got.args, minimock.Diff(*mm_want_ptrs.args, mm_got.args))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmScanOneContext.t.Errorf("NamedExecerMock.ScanOneContext got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmScanOneContext.ScanOneContextMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmScanOneContext.ScanOneContextMock.defaultExpectation.results
		if mm_results == nil {
			mmScanOneContext.t.Fatal("No results are set for the NamedExecerMock.ScanOneContext")
		}
		return (*mm_results).err
	}
	if mmScanOneContext.funcScanOneContext != nil {
		return mmScanOneContext.funcScanOneContext(ctx, deps, q, args...)
	}
	mmScanOneContext.t.Fatalf("Unexpected call to NamedExecerMock.ScanOneContext. %v %v %v %v", ctx, deps, q, args)
	return
}

// ScanOneContextAfterCounter returns a count of finished NamedExecerMock.ScanOneContext invocations
func (mmScanOneContext *NamedExecerMock) ScanOneContextAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmScanOneContext.afterScanOneContextCounter)
}

// ScanOneContextBeforeCounter returns a count of NamedExecerMock.ScanOneContext invocations
func (mmScanOneContext *NamedExecerMock) ScanOneContextBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmScanOneContext.beforeScanOneContextCounter)
}

// Calls returns a list of arguments used in each call to NamedExecerMock.ScanOneContext.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmScanOneContext *mNamedExecerMockScanOneContext) Calls() []*NamedExecerMockScanOneContextParams {
	mmScanOneContext.mutex.RLock()

	argCopy := make([]*NamedExecerMockScanOneContextParams, len(mmScanOneContext.callArgs))
	copy(argCopy, mmScanOneContext.callArgs)

	mmScanOneContext.mutex.RUnlock()

	return argCopy
}

// MinimockScanOneContextDone returns true if the count of the ScanOneContext invocations corresponds
// the number of defined expectations
func (m *NamedExecerMock) MinimockScanOneContextDone() bool {
	if m.ScanOneContextMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.ScanOneContextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.ScanOneContextMock.invocationsDone()
}

// MinimockScanOneContextInspect logs each unmet expectation
func (m *NamedExecerMock) MinimockScanOneContextInspect() {
	for _, e := range m.ScanOneContextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to NamedExecerMock.ScanOneContext at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterScanOneContextCounter := mm_atomic.LoadUint64(&m.afterScanOneContextCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.ScanOneContextMock.defaultExpectation != nil && afterScanOneContextCounter < 1 {
		if m.ScanOneContextMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to NamedExecerMock.ScanOneContext at\n%s", m.ScanOneContextMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to NamedExecerMock.ScanOneContext at\n%s with params: %#v", m.ScanOneContextMock.defaultExpectation.expectationOrigins.origin, *m.ScanOneContextMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcScanOneContext != nil && afterScanOneContextCounter < 1 {
		m.t.Errorf("Expected call to NamedExecerMock.ScanOneContext at\n%s", m.funcScanOneContextOrigin)
	}

	if !m.ScanOneContextMock.invocationsDone() && afterScanOneContextCounter > 0 {
		m.t.Errorf("Expected %d calls to NamedExecerMock.ScanOneContext at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.ScanOneContextMock.expectedInvocations), m.ScanOneContextMock.expectedInvocationsOrigin, afterScanOneContextCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *NamedExecerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockScanAllContextInspect()

			m.MinimockScanOneContextInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *NamedExecerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *NamedExecerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockScanAllContextDone() &&
		m.MinimockScanOneContextDone()
}
