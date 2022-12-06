package repository

import (
	_class "api-alta-dashboard/features/class"

	"gorm.io/gorm"
)

// struct gorm model
type Class struct {
	gorm.Model
	Name   string
	UserID uint
	// User   user.User
}

// DTO
// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore _class.Core) Class {
	classGorm := Class{
		Name:   dataCore.Name,
		UserID: dataCore.UserID,
	}
	return classGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Class) toCore() _class.Core {
	return _class.Core{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		UserID:    dataModel.UserID,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Class) []_class.Core {
	var dataCore []_class.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
