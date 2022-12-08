package service

import (
	"api-alta-dashboard/features/user"
	"api-alta-dashboard/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestCreate(t *testing.T) {
// 	repo := new(mocks.UserRepository)
// 	t.Run("Success Create user", func(t *testing.T) {
// 		inputRepo := user.Core{FullName: "alta", Email: "alta", Password: "alta", Team: "alta", Role: "User", Status: "alta"}
// 		repo.On("Create", inputRepo).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.Create(inputRepo)
// 		assert.NoError(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Failed Create class", func(t *testing.T) {
// 		inputRepo := class.Core{ClassName: "ALTA", UserID: 1}
// 		repo.On("CreateClass", inputRepo).Return(errors.New("failed")).Once()
// 		srv := New(repo)
// 		err := srv.CreateClass(inputRepo)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "failed", err.Error())
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestUpdateClass(t *testing.T) {
// 	repo := new(mocks.ClassRepository)
// 	t.Run("Success update class", func(t *testing.T) {
// 		inputRepo := class.Core{ClassName: "alta", UserID: 1}
// 		repo.On("UpdateClass", inputRepo, 1).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.UpdateClass(inputRepo, 1)
// 		assert.NoError(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Failed Update class", func(t *testing.T) {
// 		inputRepo := class.Core{ClassName: "ALTA", UserID: 1}
// 		repo.On("UpdateClass", inputRepo, 1).Return(errors.New("failed")).Once()
// 		srv := New(repo)
// 		err := srv.UpdateClass(inputRepo, 1)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "failed", err.Error())
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	repo := new(mocks.UserRepository)
// 	t.Run("Success delete user", func(t *testing.T) {
// 		repo.On("Delete", 1).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.Delete(1)
// 		assert.Nil(t, err)
// 		repo.AssertExpectations(t)
// 	})

// t.Run("Failed delete class", func(t *testing.T) {
// 	repo.On("DeleteClass", 1).Return(errors.New("failed")).Once()
// 	srv := New(repo)
// 	err := srv.DeleteClass(1)
// 	assert.NotNil(t, err)
// 	assert.Equal(t, "failed", err.Error())
// 	repo.AssertExpectations(t)
// })
// }

func TestGetAll(t *testing.T) {
	repo := new(mocks.UserRepository)
	t.Run("Success get all user", func(t *testing.T) {
		inputRepo := []user.Core{{ID: 1, FullName: "alta", Email: "alta", Password: "alta", Team: "alta", Role: "alta", Status: "alta"}}
		repo.On("GetAll").Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll("")
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].FullName, response[0].FullName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetAll("")
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})

	t.Run("Success get all with search ", func(t *testing.T) {
		inputRepo := []user.Core{{ID: 1, FullName: "alta", Email: "alta", Password: "alta", Team: "alta", Role: "alta", Status: "alta"}}
		repo.On("GetAllWithSearch", "query").Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll("query")
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].FullName, response[0].FullName)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.UserRepository)
	t.Run("Success get by id user", func(t *testing.T) {
		returnData := user.Core{ID: 1, FullName: "alta", Email: "alta", Password: "alta", Team: "alta", Role: "alta", Status: "alta"}
		repo.On("GetById", 1).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NoError(t, err)
		assert.Equal(t, returnData.FullName, response.FullName)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get by id user", func(t *testing.T) {
		inputRepo := user.Core{ID: 1, FullName: "alta", Email: "alta", Password: "alta", Team: "alta", Role: "alta", Status: "alta"}
		repo.On("GetById", 1).Return(user.Core{}, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NotNil(t, err)
		assert.NotEqual(t, inputRepo.FullName, response.FullName)
		repo.AssertExpectations(t)
	})
}
