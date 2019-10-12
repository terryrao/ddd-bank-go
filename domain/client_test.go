package domain

import (
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateClient(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	cases := []struct {
		Name   string
		Result *Client
		err    error
	}{
		{
			Name:   "save success",
			Result: nil,
			err:    nil,
		}, {
			Name:   "save fail",
			Result: nil,
			err:    fmt.Errorf("throw error"),
		},
	}

	for _, c := range cases {
		Convey(c.Name, t, func() {
			clientRepository := NewMockClientRepository(controller)
			accountAccessRepository := NewMockAccountAccessRepository(controller)
			accountRepository := NewMockAccountRepository(controller)
			client := &Client{
				Id:                      1,
				UserName:                "mockTEst",
				BirthDate:               time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
				AccountRepository:       accountRepository,
				AccountAccessRepository: accountAccessRepository,
				ClientRepository:        clientRepository,
			}
			clientRepository.EXPECT().Save(client).Return(c.Result, c.err)
			e := client.CreateClient()
			So(e, ShouldEqual, c.err)
		})
	}
}

func TestClient_Deposit(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()
	clientRepository := NewMockClientRepository(mock)
	accountAccessRepository := NewMockAccountAccessRepository(mock)
	accountRepository := NewMockAccountRepository(mock)
	client := &Client{
		Id:                      1,
		UserName:                "mockTEst",
		BirthDate:               time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
		AccountRepository:       accountRepository,
		AccountAccessRepository: accountAccessRepository,
		ClientRepository:        clientRepository,
	}

	sourceAccount := AccountNo("1111")
	destAccount := AccountNo("2222")
	e := client.Transfer(&sourceAccount, &destAccount, &Amount{euros: 100})
	if e != nil {
		t.Errorf("test fail when transfer ")
	}

}
