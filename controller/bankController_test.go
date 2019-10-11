package controller

import (
	"context"
	"ddd-bank-go/api"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

const (
	url = "localhost:9001"
)

func TestBankController_Save(t *testing.T) {
	conn, e := grpc.Dial(url, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("can not connect to %s %v", url, conn)
	}

	defer func() {
		e := conn.Close()
		if e != nil {
			log.Fatalf("close connect fail %v", conn)

		}
	}()
	c := api.NewClientApiClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Save(ctx, &api.Client{Username: "grpc", BirthDate: time.Now().Unix()})
	if err != nil {
		t.Errorf("could not call Save: %v", err)
	}
	t.Logf("result: %t", r.Result)

}
