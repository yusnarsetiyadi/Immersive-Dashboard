package service

import (
	"api-alta-dashboard/features/class"
	"api-alta-dashboard/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllClass(t *testing.T) {
	repo := new(mocks.ClassRepository)
	returnData := []class.Core{{ID: 1, Name: "be13", UserID: 1}}
	//test for success
	t.Run("success get all class", func(t *testing.T) {
		repo.On("GetAllClass").Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAllClass()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})
	//test for failed
	t.Run("failed get all class", func(t *testing.T) {
		repo.On("GetAllClass").Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetAllClass()
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestCreateClass(t *testing.T) {
	repo := new(mocks.ClassRepository)
	//test for success
	t.Run("success create data", func(t *testing.T) {
		createData := class.Core{ID: 1, Name: "be13", UserID: 1}
		repo.On("CreateClass", createData).Return(1, nil).Once()
		srv := New(repo)
		err := srv.CreateClass(createData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	// test for failed
	t.Run("failed create data", func(t *testing.T) {
		createData := class.Core{ID: 1, Name: "be13", UserID: 1}
		repo.On("CreateClass", createData).Return(0, errors.New("failed")).Once()
		srv := New(repo)
		err := srv.CreateClass(createData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed create data, name empty", func(t *testing.T) {
		createData := class.Core{ID: 1, UserID: 1}
		srv := New(repo)
		err := srv.CreateClass(createData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetByIdClass(t *testing.T) {
	repo := new(mocks.ClassRepository)
	returnData := class.Core{ID: 1, Name: "be13", UserID: 1}
	//test for success
	t.Run("success get user by id class", func(t *testing.T) {
		repo.On("GetByIdClass", 1).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetByIdClass(1)
		assert.Nil(t, err)
		assert.Equal(t, returnData.Name, response.Name)
		repo.AssertExpectations(t)
	})
	// test for failed
	// 	t.Run("failed get user by id class", func(t *testing.T) {
	// 		repo.On("GetByIdClass", 1).Return(nil, errors.New("failed")).Once()
	// 		srv := New(repo)
	// 		response, err := srv.GetByIdClass(1)
	// 		assert.NotNil(t, err)
	// 		assert.NotEqual(t, returnData.Name, response.Name)
	// 		repo.AssertExpectations(t)
	// 	})
}

// func TestDeleteClass(t *testing.T) {
// 	repo := new(mocks.ClassRepository)
// 	//test for success
// 	t.Run("success delete data", func(t *testing.T) {
// 		repo.On("DeleteClass", 1).Return(1, nil).Once()
// 		srv := New(repo)
// 		err := srv.DeleteClass(1)
// 		assert.Nil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// 	//test for failed
// 	t.Run("failed delete data", func(t *testing.T) {
// 		repo.On("DeleteClass", 1).Return(errors.New("failed")).Once()
// 		srv := New(repo)
// 		err := srv.DeleteClass(1)
// 		assert.NotNil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestUpdateClass(t *testing.T) {
// 	repo := new(mocks.ClassRepository)
// 	//test for success
// 	t.Run("success update class", func(t *testing.T) {
// 		insertData := class.Core{ID: 1, Name: "be13", UserID: 1}
// 		repo.On("UpdateClass", insertData, 1).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.UpdateClass(insertData, 1)
// 		assert.NotNil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// 	// 	//test for failed
// 	t.Run("failed update data", func(t *testing.T) {
// 		insertData := class.Core{ID: 1, Name: "be13", UserID: 1}
// 		repo.On("UpdateClass", 1, insertData).Return(errors.New("failed")).Once()
// 		srv := New(repo)
// 		err := srv.UpdateClass(insertData, 1)
// 		assert.NotNil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// }
