package service

import (
	"api-alta-dashboard/features/class"
	"api-alta-dashboard/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClass(t *testing.T) {
	repo := new(mocks.ClassRepository)
	t.Run("Success Create class", func(t *testing.T) {
		inputRepo := class.Core{Name: "alta", UserID: 1}
		repo.On("CreateClass", inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.CreateClass(inputRepo)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create class", func(t *testing.T) {
		inputRepo := class.Core{Name: "ALTA", UserID: 1}
		repo.On("CreateClass", inputRepo).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.CreateClass(inputRepo)
		assert.NotNil(t, err)
		assert.Equal(t, "failed", err.Error())
		repo.AssertExpectations(t)
	})
}
