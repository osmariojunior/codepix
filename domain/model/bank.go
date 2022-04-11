package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base `valid:"required"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
	}
	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreateAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
