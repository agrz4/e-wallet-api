package mocks

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/mock"
)

type JWTService struct {
	mock.Mock
}

func (_m *JWTService) GenerateToken(userID int) (string, error) {
	ret := _m.Called(userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *JWTService) ValidateToken(token string) (*jwt.Token, error) {
	ret := _m.Called(token)

	var r0 *jwt.Token
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewJWTService interface {
	mock.TestingT
	Cleanup(func())
}

func NewJWTService(t mockConstructorTestingTNewJWTService) *JWTService {
	mock := &JWTService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
