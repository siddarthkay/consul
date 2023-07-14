// Code generated by mockery v2.20.0. DO NOT EDIT.

package ca

import (
	x509 "crypto/x509"

	mock "github.com/stretchr/testify/mock"
)

// MockProvider is an autogenerated mock type for the Provider type
type MockProvider struct {
	mock.Mock
}

// ActiveLeafSigningCert provides a mock function with given fields:
func (_m *MockProvider) ActiveLeafSigningCert() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Cleanup provides a mock function with given fields: providerTypeChange, otherConfig
func (_m *MockProvider) Cleanup(providerTypeChange bool, otherConfig map[string]interface{}) error {
	ret := _m.Called(providerTypeChange, otherConfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, map[string]interface{}) error); ok {
		r0 = rf(providerTypeChange, otherConfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Configure provides a mock function with given fields: cfg
func (_m *MockProvider) Configure(cfg ProviderConfig) error {
	ret := _m.Called(cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(ProviderConfig) error); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CrossSignCA provides a mock function with given fields: _a0
func (_m *MockProvider) CrossSignCA(_a0 *x509.Certificate) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*x509.Certificate) (string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*x509.Certificate) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*x509.Certificate) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateCAChain provides a mock function with given fields:
func (_m *MockProvider) GenerateCAChain() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateIntermediateCSR provides a mock function with given fields:
func (_m *MockProvider) GenerateIntermediateCSR() (string, string, error) {
	ret := _m.Called()

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func() (string, string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetIntermediate provides a mock function with given fields: intermediatePEM, rootPEM, opaque
func (_m *MockProvider) SetIntermediate(intermediatePEM string, rootPEM string, opaque string) error {
	ret := _m.Called(intermediatePEM, rootPEM, opaque)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(intermediatePEM, rootPEM, opaque)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sign provides a mock function with given fields: _a0
func (_m *MockProvider) Sign(_a0 *x509.CertificateRequest) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*x509.CertificateRequest) (string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*x509.CertificateRequest) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*x509.CertificateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignIntermediate provides a mock function with given fields: _a0
func (_m *MockProvider) SignIntermediate(_a0 *x509.CertificateRequest) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*x509.CertificateRequest) (string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*x509.CertificateRequest) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*x509.CertificateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// State provides a mock function with given fields:
func (_m *MockProvider) State() (map[string]string, error) {
	ret := _m.Called()

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func() (map[string]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SupportsCrossSigning provides a mock function with given fields:
func (_m *MockProvider) SupportsCrossSigning() (bool, error) {
	ret := _m.Called()

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func() (bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockProvider creates a new instance of MockProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockProvider(t mockConstructorTestingTNewMockProvider) *MockProvider {
	mock := &MockProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
