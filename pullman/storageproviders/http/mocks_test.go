// Copyright 2021 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: ./provider.go
package httpprovider

import (
	context "context"
	tls "crypto/tls"
	x509 "crypto/x509"
	http "net/http"
	reflect "reflect"

	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
)

// MockfetcherFactory is a mock of fetcherFactory interface.
type MockfetcherFactory struct {
	ctrl     *gomock.Controller
	recorder *MockfetcherFactoryMockRecorder
}

// MockfetcherFactoryMockRecorder is the mock recorder for MockfetcherFactory.
type MockfetcherFactoryMockRecorder struct {
	mock *MockfetcherFactory
}

// NewMockfetcherFactory creates a new mock instance.
func NewMockfetcherFactory(ctrl *gomock.Controller) *MockfetcherFactory {
	mock := &MockfetcherFactory{ctrl: ctrl}
	mock.recorder = &MockfetcherFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockfetcherFactory) EXPECT() *MockfetcherFactoryMockRecorder {
	return m.recorder
}

// newClient mocks base method.
func (m *MockfetcherFactory) newClient(log logr.Logger, ca *x509.CertPool, client_tls *tls.Certificate) fetcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "newClient", log, ca, client_tls)
	ret0, _ := ret[0].(fetcher)
	return ret0
}

// newClient indicates an expected call of newClient.
func (mr *MockfetcherFactoryMockRecorder) newClient(log, ca, client_tls interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "newClient", reflect.TypeOf((*MockfetcherFactory)(nil).newClient), log, ca, client_tls)
}

// Mockfetcher is a mock of fetcher interface.
type Mockfetcher struct {
	ctrl     *gomock.Controller
	recorder *MockfetcherMockRecorder
}

// MockfetcherMockRecorder is the mock recorder for Mockfetcher.
type MockfetcherMockRecorder struct {
	mock *Mockfetcher
}

// NewMockfetcher creates a new mock instance.
func NewMockfetcher(ctrl *gomock.Controller) *Mockfetcher {
	mock := &Mockfetcher{ctrl: ctrl}
	mock.recorder = &MockfetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockfetcher) EXPECT() *MockfetcherMockRecorder {
	return m.recorder
}

// download mocks base method.
func (m *Mockfetcher) download(ctx context.Context, req *http.Request, filename string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "download", ctx, req, filename)
	ret0, _ := ret[0].(error)
	return ret0
}

// download indicates an expected call of download.
func (mr *MockfetcherMockRecorder) download(ctx, req, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "download", reflect.TypeOf((*Mockfetcher)(nil).download), ctx, req, filename)
}