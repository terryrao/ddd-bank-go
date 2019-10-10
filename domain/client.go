package domain

import (
	"ddd-bank-go/di"
	"ddd-bank-go/domain/base"
	"ddd-bank-go/domain/domain_repository"
	"fmt"
	"time"
)

type Client struct {
	base.Entity
	UserName                string
	BirthDate               time.Time
	accountRepository       domain_repository.AccountRepository
	accountAccessRepository domain_repository.AccountAccessRepository
}

//<<constructors>>

//CreateClient constructor
func CreateClient(username string, birthDay time.Time) *Client {
	var cRepository domain_repository.AccountRepository
	var aRepository domain_repository.AccountAccessRepository
	_ = di.Container.Invoke(func(repository domain_repository.AccountRepository) {
		cRepository = repository
	})
	_ = di.Container.Invoke(func(repository domain_repository.AccountAccessRepository) {
		aRepository = repository
	})
	return &Client{
		UserName:                username,
		BirthDate:               birthDay,
		accountRepository:       cRepository,
		accountAccessRepository: aRepository,
	}
}

func CreateClientByName(username string) *Client {
	var cRepository domain_repository.AccountRepository
	var aRepository domain_repository.AccountAccessRepository
	_ = di.Container.Invoke(func(repository domain_repository.AccountRepository) {
		cRepository = repository
	})
	_ = di.Container.Invoke(func(repository domain_repository.AccountAccessRepository) {
		aRepository = repository
	})
	return &Client{
		UserName:                username,
		accountRepository:       cRepository,
		accountAccessRepository: aRepository,
	}
}

//<<command>>

//CreateAccount
func (r *Client) CreateAccount(userName string) *AccountAccess {
	account := r.accountRepository.Save(NewAccount(userName))
	accountAccess := CreateAccountAccess(account, r, true)
	return r.accountAccessRepository.Save(accountAccess)
}

//Deposit
func (r *Client) Deposit(destination *AccountNo, amount *Amount) error {
	if amount.ToDouble() <= 0 {
		return fmt.Errorf("illegalArgument deposit %.2f", amount.ToDouble())
	}
	dest := r.findDestinationAccount(destination)
	dest.Balance = dest.Balance.Plus(amount)
	return nil
}

func (r *Client) findDestinationAccount(no *AccountNo) *Account {
	return r.accountRepository.Find(no)
}

//Transfer
func (r *Client) Transfer(source, destination AccountNo, amount Amount) {

}

//AddAccountManager
func (r *Client) AddAccountManager(account Account, manager Client) *AccountAccess {
	return nil
}

//<<Query>>

//FindMyAccount
func (r Client) FindMyAccount(no AccountNo) *Account {
	return nil
}

//AccountReport
func (r *Client) AccountReport() string {
	return ""
}
