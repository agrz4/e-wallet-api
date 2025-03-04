package mocks

import (
	"e-wallet-api-go/internal/dto"
	"e-wallet-api-go/internal/model"

	"github.com/stretchr/testify/mock"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: userID
func (_m *TransactionRepository) Count(userID int) (int64, error) {
	ret := _m.Called(userID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int) int64); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: userID, query
func (_m *TransactionRepository) FindAll(userID int, query *dto.TransactionRequestQuery) ([]*model.Transaction, error) {
	ret := _m.Called(userID, query)

	var r0 []*model.Transaction
	if rf, ok := ret.Get(0).(func(int, *dto.TransactionRequestQuery) []*model.Transaction); ok {
		r0 = rf(userID, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *dto.TransactionRequestQuery) error); ok {
		r1 = rf(userID, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: transaction
func (_m *TransactionRepository) Save(transaction *model.Transaction) (*model.Transaction, error) {
	ret := _m.Called(transaction)

	var r0 *model.Transaction
	if rf, ok := ret.Get(0).(func(*model.Transaction) *model.Transaction); ok {
		r0 = rf(transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Transaction) error); ok {
		r1 = rf(transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionRepository(t mockConstructorTestingTNewTransactionRepository) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
