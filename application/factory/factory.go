package factory

import (
	"github.com//osmariojunior/codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"github.com/osmariojunior/codepix/application/usecase"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}

	return transactionUseCase
}
