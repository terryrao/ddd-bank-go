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
