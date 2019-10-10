package domain_repository

import (
	"ddd-bank-go/domain"
	"time"
)

type ClientRepository interface {
	DeleteAll()
	Save(client *domain.Client) *domain.Client
	Delete(client *domain.Client)
	Find(id float64) *domain.Client
	FindByName(username string) *domain.Client
	FindAll() []*domain.Client
	FindAllBornFrom(birthday time.Time) []*domain.Client
}

type AccountRepository interface {
	Find(no *domain.AccountNo) *domain.Account
	DeleteAll()
	Save(account *domain.Account) *domain.Account
}

type AccountAccessRepository interface {
	DeleteAll()
	Save(access *domain.AccountAccess) *domain.AccountAccess
	Delete(access *domain.AccountAccess)
	FindManagedAccountsOf(client *domain.Client, asOwner bool) []*domain.Client
	FindFullAccounts(minBalance *domain.Account) []*domain.Account
	Find(client *domain.Client, account *domain.Account) *domain.AccountAccess
}
