package service

import (
	"api-alta-dashboard/features/log"
	"api-alta-dashboard/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLog(t *testing.T) {
	repo := new(mocks.LogRepository)
	t.Run("Success Create log", func(t *testing.T) {
		inputRepo := log.Core{Title: "alta", Feedback: "alta", Status: "alta", UserID: 1, MenteeID: 1}
		repo.On("CreateLog", inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.CreateLog(inputRepo)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create log", func(t *testing.T) {
		inputRepo := log.Core{Title: "alta", Feedback: "alta", Status: "alta", UserID: 1, MenteeID: 1}
		repo.On("CreateLog", inputRepo).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.CreateLog(inputRepo)
		assert.NotNil(t, err)
		assert.Equal(t, "failed", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestUpdateLog(t *testing.T) {
	repo := new(mocks.LogRepository)
	t.Run("Success update log", func(t *testing.T) {
		inputRepo := log.Core{Title: "alta", Feedback: "alta", Status: "alta", UserID: 1, MenteeID: 1}
		repo.On("UpdateLog", inputRepo, 1).Return(nil).Once()
		srv := New(repo)
		err := srv.UpdateLog(inputRepo, 1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update log", func(t *testing.T) {
		inputRepo := log.Core{Title: "alta", Feedback: "alta", Status: "alta", UserID: 1, MenteeID: 1}
		repo.On("UpdateLog", inputRepo, 1).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.UpdateLog(inputRepo, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "failed", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestDeleteLog(t *testing.T) {
	repo := new(mocks.LogRepository)
	t.Run("Success delete log", func(t *testing.T) {
		repo.On("DeleteLog", 1).Return(nil).Once()
		srv := New(repo)
		err := srv.DeleteLog(1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete log", func(t *testing.T) {
		repo.On("DeleteLog", 1).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.DeleteLog(1)
		assert.NotNil(t, err)
		assert.Equal(t, "failed", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestGetAllLog(t *testing.T) {
	repo := new(mocks.LogRepository)
	t.Run("Success get all log", func(t *testing.T) {
		inputRepo := []log.Core{{ID: 1, Title: "alta", Feedback: "alta", Status: "alta", UserID: 1, MenteeID: 1}}
		repo.On("GetAllLog").Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetAllLog("")
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].Title, response[0].Title)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all log", func(t *testing.T) {
		repo.On("GetAllLog").Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetAllLog("")
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})

	// t.Run("Success get all class", func(t *testing.T) {
	// 	inputRepo := []class.Core{{ID: 1, ClassName: "alta", UserID: 1}}
	// 	repo.On("GetAllClass").Return(inputRepo, nil).Once()
	// 	srv := New(repo)
	// 	response, err := srv.GetAllClass("query")
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, inputRepo[0].ClassName, response[0].ClassName)
	// 	repo.AssertExpectations(t)
	// })
}
