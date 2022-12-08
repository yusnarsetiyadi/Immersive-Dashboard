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
		inputRepo := class.Core{ClassName: "alta", UserID: 1}
		repo.On("CreateClass", inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.CreateClass(inputRepo)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create class", func(t *testing.T) {
		inputRepo := class.Core{ClassName: "ALTA", UserID: 1}
		repo.On("CreateClass", inputRepo).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.CreateClass(inputRepo)
		assert.NotNil(t, err)
		// assert.Equal(t, "failed", err.Error())
		repo.AssertExpectations(t)
	})

	// t.Run("Failed Create class", func(t *testing.T) {
	// 	inputRepo := class.Core{ClassName: "ALTA", UserID: 1}
	// 	repo.On("CreateClass", inputRepo).Return(errors.New("failed")).Once()
	// 	srv := New(repo)
	// 	err := srv.CreateClass(inputRepo)
	// 	assert.NotNil(t, err)
	// 	// assert.Equal(t, "failed", err.Error())
	// 	repo.AssertExpectations(t)
	// })
}

func TestUpdateClass(t *testing.T) {
	repo := new(mocks.ClassRepository)
	t.Run("Success update class", func(t *testing.T) {
		inputRepo := class.Core{ClassName: "alta", UserID: 1}
		repo.On("UpdateClass", inputRepo, 1).Return(nil).Once()
		srv := New(repo)
		err := srv.UpdateClass(inputRepo, 1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update class", func(t *testing.T) {
		inputRepo := class.Core{ClassName: "ALTA", UserID: 1}
		repo.On("UpdateClass", inputRepo, 1).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.UpdateClass(inputRepo, 1)
		assert.NotNil(t, err)
		// assert.Equal(t, "failed", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestDeleteClass(t *testing.T) {
	repo := new(mocks.ClassRepository)
	t.Run("Success delete class", func(t *testing.T) {
		repo.On("DeleteClass", 1).Return(nil).Once()
		srv := New(repo)
		err := srv.DeleteClass(1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete class", func(t *testing.T) {
		repo.On("DeleteClass", 1).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.DeleteClass(1)
		assert.NotNil(t, err)
		// assert.Equal(t, "failed", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestGetByIdClass(t *testing.T) {
	repo := new(mocks.ClassRepository)
	t.Run("Success get by id class", func(t *testing.T) {
		inputRepo := class.Core{ID: 1, ClassName: "alta", UserID: 1}
		repo.On("GetByIdClass", 1).Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetByIdClass(1)
		assert.NoError(t, err)
		assert.Equal(t, inputRepo.ClassName, response.ClassName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get by id class", func(t *testing.T) {
		inputRepo := class.Core{ID: 1, ClassName: "alta", UserID: 1}
		repo.On("GetByIdClass", 1).Return(class.Core{}, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetByIdClass(1)
		assert.NotNil(t, err)
		assert.NotEqual(t, inputRepo.ClassName, response.ClassName)
		repo.AssertExpectations(t)
	})
}

func TestGetAllClass(t *testing.T) {
	repo := new(mocks.ClassRepository)
	t.Run("Success get all class", func(t *testing.T) {
		inputRepo := []class.Core{{ID: 1, ClassName: "alta", UserID: 1}}
		repo.On("GetAllClass").Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetAllClass("")
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].ClassName, response[0].ClassName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all class", func(t *testing.T) {
		repo.On("GetAllClass").Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetAllClass("")
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})

	t.Run("Success get all with search class", func(t *testing.T) {
		inputRepo := []class.Core{{ID: 1, ClassName: "alta", UserID: 1}}
		repo.On("GetAllWithSearchClass", "query").Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetAllClass("query")
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].ClassName, response[0].ClassName)
		repo.AssertExpectations(t)
	})
}
