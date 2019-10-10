package service

import (
	"ddd-bank-go/domain"
	"time"
)

//BankService
type BankService interface {
	// <<Commands>>
	//CreateClient
	CreateClient(username string, birthday *time.Time) *domain.Client
	DeleteClient(client *domain.Client)

	//<<queries>>
	FindClient(username string) *domain.Client
	FindAllClients() []*domain.Client
	FindYoungClients(fromBirth *time.Time) []*domain.Client
	findRichClients(minBalance *domain.Amount) []*domain.Client
}

type BankServiceImpl struct {
}

func (b *BankServiceImpl) CreateClient(username string, birthday *time.Time) *domain.Client {
	panic("implement me")
}

func (b *BankServiceImpl) DeleteClient(client *domain.Client) {
	panic("implement me")
}

func (b *BankServiceImpl) FindClient(username string) *domain.Client {
	panic("implement me")
}

func (b *BankServiceImpl) FindAllClients() []*domain.Client {
	panic("implement me")
}

func (b *BankServiceImpl) FindYoungClients(fromBirth *time.Time) []*domain.Client {
	panic("implement me")
}

func (b *BankServiceImpl) findRichClients(minBalance *domain.Amount) []*domain.Client {
	panic("implement me")
}
