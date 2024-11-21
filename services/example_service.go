package services

import (
	"context"
	"server/repositories"
)

type ExampleServiceInterface interface {
	GetService(ctx context.Context) error
}

type ExampleServiceImplementation struct {
	TransactorRepository repositories.TransactorRepositoryInterface
}

func NewExampleServiceImplementation(
	TransactorRepository repositories.TransactorRepositoryInterface,
) *ExampleServiceImplementation {
	return &ExampleServiceImplementation{
		TransactorRepository: TransactorRepository,
	}
}

func (ir *ExampleServiceImplementation) GetService(ctx context.Context) error {
	err := ir.TransactorRepository.Atomic(ctx, func(cForTx context.Context) error {
		return nil
	})
	return err
}
