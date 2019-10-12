package repository

import (
	"ddd-bank-go/domain"
	"github.com/go-xorm/xorm"
	"time"
)

//type Err string
//
//func (e Err) Error() string {
//	return string(e)
//}

type ClientRepositoryImpl struct {
	Engine *xorm.Engine
}

func (c *ClientRepositoryImpl) DeleteAll() {
	panic("implement me")
}

func (c *ClientRepositoryImpl) Save(client *domain.Client) (*domain.Client, error) {
	result, e := c.Engine.Insert(client)
	if e != nil || result <= 0 {
		return nil, e
	}
	return client, nil
}

func (c *ClientRepositoryImpl) Delete(client *domain.Client) {
	panic("implement me")
}

func (c *ClientRepositoryImpl) Find(id float64) *domain.Client {
	panic("implement me")
}

func (c *ClientRepositoryImpl) FindByName(username string) *domain.Client {
	panic("implement me")
}

func (c *ClientRepositoryImpl) FindAll() []*domain.Client {
	panic("implement me")
}

func (c *ClientRepositoryImpl) FindAllBornFrom(birthday time.Time) []*domain.Client {
	panic("implement me")
}

type AccountRepositoryImpl struct {
	Engine *xorm.Engine
}

func (a *AccountRepositoryImpl) Find(no *domain.AccountNo) *domain.Account {
	panic("implement me")
}

func (a *AccountRepositoryImpl) DeleteAll() {
	panic("implement me")
}

func (a *AccountRepositoryImpl) Save(account *domain.Account) *domain.Account {
	panic("implement me")
}

type AccountAccessRepositoryImpl struct {
	Engine *xorm.Engine
}

func (a *AccountAccessRepositoryImpl) DeleteAll() {
	panic("implement me")
}

func (a *AccountAccessRepositoryImpl) Save(access *domain.AccountAccess) *domain.AccountAccess {
	panic("implement me")
}

func (a *AccountAccessRepositoryImpl) Delete(access *domain.AccountAccess) {
	panic("implement me")
}

func (a *AccountAccessRepositoryImpl) FindManagedAccountsOf(client *domain.Client, asOwner bool) []*domain.Client {
	panic("implement me")
}

func (a *AccountAccessRepositoryImpl) FindFullAccounts(minBalance *domain.Account) []*domain.Account {
	panic("implement me")
}

func (a *AccountAccessRepositoryImpl) Find(client *domain.Client, account *domain.Account) *domain.AccountAccess {
	panic("implement me")
}
