// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	model "api-bed-covid/model"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Redis is an autogenerated mock type for the Redis type
type Redis struct {
	mock.Mock
}

// Get provides a mock function with given fields: key
func (_m *Redis) Get(key string) (string, error) {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScrapedAvailableHospitals provides a mock function with given fields: url
func (_m *Redis) GetScrapedAvailableHospitals(url string) (string, error) {
	ret := _m.Called(url)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScrapedDetailHospital provides a mock function with given fields: url
func (_m *Redis) GetScrapedDetailHospital(url string) (string, error) {
	ret := _m.Called(url)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: key, value
func (_m *Redis) Set(key string, value string) error {
	ret := _m.Called(key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetEx provides a mock function with given fields: key, value, expire
func (_m *Redis) SetEx(key string, value string, expire time.Duration) error {
	ret := _m.Called(key, value, expire)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, time.Duration) error); ok {
		r0 = rf(key, value, expire)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetScrapedAvailableHospitals provides a mock function with given fields: url, hospitals
func (_m *Redis) SetScrapedAvailableHospitals(url string, hospitals []model.HospitalSummary) error {
	ret := _m.Called(url, hospitals)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []model.HospitalSummary) error); ok {
		r0 = rf(url, hospitals)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetScrapedDetailHospital provides a mock function with given fields: url, hospitals
func (_m *Redis) SetScrapedDetailHospital(url string, hospitals model.HospitalDetail) error {
	ret := _m.Called(url, hospitals)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, model.HospitalDetail) error); ok {
		r0 = rf(url, hospitals)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}