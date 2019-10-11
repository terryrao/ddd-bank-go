package service

import (
	"ddd-bank-go/domain"
	"fmt"
	"regexp"
	"time"
)

//BankService
type BankService interface {
	// <<Commands>>
	//CreateClient
	CreateClient(username string, birthday time.Time) (*domain.Client, error)
	DeleteClient(client *domain.Client)

	//<<queries>>
	FindClient(username string) *domain.Client
	FindAllClients() []*domain.Client
	FindYoungClients(fromBirth time.Time) []*domain.Client
	findRichClients(minBalance *domain.Amount) []*domain.Client
}

type BankServiceImpl struct {
	ClientRepository domain.ClientRepository
}

func (b *BankServiceImpl) CreateClient(username string, birthday time.Time) (*domain.Client, error) {
	pat := `[a-z_A-Z][a-z_A-Z0-9]{0,30}`
	compile, _ := regexp.Compile(pat)
	match := compile.Match([]byte(username))
	if !match {
		return nil, fmt.Errorf("illegal char in username")
	}
	client, e := b.ClientRepository.Save(&domain.Client{UserName: username, BirthDate: birthday})
	if e != nil {
		return nil, e
	}
	return client, nil

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

func (b *BankServiceImpl) FindYoungClients(fromBirth time.Time) []*domain.Client {
	panic("implement me")
}

func (b *BankServiceImpl) findRichClients(minBalance *domain.Amount) []*domain.Client {
	panic("implement me")
}
