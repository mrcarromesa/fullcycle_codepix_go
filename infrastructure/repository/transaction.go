package repository

import (
	"fmt"

	"github.com/mrcarromesa/fullcycle_codepix_go/domain/model"
	"gorm.io/gorm"
)

type TransactionRepositoryDB struct {
	Db *gorm.DB
}

func (t *TransactionRepositoryDB) Register(transaction *model.Transaction) error {
	err := t.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDB) Save(transaction *model.Transaction) error {
	// update
	err := t.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryDB) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	t.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ? ", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}
