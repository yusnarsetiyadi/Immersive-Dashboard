package repository

import (
	"api-alta-dashboard/features/log"
	"errors"

	"gorm.io/gorm"
)

type logRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) log.RepositoryInterface {
	return &logRepository{
		db: db,
	}
}

// Create implements class.Repository
func (repo *logRepository) CreateLog(input log.Core) error {
	logGorm := fromCore(input)
	tx := repo.db.Create(&logGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("create failed")
	}
	return nil
}

// GetAll implements class.Repository
func (repo *logRepository) GetAllLog() (data []log.Core, err error) {
	var logs []Log

	tx := repo.db.Preload("User").Find(&logs)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(logs)
	return dataCore, nil
}

// GetAll with search by name implements class.Repository
func (repo *logRepository) GetAllWithSearchLog(query string) (data []log.Core, err error) {
	var logs []Log
	tx := repo.db.Where("title LIKE ?", "%"+query+"%").Find(&logs)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(logs)
	return dataCore, nil
}

// Update implements class.Repository
func (repo *logRepository) UpdateLog(input log.Core, id int) error {
	logGorm := fromCore(input)
	var log Log
	tx := repo.db.Model(&log).Where("ID = ?", id).Updates(&logGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// Delete implements class.Repository
func (repo *logRepository) DeleteLog(id int) error {
	var log Log
	tx := repo.db.Delete(&log, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}
