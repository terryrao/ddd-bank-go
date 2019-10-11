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
	var birthday time.Time
	timestamp := req.BirthDate
	if timestamp > 0 {
		birthday = time.Unix(timestamp, 0)
	}
	_, e := b.BankService.CreateClient(req.GetUsername(), birthday)
	if e != nil {
		return &api.Result{Result: false}, e
	}
	return &api.Result{Result: true}, nil
}
