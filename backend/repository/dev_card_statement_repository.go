package repository

import (
	"monelog/model"
	"gorm.io/gorm"
)

type IDevCardStatementRepository interface {
	DeleteAllCardStatements() (int64, error)
}

type devCardStatementRepository struct {
	db *gorm.DB
}

func NewDevCardStatementRepository(db *gorm.DB) IDevCardStatementRepository {
	return &devCardStatementRepository{db}
}

func (dcsr *devCardStatementRepository) DeleteAllCardStatements() (int64, error) {
	result := dcsr.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.CardStatement{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
