package domain

import (
	"testing"
)

func TestAccount_String(t *testing.T) {
	t.Run("test String", func(t *testing.T) {
		account := Account{
			Name:    "test",
			Id:      12,
			Balance: Amount{100},
		}

		want := "Account{accountNo=12,name=test,amount=100.00}"
		got := account.String()
		if want != got {
			t.Errorf("got : %s, want : %s", want, got)
		}
	})
}
