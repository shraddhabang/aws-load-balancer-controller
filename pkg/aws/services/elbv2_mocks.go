// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/aws-load-balancer-controller/pkg/aws/services (interfaces: ELBV2)

// Package services is a generated GoMock package.
package services

import (
	context "context"
	reflect "reflect"

	elasticloadbalancingv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	gomock "github.com/golang/mock/gomock"
)

// MockELBV2 is a mock of ELBV2 interface.
type MockELBV2 struct {
	ctrl     *gomock.Controller
	recorder *MockELBV2MockRecorder
}

// MockELBV2MockRecorder is the mock recorder for MockELBV2.
type MockELBV2MockRecorder struct {
	mock *MockELBV2
}

// NewMockELBV2 creates a new mock instance.
func NewMockELBV2(ctrl *gomock.Controller) *MockELBV2 {
	mock := &MockELBV2{ctrl: ctrl}
	mock.recorder = &MockELBV2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockELBV2) EXPECT() *MockELBV2MockRecorder {
	return m.recorder
}

// AddListenerCertificatesWithContext mocks base method.
func (m *MockELBV2) AddListenerCertificatesWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.AddListenerCertificatesInput) (*elasticloadbalancingv2.AddListenerCertificatesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddListenerCertificatesWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.AddListenerCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddListenerCertificatesWithContext indicates an expected call of AddListenerCertificatesWithContext.
func (mr *MockELBV2MockRecorder) AddListenerCertificatesWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddListenerCertificatesWithContext", reflect.TypeOf((*MockELBV2)(nil).AddListenerCertificatesWithContext), arg0, arg1)
}

// AddTagsWithContext mocks base method.
func (m *MockELBV2) AddTagsWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.AddTagsInput) (*elasticloadbalancingv2.AddTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTagsWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.AddTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTagsWithContext indicates an expected call of AddTagsWithContext.
func (mr *MockELBV2MockRecorder) AddTagsWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTagsWithContext", reflect.TypeOf((*MockELBV2)(nil).AddTagsWithContext), arg0, arg1)
}

// CreateListenerWithContext mocks base method.
func (m *MockELBV2) CreateListenerWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.CreateListenerInput) (*elasticloadbalancingv2.CreateListenerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateListenerWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.CreateListenerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateListenerWithContext indicates an expected call of CreateListenerWithContext.
func (mr *MockELBV2MockRecorder) CreateListenerWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateListenerWithContext", reflect.TypeOf((*MockELBV2)(nil).CreateListenerWithContext), arg0, arg1)
}

// CreateLoadBalancerWithContext mocks base method.
func (m *MockELBV2) CreateLoadBalancerWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.CreateLoadBalancerInput) (*elasticloadbalancingv2.CreateLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoadBalancerWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.CreateLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLoadBalancerWithContext indicates an expected call of CreateLoadBalancerWithContext.
func (mr *MockELBV2MockRecorder) CreateLoadBalancerWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoadBalancerWithContext", reflect.TypeOf((*MockELBV2)(nil).CreateLoadBalancerWithContext), arg0, arg1)
}

// CreateRuleWithContext mocks base method.
func (m *MockELBV2) CreateRuleWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.CreateRuleInput) (*elasticloadbalancingv2.CreateRuleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRuleWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.CreateRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRuleWithContext indicates an expected call of CreateRuleWithContext.
func (mr *MockELBV2MockRecorder) CreateRuleWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRuleWithContext", reflect.TypeOf((*MockELBV2)(nil).CreateRuleWithContext), arg0, arg1)
}

// CreateTargetGroupWithContext mocks base method.
func (m *MockELBV2) CreateTargetGroupWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.CreateTargetGroupInput) (*elasticloadbalancingv2.CreateTargetGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTargetGroupWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.CreateTargetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTargetGroupWithContext indicates an expected call of CreateTargetGroupWithContext.
func (mr *MockELBV2MockRecorder) CreateTargetGroupWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTargetGroupWithContext", reflect.TypeOf((*MockELBV2)(nil).CreateTargetGroupWithContext), arg0, arg1)
}

// DeleteListenerWithContext mocks base method.
func (m *MockELBV2) DeleteListenerWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DeleteListenerInput) (*elasticloadbalancingv2.DeleteListenerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteListenerWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DeleteListenerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteListenerWithContext indicates an expected call of DeleteListenerWithContext.
func (mr *MockELBV2MockRecorder) DeleteListenerWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteListenerWithContext", reflect.TypeOf((*MockELBV2)(nil).DeleteListenerWithContext), arg0, arg1)
}

// DeleteLoadBalancerWithContext mocks base method.
func (m *MockELBV2) DeleteLoadBalancerWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DeleteLoadBalancerInput) (*elasticloadbalancingv2.DeleteLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLoadBalancerWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DeleteLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLoadBalancerWithContext indicates an expected call of DeleteLoadBalancerWithContext.
func (mr *MockELBV2MockRecorder) DeleteLoadBalancerWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLoadBalancerWithContext", reflect.TypeOf((*MockELBV2)(nil).DeleteLoadBalancerWithContext), arg0, arg1)
}

// DeleteRuleWithContext mocks base method.
func (m *MockELBV2) DeleteRuleWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DeleteRuleInput) (*elasticloadbalancingv2.DeleteRuleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRuleWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DeleteRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRuleWithContext indicates an expected call of DeleteRuleWithContext.
func (mr *MockELBV2MockRecorder) DeleteRuleWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRuleWithContext", reflect.TypeOf((*MockELBV2)(nil).DeleteRuleWithContext), arg0, arg1)
}

// DeleteTargetGroupWithContext mocks base method.
func (m *MockELBV2) DeleteTargetGroupWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DeleteTargetGroupInput) (*elasticloadbalancingv2.DeleteTargetGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTargetGroupWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DeleteTargetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTargetGroupWithContext indicates an expected call of DeleteTargetGroupWithContext.
func (mr *MockELBV2MockRecorder) DeleteTargetGroupWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTargetGroupWithContext", reflect.TypeOf((*MockELBV2)(nil).DeleteTargetGroupWithContext), arg0, arg1)
}

// DeregisterTargetsWithContext mocks base method.
func (m *MockELBV2) DeregisterTargetsWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DeregisterTargetsInput) (*elasticloadbalancingv2.DeregisterTargetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterTargetsWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DeregisterTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeregisterTargetsWithContext indicates an expected call of DeregisterTargetsWithContext.
func (mr *MockELBV2MockRecorder) DeregisterTargetsWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterTargetsWithContext", reflect.TypeOf((*MockELBV2)(nil).DeregisterTargetsWithContext), arg0, arg1)
}

// DescribeCapacityReservationWithContext mocks base method.
func (m *MockELBV2) DescribeCapacityReservationWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeCapacityReservationInput) (*elasticloadbalancingv2.DescribeCapacityReservationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCapacityReservationWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeCapacityReservationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCapacityReservationWithContext indicates an expected call of DescribeCapacityReservationWithContext.
func (mr *MockELBV2MockRecorder) DescribeCapacityReservationWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCapacityReservationWithContext", reflect.TypeOf((*MockELBV2)(nil).DescribeCapacityReservationWithContext), arg0, arg1)
}

// DescribeListenerAttributesWithContext mocks base method.
func (m *MockELBV2) DescribeListenerAttributesWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeListenerAttributesInput) (*elasticloadbalancingv2.DescribeListenerAttributesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeListenerAttributesWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeListenerAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeListenerAttributesWithContext indicates an expected call of DescribeListenerAttributesWithContext.
func (mr *MockELBV2MockRecorder) DescribeListenerAttributesWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeListenerAttributesWithContext", reflect.TypeOf((*MockELBV2)(nil).DescribeListenerAttributesWithContext), arg0, arg1)
}

// DescribeListenerCertificatesAsList mocks base method.
func (m *MockELBV2) DescribeListenerCertificatesAsList(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeListenerCertificatesInput) ([]types.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeListenerCertificatesAsList", arg0, arg1)
	ret0, _ := ret[0].([]types.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeListenerCertificatesAsList indicates an expected call of DescribeListenerCertificatesAsList.
func (mr *MockELBV2MockRecorder) DescribeListenerCertificatesAsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeListenerCertificatesAsList", reflect.TypeOf((*MockELBV2)(nil).DescribeListenerCertificatesAsList), arg0, arg1)
}

// DescribeListenersAsList mocks base method.
func (m *MockELBV2) DescribeListenersAsList(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeListenersInput) ([]types.Listener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeListenersAsList", arg0, arg1)
	ret0, _ := ret[0].([]types.Listener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeListenersAsList indicates an expected call of DescribeListenersAsList.
func (mr *MockELBV2MockRecorder) DescribeListenersAsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeListenersAsList", reflect.TypeOf((*MockELBV2)(nil).DescribeListenersAsList), arg0, arg1)
}

// DescribeListenersWithContext mocks base method.
func (m *MockELBV2) DescribeListenersWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeListenersInput) (*elasticloadbalancingv2.DescribeListenersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeListenersWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeListenersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeListenersWithContext indicates an expected call of DescribeListenersWithContext.
func (mr *MockELBV2MockRecorder) DescribeListenersWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeListenersWithContext", reflect.TypeOf((*MockELBV2)(nil).DescribeListenersWithContext), arg0, arg1)
}

// DescribeLoadBalancerAttributesWithContext mocks base method.
func (m *MockELBV2) DescribeLoadBalancerAttributesWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeLoadBalancerAttributesInput) (*elasticloadbalancingv2.DescribeLoadBalancerAttributesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLoadBalancerAttributesWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeLoadBalancerAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancerAttributesWithContext indicates an expected call of DescribeLoadBalancerAttributesWithContext.
func (mr *MockELBV2MockRecorder) DescribeLoadBalancerAttributesWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancerAttributesWithContext", reflect.TypeOf((*MockELBV2)(nil).DescribeLoadBalancerAttributesWithContext), arg0, arg1)
}

// DescribeLoadBalancersAsList mocks base method.
func (m *MockELBV2) DescribeLoadBalancersAsList(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeLoadBalancersInput) ([]types.LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLoadBalancersAsList", arg0, arg1)
	ret0, _ := ret[0].([]types.LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancersAsList indicates an expected call of DescribeLoadBalancersAsList.
func (mr *MockELBV2MockRecorder) DescribeLoadBalancersAsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancersAsList", reflect.TypeOf((*MockELBV2)(nil).DescribeLoadBalancersAsList), arg0, arg1)
}

// DescribeLoadBalancersWithContext mocks base method.
func (m *MockELBV2) DescribeLoadBalancersWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeLoadBalancersInput) (*elasticloadbalancingv2.DescribeLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLoadBalancersWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancersWithContext indicates an expected call of DescribeLoadBalancersWithContext.
func (mr *MockELBV2MockRecorder) DescribeLoadBalancersWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancersWithContext", reflect.TypeOf((*MockELBV2)(nil).DescribeLoadBalancersWithContext), arg0, arg1)
}

// DescribeRulesAsList mocks base method.
func (m *MockELBV2) DescribeRulesAsList(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeRulesInput) ([]types.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRulesAsList", arg0, arg1)
	ret0, _ := ret[0].([]types.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRulesAsList indicates an expected call of DescribeRulesAsList.
func (mr *MockELBV2MockRecorder) DescribeRulesAsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRulesAsList", reflect.TypeOf((*MockELBV2)(nil).DescribeRulesAsList), arg0, arg1)
}

// DescribeRulesWithContext mocks base method.
func (m *MockELBV2) DescribeRulesWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeRulesInput) (*elasticloadbalancingv2.DescribeRulesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRulesWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRulesWithContext indicates an expected call of DescribeRulesWithContext.
func (mr *MockELBV2MockRecorder) DescribeRulesWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRulesWithContext", reflect.TypeOf((*MockELBV2)(nil).DescribeRulesWithContext), arg0, arg1)
}

// DescribeTagsWithContext mocks base method.
func (m *MockELBV2) DescribeTagsWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTagsInput) (*elasticloadbalancingv2.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTagsWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTagsWithContext indicates an expected call of DescribeTagsWithContext.
func (mr *MockELBV2MockRecorder) DescribeTagsWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTagsWithContext", reflect.TypeOf((*MockELBV2)(nil).DescribeTagsWithContext), arg0, arg1)
}

// DescribeTargetGroupAttributesWithContext mocks base method.
func (m *MockELBV2) DescribeTargetGroupAttributesWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTargetGroupAttributesInput) (*elasticloadbalancingv2.DescribeTargetGroupAttributesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTargetGroupAttributesWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeTargetGroupAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTargetGroupAttributesWithContext indicates an expected call of DescribeTargetGroupAttributesWithContext.
func (mr *MockELBV2MockRecorder) DescribeTargetGroupAttributesWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTargetGroupAttributesWithContext", reflect.TypeOf((*MockELBV2)(nil).DescribeTargetGroupAttributesWithContext), arg0, arg1)
}

// DescribeTargetGroupsAsList mocks base method.
func (m *MockELBV2) DescribeTargetGroupsAsList(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTargetGroupsInput) ([]types.TargetGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTargetGroupsAsList", arg0, arg1)
	ret0, _ := ret[0].([]types.TargetGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTargetGroupsAsList indicates an expected call of DescribeTargetGroupsAsList.
func (mr *MockELBV2MockRecorder) DescribeTargetGroupsAsList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTargetGroupsAsList", reflect.TypeOf((*MockELBV2)(nil).DescribeTargetGroupsAsList), arg0, arg1)
}

// DescribeTargetGroupsWithContext mocks base method.
func (m *MockELBV2) DescribeTargetGroupsWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTargetGroupsInput) (*elasticloadbalancingv2.DescribeTargetGroupsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTargetGroupsWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeTargetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTargetGroupsWithContext indicates an expected call of DescribeTargetGroupsWithContext.
func (mr *MockELBV2MockRecorder) DescribeTargetGroupsWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTargetGroupsWithContext", reflect.TypeOf((*MockELBV2)(nil).DescribeTargetGroupsWithContext), arg0, arg1)
}

// DescribeTargetHealthWithContext mocks base method.
func (m *MockELBV2) DescribeTargetHealthWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTargetHealthInput) (*elasticloadbalancingv2.DescribeTargetHealthOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTargetHealthWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeTargetHealthOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTargetHealthWithContext indicates an expected call of DescribeTargetHealthWithContext.
func (mr *MockELBV2MockRecorder) DescribeTargetHealthWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTargetHealthWithContext", reflect.TypeOf((*MockELBV2)(nil).DescribeTargetHealthWithContext), arg0, arg1)
}

// DescribeTrustStoresWithContext mocks base method.
func (m *MockELBV2) DescribeTrustStoresWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTrustStoresInput) (*elasticloadbalancingv2.DescribeTrustStoresOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTrustStoresWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeTrustStoresOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTrustStoresWithContext indicates an expected call of DescribeTrustStoresWithContext.
func (mr *MockELBV2MockRecorder) DescribeTrustStoresWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTrustStoresWithContext", reflect.TypeOf((*MockELBV2)(nil).DescribeTrustStoresWithContext), arg0, arg1)
}

// ModifyCapacityReservationWithContext mocks base method.
func (m *MockELBV2) ModifyCapacityReservationWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.ModifyCapacityReservationInput) (*elasticloadbalancingv2.ModifyCapacityReservationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyCapacityReservationWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.ModifyCapacityReservationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyCapacityReservationWithContext indicates an expected call of ModifyCapacityReservationWithContext.
func (mr *MockELBV2MockRecorder) ModifyCapacityReservationWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyCapacityReservationWithContext", reflect.TypeOf((*MockELBV2)(nil).ModifyCapacityReservationWithContext), arg0, arg1)
}

// ModifyListenerAttributesWithContext mocks base method.
func (m *MockELBV2) ModifyListenerAttributesWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.ModifyListenerAttributesInput) (*elasticloadbalancingv2.ModifyListenerAttributesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyListenerAttributesWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.ModifyListenerAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyListenerAttributesWithContext indicates an expected call of ModifyListenerAttributesWithContext.
func (mr *MockELBV2MockRecorder) ModifyListenerAttributesWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyListenerAttributesWithContext", reflect.TypeOf((*MockELBV2)(nil).ModifyListenerAttributesWithContext), arg0, arg1)
}

// ModifyListenerWithContext mocks base method.
func (m *MockELBV2) ModifyListenerWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.ModifyListenerInput) (*elasticloadbalancingv2.ModifyListenerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyListenerWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.ModifyListenerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyListenerWithContext indicates an expected call of ModifyListenerWithContext.
func (mr *MockELBV2MockRecorder) ModifyListenerWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyListenerWithContext", reflect.TypeOf((*MockELBV2)(nil).ModifyListenerWithContext), arg0, arg1)
}

// ModifyLoadBalancerAttributesWithContext mocks base method.
func (m *MockELBV2) ModifyLoadBalancerAttributesWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.ModifyLoadBalancerAttributesInput) (*elasticloadbalancingv2.ModifyLoadBalancerAttributesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyLoadBalancerAttributesWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.ModifyLoadBalancerAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyLoadBalancerAttributesWithContext indicates an expected call of ModifyLoadBalancerAttributesWithContext.
func (mr *MockELBV2MockRecorder) ModifyLoadBalancerAttributesWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyLoadBalancerAttributesWithContext", reflect.TypeOf((*MockELBV2)(nil).ModifyLoadBalancerAttributesWithContext), arg0, arg1)
}

// ModifyRuleWithContext mocks base method.
func (m *MockELBV2) ModifyRuleWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.ModifyRuleInput) (*elasticloadbalancingv2.ModifyRuleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyRuleWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.ModifyRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyRuleWithContext indicates an expected call of ModifyRuleWithContext.
func (mr *MockELBV2MockRecorder) ModifyRuleWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyRuleWithContext", reflect.TypeOf((*MockELBV2)(nil).ModifyRuleWithContext), arg0, arg1)
}

// ModifyTargetGroupAttributesWithContext mocks base method.
func (m *MockELBV2) ModifyTargetGroupAttributesWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.ModifyTargetGroupAttributesInput) (*elasticloadbalancingv2.ModifyTargetGroupAttributesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyTargetGroupAttributesWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.ModifyTargetGroupAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyTargetGroupAttributesWithContext indicates an expected call of ModifyTargetGroupAttributesWithContext.
func (mr *MockELBV2MockRecorder) ModifyTargetGroupAttributesWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyTargetGroupAttributesWithContext", reflect.TypeOf((*MockELBV2)(nil).ModifyTargetGroupAttributesWithContext), arg0, arg1)
}

// ModifyTargetGroupWithContext mocks base method.
func (m *MockELBV2) ModifyTargetGroupWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.ModifyTargetGroupInput) (*elasticloadbalancingv2.ModifyTargetGroupOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyTargetGroupWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.ModifyTargetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyTargetGroupWithContext indicates an expected call of ModifyTargetGroupWithContext.
func (mr *MockELBV2MockRecorder) ModifyTargetGroupWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyTargetGroupWithContext", reflect.TypeOf((*MockELBV2)(nil).ModifyTargetGroupWithContext), arg0, arg1)
}

// RegisterTargetsWithContext mocks base method.
func (m *MockELBV2) RegisterTargetsWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.RegisterTargetsInput) (*elasticloadbalancingv2.RegisterTargetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterTargetsWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.RegisterTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterTargetsWithContext indicates an expected call of RegisterTargetsWithContext.
func (mr *MockELBV2MockRecorder) RegisterTargetsWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTargetsWithContext", reflect.TypeOf((*MockELBV2)(nil).RegisterTargetsWithContext), arg0, arg1)
}

// RemoveListenerCertificatesWithContext mocks base method.
func (m *MockELBV2) RemoveListenerCertificatesWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.RemoveListenerCertificatesInput) (*elasticloadbalancingv2.RemoveListenerCertificatesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveListenerCertificatesWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.RemoveListenerCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveListenerCertificatesWithContext indicates an expected call of RemoveListenerCertificatesWithContext.
func (mr *MockELBV2MockRecorder) RemoveListenerCertificatesWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveListenerCertificatesWithContext", reflect.TypeOf((*MockELBV2)(nil).RemoveListenerCertificatesWithContext), arg0, arg1)
}

// RemoveTagsWithContext mocks base method.
func (m *MockELBV2) RemoveTagsWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.RemoveTagsInput) (*elasticloadbalancingv2.RemoveTagsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTagsWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.RemoveTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveTagsWithContext indicates an expected call of RemoveTagsWithContext.
func (mr *MockELBV2MockRecorder) RemoveTagsWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTagsWithContext", reflect.TypeOf((*MockELBV2)(nil).RemoveTagsWithContext), arg0, arg1)
}

// SetIpAddressTypeWithContext mocks base method.
func (m *MockELBV2) SetIpAddressTypeWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.SetIpAddressTypeInput) (*elasticloadbalancingv2.SetIpAddressTypeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetIpAddressTypeWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.SetIpAddressTypeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetIpAddressTypeWithContext indicates an expected call of SetIpAddressTypeWithContext.
func (mr *MockELBV2MockRecorder) SetIpAddressTypeWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIpAddressTypeWithContext", reflect.TypeOf((*MockELBV2)(nil).SetIpAddressTypeWithContext), arg0, arg1)
}

// SetSecurityGroupsWithContext mocks base method.
func (m *MockELBV2) SetSecurityGroupsWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.SetSecurityGroupsInput) (*elasticloadbalancingv2.SetSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSecurityGroupsWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.SetSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSecurityGroupsWithContext indicates an expected call of SetSecurityGroupsWithContext.
func (mr *MockELBV2MockRecorder) SetSecurityGroupsWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecurityGroupsWithContext", reflect.TypeOf((*MockELBV2)(nil).SetSecurityGroupsWithContext), arg0, arg1)
}

// SetSubnetsWithContext mocks base method.
func (m *MockELBV2) SetSubnetsWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.SetSubnetsInput) (*elasticloadbalancingv2.SetSubnetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSubnetsWithContext", arg0, arg1)
	ret0, _ := ret[0].(*elasticloadbalancingv2.SetSubnetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSubnetsWithContext indicates an expected call of SetSubnetsWithContext.
func (mr *MockELBV2MockRecorder) SetSubnetsWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSubnetsWithContext", reflect.TypeOf((*MockELBV2)(nil).SetSubnetsWithContext), arg0, arg1)
}

// WaitUntilLoadBalancerAvailableWithContext mocks base method.
func (m *MockELBV2) WaitUntilLoadBalancerAvailableWithContext(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeLoadBalancersInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitUntilLoadBalancerAvailableWithContext", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilLoadBalancerAvailableWithContext indicates an expected call of WaitUntilLoadBalancerAvailableWithContext.
func (mr *MockELBV2MockRecorder) WaitUntilLoadBalancerAvailableWithContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilLoadBalancerAvailableWithContext", reflect.TypeOf((*MockELBV2)(nil).WaitUntilLoadBalancerAvailableWithContext), arg0, arg1)
}
