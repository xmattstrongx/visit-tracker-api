package mocks

import mock "github.com/stretchr/testify/mock"
import models "RestApiProject/models"
import repositories "RestApiProject/repositories"

// CitiesRepository is an autogenerated mock type for the CitiesRepository type
type CitiesRepository struct {
	mock.Mock
}

// GetCitiesForSpecificState provides a mock function with given fields: _a0
func (_m *CitiesRepository) GetCitiesForSpecificState(_a0 int) ([]models.City, error) {
	ret := _m.Called(_a0)

	var r0 []models.City
	if rf, ok := ret.Get(0).(func(int) []models.City); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.City)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetailsForSpecificCity provides a mock function with given fields: _a0, _a1
func (_m *CitiesRepository) GetDetailsForSpecificCity(_a0 int, _a1 string) (models.CityDetails, error) {
	ret := _m.Called(_a0, _a1)

	var r0 models.CityDetails
	if rf, ok := ret.Get(0).(func(int, string) models.CityDetails); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(models.CityDetails)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ repositories.CitiesRepository = (*CitiesRepository)(nil)
