package domain

import "time"

type AccountRepository interface {
	Find(no *AccountNo) *Account
	DeleteAll()
	Save(account *Account) *Account
}

type AccountAccessRepository interface {
	DeleteAll()
	Save(access *AccountAccess) *AccountAccess
	Delete(access *AccountAccess)
	FindManagedAccountsOf(client *Client, asOwner bool) []*Client
	FindFullAccounts(minBalance *Account) []*Account
	Find(client *Client, account *Account) *AccountAccess
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
