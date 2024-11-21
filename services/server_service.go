package services

import (
	"context"
	"server/repositories"
)

type ServerServiceInterface interface {
	GetService(ctx context.Context) error
}

type ServerServiceImplementation struct {
	TransactorRepository repositories.TransactorRepositoryInterface
}

func NewServerServiceImplementation(
	TransactorRepository repositories.TransactorRepositoryInterface,
) *ServerServiceImplementation {
	return &ServerServiceImplementation{
		TransactorRepository: TransactorRepository,
	}
}

func (ir *ServerServiceImplementation) GetService(ctx context.Context) error {
	err := ir.TransactorRepository.Atomic(ctx, func(cForTx context.Context) error {
		return nil
	})
	return err
}
