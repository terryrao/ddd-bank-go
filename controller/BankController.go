package controller

import (
	"context"
	"ddd-bank-go/api"
	"ddd-bank-go/service"
	"time"
)

type BankController struct {
	BankService service.BankService
}

func (b *BankController) Save(ctx context.Context, req *api.Client) (*api.Result, error) {
	username := req.GetUsername()
	date := req.GetBirthDate()
	var birthday time.Time
	if date > 0 {
		birthday = time.Unix(date, 0)

	}
	client := b.BankService.CreateClient(0, username, birthday)
	e := client.CreateClient()
	if e != nil {
		return &api.Result{Result: false}, e
	}
	return &api.Result{Result: true}, nil
}
