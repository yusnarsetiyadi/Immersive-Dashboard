package repository

import (
	"api-alta-dashboard/features/class"
	"errors"

	"gorm.io/gorm"
)

type classRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) class.RepositoryInterface {
	return &classRepository{
		db: db,
	}
}

// Create implements class.Repository
func (repo *classRepository) CreateClass(input class.Core) (row int, err error) {
	classGorm := fromCore(input)
	tx := repo.db.Create(&classGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("create failed")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements class.Repository
func (repo *classRepository) GetAllClass() (data []class.Core, err error) {
	var classes []Class

	tx := repo.db.Find(&classes)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(classes)
	return dataCore, nil
}

// GetAll with search by name implements class.Repository
func (repo *classRepository) GetAllWithSearchClass(query string) (data []class.Core, err error) {
	var classes []Class

	tx := repo.db.Where("full_name LIKE ?", "%"+query+"%").Find(&classes)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(classes)
	return dataCore, nil
}

// GetById implements class.RepositoryInterface
func (repo *classRepository) GetByIdClass(id int) (data class.Core, err error) {
	var class Class

	tx := repo.db.First(&class, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, errors.New("error")
	}

	var dataCore = class.toCore()
	return dataCore, nil
}

// Update implements class.Repository
func (repo *classRepository) UpdateClass(input class.Core, id int) (row int, err error) {
	classGorm := fromCore(input)
	var class Class
	tx := repo.db.Model(&class).Where("ID = ?", id).Updates(&classGorm) // proses update
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("update failed")
	}
	return int(tx.RowsAffected), nil
}

// Delete implements class.Repository
func (repo *classRepository) DeleteClass(id int) (row int, err error) {
	var class Class
	tx := repo.db.Delete(&class, id) // proses delete
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("delete failed")
	}
	return int(tx.RowsAffected), nil
}
