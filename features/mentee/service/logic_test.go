package service

import (
	"api-alta-dashboard/features/mentee"
	"api-alta-dashboard/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.MenteeRepository)
	t.Run("Success get all", func(t *testing.T) {
		inputRepo := []mentee.Core{{ID: 1, Name: "alta", Email: "alta", Status: "alta", Gender: "alta", Nickname: "alta", Address: "alta", HomeAddress: "alta", Telegram: "alta", Discord: "alta", Phone: "alta", EmergencyName: "alta", EmergencyPhone: "alta", EmergencyStatus: "alta", EducationType: "alta", EducationMajor: "alta", EducationGraduate: "alta", ClassID: 1}}
		repo.On("GetAll").Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll("", "", "", "")
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetAll("", "", "", "")
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})

	// t.Run("Success get all with search ", func(t *testing.T) {
	// 	inputRepo := []mentee.Core{{ID: 1, Name: "alta", Email: "alta", Status: "alta", Gender: "alta", Nickname: "alta", Address: "alta", HomeAddress: "alta", Telegram: "alta", Discord: "alta", Phone: "alta", EmergencyName: "alta", EmergencyPhone: "alta", EmergencyStatus: "alta", EducationType: "alta", EducationMajor: "alta", EducationGraduate: "alta", ClassID: 1}}
	// 	repo.On("GetAllWithSearch", "").Return(inputRepo, nil).Once()
	// 	srv := New(repo)
	// 	response, err := srv.GetAll("query", "alta", "alta", "alta")
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, inputRepo[0].Name, response[0].Name)
	// 	repo.AssertExpectations(t)
	// })
}

func TestGetById(t *testing.T) {
	repo := new(mocks.MenteeRepository)
	t.Run("Success get by id", func(t *testing.T) {
		returnData := mentee.Core{ID: 1, Name: "alta", Email: "alta", Status: "alta", Gender: "alta", Nickname: "alta", Address: "alta", HomeAddress: "alta", Telegram: "alta", Discord: "alta", Phone: "alta", EmergencyName: "alta", EmergencyPhone: "alta", EmergencyStatus: "alta", EducationType: "alta", EducationMajor: "alta", EducationGraduate: "alta", ClassID: 1}
		repo.On("GetById", 1).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NoError(t, err)
		assert.Equal(t, returnData.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get by id ", func(t *testing.T) {
		inputRepo := mentee.Core{ID: 1, Name: "alta", Email: "alta", Status: "alta", Gender: "alta", Nickname: "alta", Address: "alta", HomeAddress: "alta", Telegram: "alta", Discord: "alta", Phone: "alta", EmergencyName: "alta", EmergencyPhone: "alta", EmergencyStatus: "alta", EducationType: "alta", EducationMajor: "alta", EducationGraduate: "alta", ClassID: 1}
		repo.On("GetById", 1).Return(mentee.Core{}, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NotNil(t, err)
		assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})
}
