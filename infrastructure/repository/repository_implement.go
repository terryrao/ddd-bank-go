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
