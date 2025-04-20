package repository

import (
	"fmt"
	"monelog/model"

	"gorm.io/gorm"
)

type ICSVHistoryRepository interface {
	GetAllCSVHistories(userId uint) ([]model.CSVHistory, error)
	GetCSVHistoryById(userId uint, csvHistoryId uint) (model.CSVHistory, error)
	CreateCSVHistory(csvHistory *model.CSVHistory) error
	DeleteCSVHistory(userId uint, csvHistoryId uint) error
}

type csvHistoryRepository struct {
	db *gorm.DB
}

func NewCSVHistoryRepository(db *gorm.DB) ICSVHistoryRepository {
	return &csvHistoryRepository{db}
}

func (chr *csvHistoryRepository) GetAllCSVHistories(userId uint) ([]model.CSVHistory, error) {
	var csvHistories []model.CSVHistory
	if err := chr.db.Where("user_id=?", userId).Order("created_at DESC").Find(&csvHistories).Error; err != nil {
		return nil, err
	}
	return csvHistories, nil
}

func (chr *csvHistoryRepository) GetCSVHistoryById(userId uint, csvHistoryId uint) (model.CSVHistory, error) {
	var csvHistory model.CSVHistory
	result := chr.db.Where("user_id=?", userId).First(&csvHistory, csvHistoryId)
	if result.Error != nil {
		return model.CSVHistory{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.CSVHistory{}, fmt.Errorf("CSV history not found")
	}
	return csvHistory, nil
}

func (chr *csvHistoryRepository) CreateCSVHistory(csvHistory *model.CSVHistory) error {
	return chr.db.Create(csvHistory).Error
}

func (chr *csvHistoryRepository) DeleteCSVHistory(userId uint, csvHistoryId uint) error {
	result := chr.db.Where("user_id=? AND id=?", userId, csvHistoryId).Delete(&model.CSVHistory{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("CSV history not found or not authorized to delete")
	}
	return nil
}
