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
		inputRepo := log.Core{Title: "alta", Feedback: "good", Status: "unit1", MenteeID: 1}
		repo.On("CreateLog", inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.CreateLog(inputRepo)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create log", func(t *testing.T) {
		inputRepo := log.Core{Title: "alta", Feedback: "good", Status: "unit1", MenteeID: 1}
		repo.On("CreateLog", inputRepo).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.CreateLog(inputRepo)
		assert.NotNil(t, err)
		assert.Equal(t, "failed", err.Error())
		repo.AssertExpectations(t)
	})
}
