// Code generated by MockGen. DO NOT EDIT.
// Source: generics.go

// Package source is a generated GoMock package.
package source

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	typed "github.com/golang/mock/mockgen/internal/tests/typed"
	other "github.com/golang/mock/mockgen/internal/tests/typed/other"
)

// MockBar is a mock of Bar interface.
type MockBar[T any, R any] struct {
	ctrl     *gomock.Controller
	recorder *MockBarMockRecorder[T, R]
}

// MockBarMockRecorder is the mock recorder for MockBar.
type MockBarMockRecorder[T any, R any] struct {
	mock *MockBar[T, R]
}

// NewMockBar creates a new mock instance.
func NewMockBar[T any, R any](ctrl *gomock.Controller) *MockBar[T, R] {
	mock := &MockBar[T, R]{ctrl: ctrl}
	mock.recorder = &MockBarMockRecorder[T, R]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBar[T, R]) EXPECT() *MockBarMockRecorder[T, R] {
	return m.recorder
}

// Eight mocks base method.
func (m *MockBar[T, R]) Eight(arg0 T) other.Two[T, R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eight", arg0)
	ret0, _ := ret[0].(other.Two[T, R])
	return ret0
}

// Eight indicates an expected call of Eight.
func (mr *MockBarMockRecorder[T, R]) Eight(arg0 interface{}) *BarEightCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eight", reflect.TypeOf((*MockBar[T, R])(nil).Eight), arg0)
	return &BarEightCall[T, R]{Call: call}
}

//  BarEightCall wrap *gomock.Call
type BarEightCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarEightCall[T, R]) Return(arg0 other.Two[T, R]) *BarEightCall[T, R] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarEightCall[T, R]) Do(f func(T) other.Two[T, R]) *BarEightCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarEightCall[T, R]) DoAndReturn(f func(T) other.Two[T, R]) *BarEightCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Eighteen mocks base method.
func (m *MockBar[T, R]) Eighteen() (typed.Iface[*other.Five], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eighteen")
	ret0, _ := ret[0].(typed.Iface[*other.Five])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Eighteen indicates an expected call of Eighteen.
func (mr *MockBarMockRecorder[T, R]) Eighteen() *BarEighteenCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eighteen", reflect.TypeOf((*MockBar[T, R])(nil).Eighteen))
	return &BarEighteenCall[T, R]{Call: call}
}

//  BarEighteenCall wrap *gomock.Call
type BarEighteenCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarEighteenCall[T, R]) Return(arg0 typed.Iface[*other.Five], arg1 error) *BarEighteenCall[T, R] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarEighteenCall[T, R]) Do(f func() (typed.Iface[*other.Five], error)) *BarEighteenCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarEighteenCall[T, R]) DoAndReturn(f func() (typed.Iface[*other.Five], error)) *BarEighteenCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Eleven mocks base method.
func (m *MockBar[T, R]) Eleven() (*other.One[T], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eleven")
	ret0, _ := ret[0].(*other.One[T])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Eleven indicates an expected call of Eleven.
func (mr *MockBarMockRecorder[T, R]) Eleven() *BarElevenCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eleven", reflect.TypeOf((*MockBar[T, R])(nil).Eleven))
	return &BarElevenCall[T, R]{Call: call}
}

//  BarElevenCall wrap *gomock.Call
type BarElevenCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarElevenCall[T, R]) Return(arg0 *other.One[T], arg1 error) *BarElevenCall[T, R] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarElevenCall[T, R]) Do(f func() (*other.One[T], error)) *BarElevenCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarElevenCall[T, R]) DoAndReturn(f func() (*other.One[T], error)) *BarElevenCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Fifteen mocks base method.
func (m *MockBar[T, R]) Fifteen() (typed.Iface[typed.StructType], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fifteen")
	ret0, _ := ret[0].(typed.Iface[typed.StructType])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fifteen indicates an expected call of Fifteen.
func (mr *MockBarMockRecorder[T, R]) Fifteen() *BarFifteenCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fifteen", reflect.TypeOf((*MockBar[T, R])(nil).Fifteen))
	return &BarFifteenCall[T, R]{Call: call}
}

//  BarFifteenCall wrap *gomock.Call
type BarFifteenCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarFifteenCall[T, R]) Return(arg0 typed.Iface[typed.StructType], arg1 error) *BarFifteenCall[T, R] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarFifteenCall[T, R]) Do(f func() (typed.Iface[typed.StructType], error)) *BarFifteenCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarFifteenCall[T, R]) DoAndReturn(f func() (typed.Iface[typed.StructType], error)) *BarFifteenCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Five mocks base method.
func (m *MockBar[T, R]) Five(arg0 T) typed.Baz[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Five", arg0)
	ret0, _ := ret[0].(typed.Baz[T])
	return ret0
}

// Five indicates an expected call of Five.
func (mr *MockBarMockRecorder[T, R]) Five(arg0 interface{}) *BarFiveCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Five", reflect.TypeOf((*MockBar[T, R])(nil).Five), arg0)
	return &BarFiveCall[T, R]{Call: call}
}

//  BarFiveCall wrap *gomock.Call
type BarFiveCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarFiveCall[T, R]) Return(arg0 typed.Baz[T]) *BarFiveCall[T, R] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarFiveCall[T, R]) Do(f func(T) typed.Baz[T]) *BarFiveCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarFiveCall[T, R]) DoAndReturn(f func(T) typed.Baz[T]) *BarFiveCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Four mocks base method.
func (m *MockBar[T, R]) Four(arg0 T) typed.Foo[T, R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Four", arg0)
	ret0, _ := ret[0].(typed.Foo[T, R])
	return ret0
}

// Four indicates an expected call of Four.
func (mr *MockBarMockRecorder[T, R]) Four(arg0 interface{}) *BarFourCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Four", reflect.TypeOf((*MockBar[T, R])(nil).Four), arg0)
	return &BarFourCall[T, R]{Call: call}
}

//  BarFourCall wrap *gomock.Call
type BarFourCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarFourCall[T, R]) Return(arg0 typed.Foo[T, R]) *BarFourCall[T, R] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarFourCall[T, R]) Do(f func(T) typed.Foo[T, R]) *BarFourCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarFourCall[T, R]) DoAndReturn(f func(T) typed.Foo[T, R]) *BarFourCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Fourteen mocks base method.
func (m *MockBar[T, R]) Fourteen() (*typed.Foo[typed.StructType, typed.StructType2], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fourteen")
	ret0, _ := ret[0].(*typed.Foo[typed.StructType, typed.StructType2])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fourteen indicates an expected call of Fourteen.
func (mr *MockBarMockRecorder[T, R]) Fourteen() *BarFourteenCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fourteen", reflect.TypeOf((*MockBar[T, R])(nil).Fourteen))
	return &BarFourteenCall[T, R]{Call: call}
}

//  BarFourteenCall wrap *gomock.Call
type BarFourteenCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarFourteenCall[T, R]) Return(arg0 *typed.Foo[typed.StructType, typed.StructType2], arg1 error) *BarFourteenCall[T, R] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarFourteenCall[T, R]) Do(f func() (*typed.Foo[typed.StructType, typed.StructType2], error)) *BarFourteenCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarFourteenCall[T, R]) DoAndReturn(f func() (*typed.Foo[typed.StructType, typed.StructType2], error)) *BarFourteenCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Nine mocks base method.
func (m *MockBar[T, R]) Nine(arg0 typed.Iface[T]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Nine", arg0)
}

// Nine indicates an expected call of Nine.
func (mr *MockBarMockRecorder[T, R]) Nine(arg0 interface{}) *BarNineCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nine", reflect.TypeOf((*MockBar[T, R])(nil).Nine), arg0)
	return &BarNineCall[T, R]{Call: call}
}

//  BarNineCall wrap *gomock.Call
type BarNineCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarNineCall[T, R]) Return() *BarNineCall[T, R] {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarNineCall[T, R]) Do(f func(typed.Iface[T])) *BarNineCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarNineCall[T, R]) DoAndReturn(f func(typed.Iface[T])) *BarNineCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Nineteen mocks base method.
func (m *MockBar[T, R]) Nineteen() typed.AliasType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nineteen")
	ret0, _ := ret[0].(typed.AliasType)
	return ret0
}

// Nineteen indicates an expected call of Nineteen.
func (mr *MockBarMockRecorder[T, R]) Nineteen() *BarNineteenCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nineteen", reflect.TypeOf((*MockBar[T, R])(nil).Nineteen))
	return &BarNineteenCall[T, R]{Call: call}
}

//  BarNineteenCall wrap *gomock.Call
type BarNineteenCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarNineteenCall[T, R]) Return(arg0 typed.AliasType) *BarNineteenCall[T, R] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarNineteenCall[T, R]) Do(f func() typed.AliasType) *BarNineteenCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarNineteenCall[T, R]) DoAndReturn(f func() typed.AliasType) *BarNineteenCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// One mocks base method.
func (m *MockBar[T, R]) One(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// One indicates an expected call of One.
func (mr *MockBarMockRecorder[T, R]) One(arg0 interface{}) *BarOneCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockBar[T, R])(nil).One), arg0)
	return &BarOneCall[T, R]{Call: call}
}

//  BarOneCall wrap *gomock.Call
type BarOneCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarOneCall[T, R]) Return(arg0 string) *BarOneCall[T, R] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarOneCall[T, R]) Do(f func(string) string) *BarOneCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarOneCall[T, R]) DoAndReturn(f func(string) string) *BarOneCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Seven mocks base method.
func (m *MockBar[T, R]) Seven(arg0 T) other.One[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seven", arg0)
	ret0, _ := ret[0].(other.One[T])
	return ret0
}

// Seven indicates an expected call of Seven.
func (mr *MockBarMockRecorder[T, R]) Seven(arg0 interface{}) *BarSevenCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seven", reflect.TypeOf((*MockBar[T, R])(nil).Seven), arg0)
	return &BarSevenCall[T, R]{Call: call}
}

//  BarSevenCall wrap *gomock.Call
type BarSevenCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarSevenCall[T, R]) Return(arg0 other.One[T]) *BarSevenCall[T, R] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarSevenCall[T, R]) Do(f func(T) other.One[T]) *BarSevenCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarSevenCall[T, R]) DoAndReturn(f func(T) other.One[T]) *BarSevenCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Seventeen mocks base method.
func (m *MockBar[T, R]) Seventeen() (*typed.Foo[other.Three, other.Four], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seventeen")
	ret0, _ := ret[0].(*typed.Foo[other.Three, other.Four])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seventeen indicates an expected call of Seventeen.
func (mr *MockBarMockRecorder[T, R]) Seventeen() *BarSeventeenCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seventeen", reflect.TypeOf((*MockBar[T, R])(nil).Seventeen))
	return &BarSeventeenCall[T, R]{Call: call}
}

//  BarSeventeenCall wrap *gomock.Call
type BarSeventeenCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarSeventeenCall[T, R]) Return(arg0 *typed.Foo[other.Three, other.Four], arg1 error) *BarSeventeenCall[T, R] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarSeventeenCall[T, R]) Do(f func() (*typed.Foo[other.Three, other.Four], error)) *BarSeventeenCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarSeventeenCall[T, R]) DoAndReturn(f func() (*typed.Foo[other.Three, other.Four], error)) *BarSeventeenCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Six mocks base method.
func (m *MockBar[T, R]) Six(arg0 T) *typed.Baz[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Six", arg0)
	ret0, _ := ret[0].(*typed.Baz[T])
	return ret0
}

// Six indicates an expected call of Six.
func (mr *MockBarMockRecorder[T, R]) Six(arg0 interface{}) *BarSixCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Six", reflect.TypeOf((*MockBar[T, R])(nil).Six), arg0)
	return &BarSixCall[T, R]{Call: call}
}

//  BarSixCall wrap *gomock.Call
type BarSixCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarSixCall[T, R]) Return(arg0 *typed.Baz[T]) *BarSixCall[T, R] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarSixCall[T, R]) Do(f func(T) *typed.Baz[T]) *BarSixCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarSixCall[T, R]) DoAndReturn(f func(T) *typed.Baz[T]) *BarSixCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Sixteen mocks base method.
func (m *MockBar[T, R]) Sixteen() (typed.Baz[other.Three], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sixteen")
	ret0, _ := ret[0].(typed.Baz[other.Three])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sixteen indicates an expected call of Sixteen.
func (mr *MockBarMockRecorder[T, R]) Sixteen() *BarSixteenCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sixteen", reflect.TypeOf((*MockBar[T, R])(nil).Sixteen))
	return &BarSixteenCall[T, R]{Call: call}
}

//  BarSixteenCall wrap *gomock.Call
type BarSixteenCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarSixteenCall[T, R]) Return(arg0 typed.Baz[other.Three], arg1 error) *BarSixteenCall[T, R] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarSixteenCall[T, R]) Do(f func() (typed.Baz[other.Three], error)) *BarSixteenCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarSixteenCall[T, R]) DoAndReturn(f func() (typed.Baz[other.Three], error)) *BarSixteenCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Ten mocks base method.
func (m *MockBar[T, R]) Ten(arg0 *T) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ten", arg0)
}

// Ten indicates an expected call of Ten.
func (mr *MockBarMockRecorder[T, R]) Ten(arg0 interface{}) *BarTenCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ten", reflect.TypeOf((*MockBar[T, R])(nil).Ten), arg0)
	return &BarTenCall[T, R]{Call: call}
}

//  BarTenCall wrap *gomock.Call
type BarTenCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarTenCall[T, R]) Return() *BarTenCall[T, R] {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarTenCall[T, R]) Do(f func(*T)) *BarTenCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarTenCall[T, R]) DoAndReturn(f func(*T)) *BarTenCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Thirteen mocks base method.
func (m *MockBar[T, R]) Thirteen() (typed.Baz[typed.StructType], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Thirteen")
	ret0, _ := ret[0].(typed.Baz[typed.StructType])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Thirteen indicates an expected call of Thirteen.
func (mr *MockBarMockRecorder[T, R]) Thirteen() *BarThirteenCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Thirteen", reflect.TypeOf((*MockBar[T, R])(nil).Thirteen))
	return &BarThirteenCall[T, R]{Call: call}
}

//  BarThirteenCall wrap *gomock.Call
type BarThirteenCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarThirteenCall[T, R]) Return(arg0 typed.Baz[typed.StructType], arg1 error) *BarThirteenCall[T, R] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarThirteenCall[T, R]) Do(f func() (typed.Baz[typed.StructType], error)) *BarThirteenCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarThirteenCall[T, R]) DoAndReturn(f func() (typed.Baz[typed.StructType], error)) *BarThirteenCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Three mocks base method.
func (m *MockBar[T, R]) Three(arg0 T) R {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Three", arg0)
	ret0, _ := ret[0].(R)
	return ret0
}

// Three indicates an expected call of Three.
func (mr *MockBarMockRecorder[T, R]) Three(arg0 interface{}) *BarThreeCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Three", reflect.TypeOf((*MockBar[T, R])(nil).Three), arg0)
	return &BarThreeCall[T, R]{Call: call}
}

//  BarThreeCall wrap *gomock.Call
type BarThreeCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarThreeCall[T, R]) Return(arg0 R) *BarThreeCall[T, R] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarThreeCall[T, R]) Do(f func(T) R) *BarThreeCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarThreeCall[T, R]) DoAndReturn(f func(T) R) *BarThreeCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Twelve mocks base method.
func (m *MockBar[T, R]) Twelve() (*other.Two[T, R], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Twelve")
	ret0, _ := ret[0].(*other.Two[T, R])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Twelve indicates an expected call of Twelve.
func (mr *MockBarMockRecorder[T, R]) Twelve() *BarTwelveCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Twelve", reflect.TypeOf((*MockBar[T, R])(nil).Twelve))
	return &BarTwelveCall[T, R]{Call: call}
}

//  BarTwelveCall wrap *gomock.Call
type BarTwelveCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarTwelveCall[T, R]) Return(arg0 *other.Two[T, R], arg1 error) *BarTwelveCall[T, R] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarTwelveCall[T, R]) Do(f func() (*other.Two[T, R], error)) *BarTwelveCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarTwelveCall[T, R]) DoAndReturn(f func() (*other.Two[T, R], error)) *BarTwelveCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Two mocks base method.
func (m *MockBar[T, R]) Two(arg0 T) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Two", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Two indicates an expected call of Two.
func (mr *MockBarMockRecorder[T, R]) Two(arg0 interface{}) *BarTwoCall[T, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Two", reflect.TypeOf((*MockBar[T, R])(nil).Two), arg0)
	return &BarTwoCall[T, R]{Call: call}
}

//  BarTwoCall wrap *gomock.Call
type BarTwoCall[T any, R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BarTwoCall[T, R]) Return(arg0 string) *BarTwoCall[T, R] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BarTwoCall[T, R]) Do(f func(T) string) *BarTwoCall[T, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BarTwoCall[T, R]) DoAndReturn(f func(T) string) *BarTwoCall[T, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockIface is a mock of Iface interface.
type MockIface[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockIfaceMockRecorder[T]
}

// MockIfaceMockRecorder is the mock recorder for MockIface.
type MockIfaceMockRecorder[T any] struct {
	mock *MockIface[T]
}

// NewMockIface creates a new mock instance.
func NewMockIface[T any](ctrl *gomock.Controller) *MockIface[T] {
	mock := &MockIface[T]{ctrl: ctrl}
	mock.recorder = &MockIfaceMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIface[T]) EXPECT() *MockIfaceMockRecorder[T] {
	return m.recorder
}
