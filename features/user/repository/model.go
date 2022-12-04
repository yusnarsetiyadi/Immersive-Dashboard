package repository

import (
	_user "api-alta-dashboard/features/user"

	"gorm.io/gorm"
)

//struct gorm model
type User struct {
	gorm.Model
	FullName string
	Email    string `validate:"required,email"`
	Password string `valudate:"required"`
	Team     string
	Role     string `valudate:"required"`
	Status   string
	// Books    []Book
}

// DTO
// mapping

//mengubah struct core ke struct model gorm
func fromCore(dataCore _user.Core) User {
	userGorm := User{
		FullName: dataCore.FullName,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Team:     dataCore.Team,
		Role:     dataCore.Role,
		Status:   dataCore.Status,
	}
	return userGorm
}

//mengubah struct model gorm ke struct core
func (dataModel *User) toCore() _user.Core {
	return _user.Core{
		ID:        dataModel.ID,
		FullName:  dataModel.FullName,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Team:      dataModel.Team,
		Role:      dataModel.Role,
		Status:    dataModel.Status,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

//mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []User) []_user.Core {
	var dataCore []_user.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
