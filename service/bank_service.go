package service

import (
	"ddd-bank-go/domain"
	"time"
)

//BankService
type BankService interface {
	// <<Commands>>
	CreateClient(id int64, username string, birthday time.Time) *domain.Client
	//DeleteClient(client *domain.Client)

	//<<queries>>
	//FindClient(username string) *domain.Client
	//FindAllClients() []*domain.Client
	//FindYoungClients(fromBirth time.Time) []*domain.Client
	//findRichClients(minBalance *domain.Amount) []*domain.Client
}

type BankServiceImpl struct {
	ClientRepository        domain.ClientRepository
	AccountRepository       domain.AccountRepository
	AccountAccessRepository domain.AccountAccessRepository
}

func (b *BankServiceImpl) CreateClient(id int64, username string, birthday time.Time) *domain.Client {
	return &domain.Client{
		Id:                      id,
		UserName:                username,
		BirthDate:               birthday,
		AccountRepository:       b.AccountRepository,
		AccountAccessRepository: b.AccountAccessRepository,
		ClientRepository:        b.ClientRepository,
	}
}
