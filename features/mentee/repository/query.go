package repository

import (
	"api-alta-dashboard/features/mentee"
	"api-alta-dashboard/utils/helper"
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type menteeRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.RepositoryInterface {
	return &menteeRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *menteeRepository) Create(input mentee.Core) (row int, err error) {
	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements user.Repository
func (repo *menteeRepository) GetAll() (data []mentee.Core, err error) {
	var mentees []Mentee
	var queryStatus, queryIdClass, queryEdType string

	if len(queryIdClass) == 0 {
		queryIdClass = "0"
	}
	intIdClass, errConv := strconv.Atoi(queryIdClass)
	if errConv != nil {
		return nil, errors.New("error conver class id to filter")
	}

	tx := repo.db.Where(&Mentee{Status: queryStatus, ClassID: uint(intIdClass), EducationType: queryEdType}).Preload("Class").Find(&mentees)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(mentees)
	return dataCore, nil
}

// GetAll with search by name implements user.Repository
func (repo *menteeRepository) GetAllWithSearch(queryName, queryStatus, queryIdClass, queryEdType string) (data []mentee.Core, err error) {
	var mentees []Mentee

	helper.LogDebug("\n iniiii isi queryName = ", queryName)
	helper.LogDebug("\n isi queryIDClass= ", queryIdClass)
	helper.LogDebug("\n isi queryStatus = ", queryStatus)
	helper.LogDebug("\n isi queryEdType = ", queryEdType)

	if len(queryIdClass) == 0 {
		queryIdClass = "0"
	}
	intIdClass, errConv := strconv.Atoi(queryIdClass)
	if errConv != nil {
		return nil, errors.New("error conver class id to filter")
	}

	yx := repo.db.Where("name LIKE ?", "%"+queryName+"%")
	if yx.Error != nil {
		return nil, yx.Error
	}

	tx := repo.db.Where(&Mentee{Status: queryStatus, ClassID: uint(intIdClass), EducationType: queryEdType}).Find(&mentees)
	// tx = repo.db.Where("name LIKE ?", "%"+queryName+"%").Find(&mentees)
	// tx := repo.db.Where("name LIKE ?", "%"+queryName+"%")
	// tx = repo.db.Where("status = ?", queryStatus)
	// tx = repo.db.Where("id_class = ?", queryIdClass)
	// tx = repo.db.Where("education_type = ?", queryEdType)
	// tx = repo.db.Find(&mentees)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(mentees)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (repo *menteeRepository) GetById(id int) (data mentee.Core, err error) {
	var mentee Mentee

	tx := repo.db.Preload("Logs.User").First(&mentee, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, errors.New("data not found")
	}
	fmt.Println("\n\n\n data mentee", mentee)

	var dataCore = mentee.toCore()
	return dataCore, nil
}

// Update implements user.Repository
func (repo *menteeRepository) Update(input mentee.Core, id int) (row int, err error) {
	userGorm := fromCore(input)
	var user Mentee
	tx := repo.db.Model(&user).Where("ID = ?", id).Updates(&userGorm) // proses update
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("update failed")
	}
	return int(tx.RowsAffected), nil
}

// Delete implements user.Repository
func (repo *menteeRepository) Delete(id int) (row int, err error) {
	var user Mentee
	tx := repo.db.Delete(&user, id) // proses delete
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("delete failed")
	}
	return int(tx.RowsAffected), nil
}

func (repo *menteeRepository) FindUser(email string) (result mentee.Core, rowAffected int, err error) {
	var userData Mentee
	tx := repo.db.Where("email = ?", email).First(&userData)
	rowAffected = int(tx.RowsAffected)
	if tx.Error != nil {
		return mentee.Core{}, rowAffected, tx.Error
	}

	result = userData.toCore()

	return result, rowAffected, nil
}
