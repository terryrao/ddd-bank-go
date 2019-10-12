package main

import (
	"ddd-bank-go/api"
	"ddd-bank-go/controller"
	"ddd-bank-go/di"
	"github.com/go-xorm/xorm"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	start()
	defer stop()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var server *grpc.Server

func start() {

	initGRPC()

}

func stop() {
	_ = di.Container.Invoke(func(engine *xorm.Engine) {
		e := engine.Close()
		if e != nil {
			log.Fatalf("xorm close fail %v", e)
		}
	})
	server.Stop()
}

//func test() {
//	var bankService service.BankService
//	e := di.Container.Invoke(func(service service.BankService) {
//		bankService = service
//	})
//	checkErr(e)
//	client, e := bankService.CreateClient("test", time.Now())
//	checkErr(e)
//	fmt.Printf("%v\n", client)
//}
//
var (
	network = "tcp"
	address = ":9001"
)

func initGRPC() {
	listener, err := net.Listen(network, address)
	checkErr(err)
	server = grpc.NewServer()
	regGRPCService(server)
	if err = server.Serve(listener); err != nil {
		log.Fatalf("server listening failed %s", err)
	}

}

func regGRPCService(service *grpc.Server) {
	_ = di.Container.Invoke(func(controller *controller.BankController) {
		api.RegisterClientApiServer(service, controller)
	})
}
