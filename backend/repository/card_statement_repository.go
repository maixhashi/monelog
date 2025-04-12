package repository

import (
	"fmt"
	"monelog/model"

	"gorm.io/gorm"
)

type ICardStatementRepository interface {
	GetAllCardStatements(userId uint) ([]model.CardStatement, error)
	GetCardStatementById(userId uint, cardStatementId uint) (model.CardStatement, error)
	CreateCardStatement(cardStatement *model.CardStatement) error
	CreateCardStatements(cardStatements []model.CardStatement) error
	DeleteCardStatements(userId uint) error
}

type cardStatementRepository struct {
	db *gorm.DB
}

func NewCardStatementRepository(db *gorm.DB) ICardStatementRepository {
	return &cardStatementRepository{db}
}

func (csr *cardStatementRepository) GetAllCardStatements(userId uint) ([]model.CardStatement, error) {
	var cardStatements []model.CardStatement
	if err := csr.db.Where("user_id=?", userId).Order("statement_no, payment_count").Find(&cardStatements).Error; err != nil {
		return nil, err
	}
	return cardStatements, nil
}

func (csr *cardStatementRepository) GetCardStatementById(userId uint, cardStatementId uint) (model.CardStatement, error) {
	var cardStatement model.CardStatement
	result := csr.db.Where("user_id=?", userId).First(&cardStatement, cardStatementId)
	if result.Error != nil {
		return model.CardStatement{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.CardStatement{}, fmt.Errorf("card statement not found")
	}
	return cardStatement, nil
}

func (csr *cardStatementRepository) CreateCardStatement(cardStatement *model.CardStatement) error {
	return csr.db.Create(cardStatement).Error
}

func (csr *cardStatementRepository) CreateCardStatements(cardStatements []model.CardStatement) error {
	return csr.db.Create(&cardStatements).Error
}

func (csr *cardStatementRepository) DeleteCardStatements(userId uint) error {
	result := csr.db.Where("user_id=?", userId).Delete(&model.CardStatement{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
