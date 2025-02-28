// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extension-provider-gcp/pkg/gcp/client (interfaces: Factory,DNSClient,ComputeClient,StorageClient)
//
// Generated by this command:
//
//	mockgen -package client -destination=mocks.go github.com/gardener/gardener-extension-provider-gcp/pkg/gcp/client Factory,DNSClient,ComputeClient,StorageClient
//

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	storage "cloud.google.com/go/storage"
	client "github.com/gardener/gardener-extension-provider-gcp/pkg/gcp/client"
	gomock "go.uber.org/mock/gomock"
	compute "google.golang.org/api/compute/v1"
	v1 "k8s.io/api/core/v1"
	client0 "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
	isgomock struct{}
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// Compute mocks base method.
func (m *MockFactory) Compute(arg0 context.Context, arg1 client0.Client, arg2 v1.SecretReference) (client.ComputeClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compute", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.ComputeClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compute indicates an expected call of Compute.
func (mr *MockFactoryMockRecorder) Compute(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compute", reflect.TypeOf((*MockFactory)(nil).Compute), arg0, arg1, arg2)
}

// DNS mocks base method.
func (m *MockFactory) DNS(arg0 context.Context, arg1 client0.Client, arg2 v1.SecretReference) (client.DNSClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DNS", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.DNSClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DNS indicates an expected call of DNS.
func (mr *MockFactoryMockRecorder) DNS(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DNS", reflect.TypeOf((*MockFactory)(nil).DNS), arg0, arg1, arg2)
}

// IAM mocks base method.
func (m *MockFactory) IAM(arg0 context.Context, arg1 client0.Client, arg2 v1.SecretReference) (client.IAMClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IAM", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.IAMClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IAM indicates an expected call of IAM.
func (mr *MockFactoryMockRecorder) IAM(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IAM", reflect.TypeOf((*MockFactory)(nil).IAM), arg0, arg1, arg2)
}

// Storage mocks base method.
func (m *MockFactory) Storage(arg0 context.Context, arg1 client0.Client, arg2 v1.SecretReference) (client.StorageClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Storage", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.StorageClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Storage indicates an expected call of Storage.
func (mr *MockFactoryMockRecorder) Storage(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockFactory)(nil).Storage), arg0, arg1, arg2)
}

// MockDNSClient is a mock of DNSClient interface.
type MockDNSClient struct {
	ctrl     *gomock.Controller
	recorder *MockDNSClientMockRecorder
	isgomock struct{}
}

// MockDNSClientMockRecorder is the mock recorder for MockDNSClient.
type MockDNSClientMockRecorder struct {
	mock *MockDNSClient
}

// NewMockDNSClient creates a new mock instance.
func NewMockDNSClient(ctrl *gomock.Controller) *MockDNSClient {
	mock := &MockDNSClient{ctrl: ctrl}
	mock.recorder = &MockDNSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDNSClient) EXPECT() *MockDNSClientMockRecorder {
	return m.recorder
}

// CreateOrUpdateRecordSet mocks base method.
func (m *MockDNSClient) CreateOrUpdateRecordSet(ctx context.Context, managedZone, name, recordType string, rrdatas []string, ttl int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateRecordSet", ctx, managedZone, name, recordType, rrdatas, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateRecordSet indicates an expected call of CreateOrUpdateRecordSet.
func (mr *MockDNSClientMockRecorder) CreateOrUpdateRecordSet(ctx, managedZone, name, recordType, rrdatas, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateRecordSet", reflect.TypeOf((*MockDNSClient)(nil).CreateOrUpdateRecordSet), ctx, managedZone, name, recordType, rrdatas, ttl)
}

// DeleteRecordSet mocks base method.
func (m *MockDNSClient) DeleteRecordSet(ctx context.Context, managedZone, name, recordType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecordSet", ctx, managedZone, name, recordType)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecordSet indicates an expected call of DeleteRecordSet.
func (mr *MockDNSClientMockRecorder) DeleteRecordSet(ctx, managedZone, name, recordType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecordSet", reflect.TypeOf((*MockDNSClient)(nil).DeleteRecordSet), ctx, managedZone, name, recordType)
}

// GetManagedZones mocks base method.
func (m *MockDNSClient) GetManagedZones(ctx context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedZones", ctx)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagedZones indicates an expected call of GetManagedZones.
func (mr *MockDNSClientMockRecorder) GetManagedZones(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedZones", reflect.TypeOf((*MockDNSClient)(nil).GetManagedZones), ctx)
}

// MockComputeClient is a mock of ComputeClient interface.
type MockComputeClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeClientMockRecorder
	isgomock struct{}
}

// MockComputeClientMockRecorder is the mock recorder for MockComputeClient.
type MockComputeClientMockRecorder struct {
	mock *MockComputeClient
}

// NewMockComputeClient creates a new mock instance.
func NewMockComputeClient(ctrl *gomock.Controller) *MockComputeClient {
	mock := &MockComputeClient{ctrl: ctrl}
	mock.recorder = &MockComputeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputeClient) EXPECT() *MockComputeClientMockRecorder {
	return m.recorder
}

// DeleteDisk mocks base method.
func (m *MockComputeClient) DeleteDisk(ctx context.Context, zone, diskName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDisk", ctx, zone, diskName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDisk indicates an expected call of DeleteDisk.
func (mr *MockComputeClientMockRecorder) DeleteDisk(ctx, zone, diskName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDisk", reflect.TypeOf((*MockComputeClient)(nil).DeleteDisk), ctx, zone, diskName)
}

// DeleteFirewallRule mocks base method.
func (m *MockComputeClient) DeleteFirewallRule(ctx context.Context, firewall string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFirewallRule", ctx, firewall)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFirewallRule indicates an expected call of DeleteFirewallRule.
func (mr *MockComputeClientMockRecorder) DeleteFirewallRule(ctx, firewall any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFirewallRule", reflect.TypeOf((*MockComputeClient)(nil).DeleteFirewallRule), ctx, firewall)
}

// DeleteInstance mocks base method.
func (m *MockComputeClient) DeleteInstance(ctx context.Context, zone, instanceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInstance", ctx, zone, instanceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInstance indicates an expected call of DeleteInstance.
func (mr *MockComputeClientMockRecorder) DeleteInstance(ctx, zone, instanceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstance", reflect.TypeOf((*MockComputeClient)(nil).DeleteInstance), ctx, zone, instanceName)
}

// DeleteNetwork mocks base method.
func (m *MockComputeClient) DeleteNetwork(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetwork", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNetwork indicates an expected call of DeleteNetwork.
func (mr *MockComputeClientMockRecorder) DeleteNetwork(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetwork", reflect.TypeOf((*MockComputeClient)(nil).DeleteNetwork), ctx, id)
}

// DeleteRoute mocks base method.
func (m *MockComputeClient) DeleteRoute(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoute", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoute indicates an expected call of DeleteRoute.
func (mr *MockComputeClientMockRecorder) DeleteRoute(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoute", reflect.TypeOf((*MockComputeClient)(nil).DeleteRoute), ctx, id)
}

// DeleteRouter mocks base method.
func (m *MockComputeClient) DeleteRouter(ctx context.Context, region, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRouter", ctx, region, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRouter indicates an expected call of DeleteRouter.
func (mr *MockComputeClientMockRecorder) DeleteRouter(ctx, region, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRouter", reflect.TypeOf((*MockComputeClient)(nil).DeleteRouter), ctx, region, id)
}

// DeleteSubnet mocks base method.
func (m *MockComputeClient) DeleteSubnet(ctx context.Context, region, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubnet", ctx, region, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubnet indicates an expected call of DeleteSubnet.
func (mr *MockComputeClientMockRecorder) DeleteSubnet(ctx, region, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubnet", reflect.TypeOf((*MockComputeClient)(nil).DeleteSubnet), ctx, region, id)
}

// ExpandSubnet mocks base method.
func (m *MockComputeClient) ExpandSubnet(ctx context.Context, region, id, cidr string) (*compute.Subnetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpandSubnet", ctx, region, id, cidr)
	ret0, _ := ret[0].(*compute.Subnetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExpandSubnet indicates an expected call of ExpandSubnet.
func (mr *MockComputeClientMockRecorder) ExpandSubnet(ctx, region, id, cidr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpandSubnet", reflect.TypeOf((*MockComputeClient)(nil).ExpandSubnet), ctx, region, id, cidr)
}

// GetAddress mocks base method.
func (m *MockComputeClient) GetAddress(ctx context.Context, region, name string) (*compute.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress", ctx, region, name)
	ret0, _ := ret[0].(*compute.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddress indicates an expected call of GetAddress.
func (mr *MockComputeClientMockRecorder) GetAddress(ctx, region, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockComputeClient)(nil).GetAddress), ctx, region, name)
}

// GetDisk mocks base method.
func (m *MockComputeClient) GetDisk(ctx context.Context, zone, instanceName string) (*compute.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDisk", ctx, zone, instanceName)
	ret0, _ := ret[0].(*compute.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDisk indicates an expected call of GetDisk.
func (mr *MockComputeClientMockRecorder) GetDisk(ctx, zone, instanceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDisk", reflect.TypeOf((*MockComputeClient)(nil).GetDisk), ctx, zone, instanceName)
}

// GetExternalAddresses mocks base method.
func (m *MockComputeClient) GetExternalAddresses(ctx context.Context, region string) (map[string][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalAddresses", ctx, region)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExternalAddresses indicates an expected call of GetExternalAddresses.
func (mr *MockComputeClientMockRecorder) GetExternalAddresses(ctx, region any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalAddresses", reflect.TypeOf((*MockComputeClient)(nil).GetExternalAddresses), ctx, region)
}

// GetFirewallRule mocks base method.
func (m *MockComputeClient) GetFirewallRule(ctx context.Context, firewall string) (*compute.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirewallRule", ctx, firewall)
	ret0, _ := ret[0].(*compute.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirewallRule indicates an expected call of GetFirewallRule.
func (mr *MockComputeClientMockRecorder) GetFirewallRule(ctx, firewall any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirewallRule", reflect.TypeOf((*MockComputeClient)(nil).GetFirewallRule), ctx, firewall)
}

// GetInstance mocks base method.
func (m *MockComputeClient) GetInstance(ctx context.Context, zone, instanceName string) (*compute.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstance", ctx, zone, instanceName)
	ret0, _ := ret[0].(*compute.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstance indicates an expected call of GetInstance.
func (mr *MockComputeClientMockRecorder) GetInstance(ctx, zone, instanceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstance", reflect.TypeOf((*MockComputeClient)(nil).GetInstance), ctx, zone, instanceName)
}

// GetNetwork mocks base method.
func (m *MockComputeClient) GetNetwork(ctx context.Context, id string) (*compute.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetwork", ctx, id)
	ret0, _ := ret[0].(*compute.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetwork indicates an expected call of GetNetwork.
func (mr *MockComputeClientMockRecorder) GetNetwork(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetwork", reflect.TypeOf((*MockComputeClient)(nil).GetNetwork), ctx, id)
}

// GetRegion mocks base method.
func (m *MockComputeClient) GetRegion(ctx context.Context, region string) (*compute.Region, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegion", ctx, region)
	ret0, _ := ret[0].(*compute.Region)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegion indicates an expected call of GetRegion.
func (mr *MockComputeClientMockRecorder) GetRegion(ctx, region any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegion", reflect.TypeOf((*MockComputeClient)(nil).GetRegion), ctx, region)
}

// GetRouter mocks base method.
func (m *MockComputeClient) GetRouter(ctx context.Context, region, id string) (*compute.Router, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouter", ctx, region, id)
	ret0, _ := ret[0].(*compute.Router)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRouter indicates an expected call of GetRouter.
func (mr *MockComputeClientMockRecorder) GetRouter(ctx, region, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouter", reflect.TypeOf((*MockComputeClient)(nil).GetRouter), ctx, region, id)
}

// GetSubnet mocks base method.
func (m *MockComputeClient) GetSubnet(ctx context.Context, region, id string) (*compute.Subnetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnet", ctx, region, id)
	ret0, _ := ret[0].(*compute.Subnetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnet indicates an expected call of GetSubnet.
func (mr *MockComputeClientMockRecorder) GetSubnet(ctx, region, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnet", reflect.TypeOf((*MockComputeClient)(nil).GetSubnet), ctx, region, id)
}

// InsertDisk mocks base method.
func (m *MockComputeClient) InsertDisk(ctx context.Context, zone string, disk *compute.Disk) (*compute.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertDisk", ctx, zone, disk)
	ret0, _ := ret[0].(*compute.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertDisk indicates an expected call of InsertDisk.
func (mr *MockComputeClientMockRecorder) InsertDisk(ctx, zone, disk any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertDisk", reflect.TypeOf((*MockComputeClient)(nil).InsertDisk), ctx, zone, disk)
}

// InsertFirewallRule mocks base method.
func (m *MockComputeClient) InsertFirewallRule(ctx context.Context, firewall *compute.Firewall) (*compute.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertFirewallRule", ctx, firewall)
	ret0, _ := ret[0].(*compute.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertFirewallRule indicates an expected call of InsertFirewallRule.
func (mr *MockComputeClientMockRecorder) InsertFirewallRule(ctx, firewall any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertFirewallRule", reflect.TypeOf((*MockComputeClient)(nil).InsertFirewallRule), ctx, firewall)
}

// InsertInstance mocks base method.
func (m *MockComputeClient) InsertInstance(ctx context.Context, zone string, instance *compute.Instance) (*compute.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertInstance", ctx, zone, instance)
	ret0, _ := ret[0].(*compute.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertInstance indicates an expected call of InsertInstance.
func (mr *MockComputeClientMockRecorder) InsertInstance(ctx, zone, instance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertInstance", reflect.TypeOf((*MockComputeClient)(nil).InsertInstance), ctx, zone, instance)
}

// InsertNetwork mocks base method.
func (m *MockComputeClient) InsertNetwork(ctx context.Context, nw *compute.Network) (*compute.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNetwork", ctx, nw)
	ret0, _ := ret[0].(*compute.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertNetwork indicates an expected call of InsertNetwork.
func (mr *MockComputeClientMockRecorder) InsertNetwork(ctx, nw any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNetwork", reflect.TypeOf((*MockComputeClient)(nil).InsertNetwork), ctx, nw)
}

// InsertRouter mocks base method.
func (m *MockComputeClient) InsertRouter(ctx context.Context, region string, router *compute.Router) (*compute.Router, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertRouter", ctx, region, router)
	ret0, _ := ret[0].(*compute.Router)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertRouter indicates an expected call of InsertRouter.
func (mr *MockComputeClientMockRecorder) InsertRouter(ctx, region, router any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertRouter", reflect.TypeOf((*MockComputeClient)(nil).InsertRouter), ctx, region, router)
}

// InsertSubnet mocks base method.
func (m *MockComputeClient) InsertSubnet(ctx context.Context, region string, subnet *compute.Subnetwork) (*compute.Subnetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertSubnet", ctx, region, subnet)
	ret0, _ := ret[0].(*compute.Subnetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertSubnet indicates an expected call of InsertSubnet.
func (mr *MockComputeClientMockRecorder) InsertSubnet(ctx, region, subnet any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertSubnet", reflect.TypeOf((*MockComputeClient)(nil).InsertSubnet), ctx, region, subnet)
}

// ListFirewallRules mocks base method.
func (m *MockComputeClient) ListFirewallRules(ctx context.Context, opts client.FirewallListOpts) ([]*compute.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFirewallRules", ctx, opts)
	ret0, _ := ret[0].([]*compute.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFirewallRules indicates an expected call of ListFirewallRules.
func (mr *MockComputeClientMockRecorder) ListFirewallRules(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFirewallRules", reflect.TypeOf((*MockComputeClient)(nil).ListFirewallRules), ctx, opts)
}

// ListImages mocks base method.
func (m *MockComputeClient) ListImages(ctx context.Context, imageName, orderBy, fields string) (*compute.ImageList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", ctx, imageName, orderBy, fields)
	ret0, _ := ret[0].(*compute.ImageList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockComputeClientMockRecorder) ListImages(ctx, imageName, orderBy, fields any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockComputeClient)(nil).ListImages), ctx, imageName, orderBy, fields)
}

// ListRoutes mocks base method.
func (m *MockComputeClient) ListRoutes(ctx context.Context, opts client.RouteListOpts) ([]*compute.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoutes", ctx, opts)
	ret0, _ := ret[0].([]*compute.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoutes indicates an expected call of ListRoutes.
func (mr *MockComputeClientMockRecorder) ListRoutes(ctx, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoutes", reflect.TypeOf((*MockComputeClient)(nil).ListRoutes), ctx, opts)
}

// PatchFirewallRule mocks base method.
func (m *MockComputeClient) PatchFirewallRule(ctx context.Context, name string, firewall *compute.Firewall) (*compute.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchFirewallRule", ctx, name, firewall)
	ret0, _ := ret[0].(*compute.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchFirewallRule indicates an expected call of PatchFirewallRule.
func (mr *MockComputeClientMockRecorder) PatchFirewallRule(ctx, name, firewall any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchFirewallRule", reflect.TypeOf((*MockComputeClient)(nil).PatchFirewallRule), ctx, name, firewall)
}

// PatchNetwork mocks base method.
func (m *MockComputeClient) PatchNetwork(ctx context.Context, id string, nw *compute.Network) (*compute.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchNetwork", ctx, id, nw)
	ret0, _ := ret[0].(*compute.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchNetwork indicates an expected call of PatchNetwork.
func (mr *MockComputeClientMockRecorder) PatchNetwork(ctx, id, nw any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchNetwork", reflect.TypeOf((*MockComputeClient)(nil).PatchNetwork), ctx, id, nw)
}

// PatchRouter mocks base method.
func (m *MockComputeClient) PatchRouter(ctx context.Context, region, id string, router *compute.Router) (*compute.Router, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchRouter", ctx, region, id, router)
	ret0, _ := ret[0].(*compute.Router)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchRouter indicates an expected call of PatchRouter.
func (mr *MockComputeClientMockRecorder) PatchRouter(ctx, region, id, router any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchRouter", reflect.TypeOf((*MockComputeClient)(nil).PatchRouter), ctx, region, id, router)
}

// PatchSubnet mocks base method.
func (m *MockComputeClient) PatchSubnet(ctx context.Context, region, id string, subnet *compute.Subnetwork) (*compute.Subnetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchSubnet", ctx, region, id, subnet)
	ret0, _ := ret[0].(*compute.Subnetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchSubnet indicates an expected call of PatchSubnet.
func (mr *MockComputeClientMockRecorder) PatchSubnet(ctx, region, id, subnet any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchSubnet", reflect.TypeOf((*MockComputeClient)(nil).PatchSubnet), ctx, region, id, subnet)
}

// WaitForIPv6Cidr mocks base method.
func (m *MockComputeClient) WaitForIPv6Cidr(ctx context.Context, region, subnetID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForIPv6Cidr", ctx, region, subnetID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForIPv6Cidr indicates an expected call of WaitForIPv6Cidr.
func (mr *MockComputeClientMockRecorder) WaitForIPv6Cidr(ctx, region, subnetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForIPv6Cidr", reflect.TypeOf((*MockComputeClient)(nil).WaitForIPv6Cidr), ctx, region, subnetID)
}

// MockStorageClient is a mock of StorageClient interface.
type MockStorageClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageClientMockRecorder
	isgomock struct{}
}

// MockStorageClientMockRecorder is the mock recorder for MockStorageClient.
type MockStorageClientMockRecorder struct {
	mock *MockStorageClient
}

// NewMockStorageClient creates a new mock instance.
func NewMockStorageClient(ctrl *gomock.Controller) *MockStorageClient {
	mock := &MockStorageClient{ctrl: ctrl}
	mock.recorder = &MockStorageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageClient) EXPECT() *MockStorageClientMockRecorder {
	return m.recorder
}

// Attrs mocks base method.
func (m *MockStorageClient) Attrs(ctx context.Context, bucketName string) (*storage.BucketAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attrs", ctx, bucketName)
	ret0, _ := ret[0].(*storage.BucketAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attrs indicates an expected call of Attrs.
func (mr *MockStorageClientMockRecorder) Attrs(ctx, bucketName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attrs", reflect.TypeOf((*MockStorageClient)(nil).Attrs), ctx, bucketName)
}

// CreateBucket mocks base method.
func (m *MockStorageClient) CreateBucket(ctx context.Context, atts *storage.BucketAttrs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", ctx, atts)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockStorageClientMockRecorder) CreateBucket(ctx, atts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockStorageClient)(nil).CreateBucket), ctx, atts)
}

// DeleteBucketIfExists mocks base method.
func (m *MockStorageClient) DeleteBucketIfExists(ctx context.Context, bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucketIfExists", ctx, bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucketIfExists indicates an expected call of DeleteBucketIfExists.
func (mr *MockStorageClientMockRecorder) DeleteBucketIfExists(ctx, bucketName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucketIfExists", reflect.TypeOf((*MockStorageClient)(nil).DeleteBucketIfExists), ctx, bucketName)
}

// DeleteObjectsWithPrefix mocks base method.
func (m *MockStorageClient) DeleteObjectsWithPrefix(ctx context.Context, bucketName, prefix string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjectsWithPrefix", ctx, bucketName, prefix)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObjectsWithPrefix indicates an expected call of DeleteObjectsWithPrefix.
func (mr *MockStorageClientMockRecorder) DeleteObjectsWithPrefix(ctx, bucketName, prefix any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectsWithPrefix", reflect.TypeOf((*MockStorageClient)(nil).DeleteObjectsWithPrefix), ctx, bucketName, prefix)
}

// LockBucket mocks base method.
func (m *MockStorageClient) LockBucket(ctx context.Context, bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockBucket", ctx, bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockBucket indicates an expected call of LockBucket.
func (mr *MockStorageClientMockRecorder) LockBucket(ctx, bucketName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockBucket", reflect.TypeOf((*MockStorageClient)(nil).LockBucket), ctx, bucketName)
}

// UpdateBucket mocks base method.
func (m *MockStorageClient) UpdateBucket(ctx context.Context, bucketName string, bucketAttrsToUpdate storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBucket", ctx, bucketName, bucketAttrsToUpdate)
	ret0, _ := ret[0].(*storage.BucketAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBucket indicates an expected call of UpdateBucket.
func (mr *MockStorageClientMockRecorder) UpdateBucket(ctx, bucketName, bucketAttrsToUpdate any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBucket", reflect.TypeOf((*MockStorageClient)(nil).UpdateBucket), ctx, bucketName, bucketAttrsToUpdate)
}
