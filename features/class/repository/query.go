package repository

import (
	"api-alta-dashboard/features/class"
	"errors"
	"fmt"

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
func (repo *classRepository) CreateClass(input class.Core) error {
	classGorm := fromCore(input)
	tx := repo.db.Create(&classGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("create failed")
	}
	return nil
}

// GetAll implements class.Repository
func (repo *classRepository) GetAllClass() (data []class.Core, err error) {
	var classes []Class

	tx := repo.db.Preload("User").Find(&classes)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, errors.New("Data not found.")
	}

	fmt.Println("\n\n 1 getall class = ", classes)

	var dataCore = toCoreList(classes)
	return dataCore, nil
}

// GetAll with search by name implements class.Repository
func (repo *classRepository) GetAllWithSearchClass(query string) (data []class.Core, err error) {
	var classes []Class

	tx := repo.db.Preload("User").Where("class_name LIKE ?", "%"+query+"%").Find(&classes)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, errors.New("Data not found.")
	}

	fmt.Println("\n\n 2 getall class = ", classes)

	var dataCore = toCoreList(classes)
	return dataCore, nil
}

// GetById implements class.RepositoryInterface
func (repo *classRepository) GetByIdClass(id int) (data class.Core, err error) {
	var class Class

	tx := repo.db.Preload("User").First(&class, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, errors.New("Data not found.")
	}

	fmt.Println("\n\n 3 getall class = ", class)

	var dataCore = class.toCore()
	return dataCore, nil
}

// Update implements class.Repository
func (repo *classRepository) UpdateClass(input class.Core, id int) error {
	classGorm := fromCore(input)
	var class Class
	tx := repo.db.Model(&class).Where("ID = ?", id).Updates(&classGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// Delete implements class.Repository
func (repo *classRepository) DeleteClass(id int) error {
	var class Class
	tx := repo.db.Delete(&class, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}
