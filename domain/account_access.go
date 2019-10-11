package domain

//AccountAccess Many-to-Many association "Client may access Account" with specific privileges
type AccountAccess struct {
	Client  *Client
	Account *Account
	isOwner bool
}

func CreateAccountAccess(account *Account, client *Client, isOwner bool) *AccountAccess {
	return &AccountAccess{
		Client:  client,
		Account: account,
		isOwner: isOwner,
	}
}

type AccountAccessRepository interface {
	DeleteAll()
	Save(access *AccountAccess) *AccountAccess
	Delete(access *AccountAccess)
	FindManagedAccountsOf(client *Client, asOwner bool) []*Client
	FindFullAccounts(minBalance *Account) []*Account
	Find(client *Client, account *Account) *AccountAccess
}
