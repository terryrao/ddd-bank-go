package domain

import (
	"fmt"
	"time"
)

type Client struct {
	// 如果此处 `xorm:"id"`，则插入数据的时候，会默认为0，插入成功后不会把新插入的id返回，如果想得到插入后的主键id，则id不需要写`xorm:"id"`
	Id                      int64 // `xorm:"id"`
	UserName                string
	BirthDate               time.Time
	AccountRepository       AccountRepository       `xorm:"-"` //忽略字段
	AccountAccessRepository AccountAccessRepository `xorm:"-"` // 忽略字段
}

//<<constructors>>

//CreateClient constructor
//func CreateClient(username string, birthDay time.Time) *Client {
//	var cRepository domain_repository.AccountRepository
//	var aRepository domain_repository.AccountAccessRepository
//	_ = di.Container.Invoke(func(repository domain_repository.AccountRepository) {
//		cRepository = repository
//	})
//	_ = di.Container.Invoke(func(repository domain_repository.AccountAccessRepository) {
//		aRepository = repository
//	})
//	return &Client{
//		UserName:                username,
//		BirthDate:               birthDay,
//		accountRepository:       cRepository,
//		accountAccessRepository: aRepository,
//	}
//}
//
//func CreateClientByName(username string) *Client {
//	var cRepository domain_repository.AccountRepository
//	var aRepository domain_repository.AccountAccessRepository
//	_ = di.Container.Invoke(func(repository domain_repository.AccountRepository) {
//		cRepository = repository
//	})
//	_ = di.Container.Invoke(func(repository domain_repository.AccountAccessRepository) {
//		aRepository = repository
//	})
//	return &Client{
//		UserName:                username,
//		accountRepository:       cRepository,
//		accountAccessRepository: aRepository,
//	}
//}

//<<command>>

//CreateAccount
func (r *Client) CreateAccount(userName string) *AccountAccess {
	account := r.AccountRepository.Save(NewAccount(userName))
	accountAccess := CreateAccountAccess(account, r, true)
	return r.AccountAccessRepository.Save(accountAccess)
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
	return r.AccountRepository.Find(no)
}

//Transfer
func (r *Client) Transfer(source, destination *AccountNo, amount *Amount) error {
	if amount.ToDouble() <= 0 {
		return fmt.Errorf("illegalArgument trnafer money %.2f", amount.ToDouble())
	}

	dest := r.findDestinationAccount(destination)
	sourceAccount := r.findDestinationAccount(source)
	if sourceAccount == nil {
		return fmt.Errorf("illegalArgument source Account not foud %d", source)
	}
	dest.Balance = dest.Balance.Plus(amount)
	sourceAccount.Balance = sourceAccount.Balance.Minus(amount)
	r.AccountRepository.Save(dest)
	r.AccountRepository.Save(sourceAccount)
	return nil

}

//AddAccountManager Adds the given manager Client to the given account in the role as
// manager, but not owner
func (r *Client) AddAccountManager(account *Account, manager *Client) (*AccountAccess, error) {
	accountAccess := r.AccountAccessRepository.Find(r, account)
	if accountAccess == nil {
		return nil, fmt.Errorf("can not find account=%v,client=%v", account, r)
	}
	if !accountAccess.isOwner {
		return nil, fmt.Errorf("not owner of the account[%s]", account)
	}

	managerAccess := r.AccountAccessRepository.Find(manager, account)
	if accountAccess != nil {
		return nil, fmt.Errorf("manager is duplicate, account=%v,manager=%v", account, manager)
	}

	managerAccess = CreateAccountAccess(account, manager, false)
	r.AccountAccessRepository.Save(managerAccess)
	return managerAccess, nil
}

//<<Query>>

//FindMyAccount
func (r *Client) FindMyAccount(no *AccountNo) (*Account, error) {
	account := r.AccountRepository.Find(no)
	if account == nil {
		return nil, fmt.Errorf("could not find Account account=%d", no)
	}
	return account, nil
}

//AccountReport
func (r *Client) AccountReport() string {
	return ""
}

type ClientRepository interface {
	DeleteAll()
	Save(client *Client) (*Client, error)
	Delete(client *Client)
	Find(id float64) *Client
	FindByName(username string) *Client
	FindAll() []*Client
	FindAllBornFrom(birthday time.Time) []*Client
}
