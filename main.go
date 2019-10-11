package main

import (
	"ddd-bank-go/di"
	"ddd-bank-go/service"
	"fmt"
	"time"
)

func main() {
	var bankService service.BankService
	e := di.Container.Invoke(func(service service.BankService) {
		bankService = service
	})
	checkErr(e)
	client, e := bankService.CreateClient("test", time.Now())
	checkErr(e)
	fmt.Printf("%v\n", client)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
