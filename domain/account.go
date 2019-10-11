package domain

import (
	"fmt"
)

//Account userAccount Domain entity
type Account struct {
	Id      int64
	Name    string
	Balance Amount
}

//String toString
func (a *Account) String() string {
	return fmt.Sprintf("Account{accountNo=%d,name=%s,amount=%.2f}", a.Id, a.Name, a.Balance.ToDouble())
}

//GetAccountNo 获取账户号
func (a *Account) GetAccountNo() AccountNo {
	return AccountNo(a.Id)
}

//GetMinimumBalance 获取最小值
func GetMinimumBalance() Amount {
	return Amount{-100}
}

//NewAccount 构造器
func NewAccount(name string) *Account {
	return &Account{Name: name}
}
