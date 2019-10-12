package di

import (
	"ddd-bank-go/controller"
	"ddd-bank-go/domain"
	"ddd-bank-go/infrastructure/repository"
	"ddd-bank-go/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go.uber.org/dig"
	"time"
)

var Container *dig.Container

var (
	dbhostip   = "localhost:3306"
	dbusername = "root"
	dbpassword = "123456"
	dbname     = "event"
)

func init() {
	Container = dig.New()
	checkErr(Container.Provide(func() (*xorm.Engine, error) {
		engine, err := xorm.NewEngine("mysql", dbusername+":"+dbpassword+"@tcp("+dbhostip+")/"+dbname+"?charset=utf8")
		checkErr(err)
		engine.ShowSQL(true)

		return engine, nil
	}))

	checkErr(Container.Provide(func(engine *xorm.Engine) domain.ClientRepository {
		return &repository.ClientRepositoryImpl{Engine: engine}
	}))

	checkErr(Container.Provide(func(engine *xorm.Engine) domain.AccountRepository {
		return &repository.AccountRepositoryImpl{Engine: engine}
	}))

	checkErr(Container.Provide(func(engine *xorm.Engine) domain.AccountAccessRepository {
		return &repository.AccountAccessRepositoryImpl{Engine: engine}
	}))
	checkErr(Container.Provide(func(c ClientParamsHandler) service.BankService {

		return &service.BankServiceImpl{
			AccountAccessRepository: c.AccountAccessRepository,
			AccountRepository:       c.AccountRepository,
			ClientRepository:        c.ClientRepository,
		}
	}))
	//Container.Invoke(func(clientRepository domain_repository.ClientRepository,
	//	accessRepository domain_repository.AccountAccessRepository,
	//	accountRepository domain_repository.AccountRepository) *domain.Client {
	//})
	checkErr(Container.Provide(func(bankService service.BankService) *controller.BankController {
		return &controller.BankController{BankService: bankService}
	}))
	checkErr(Container.Provide(func(p ClientParamsHandler) *domain.Client {
		return &domain.Client{
			Id:                      0,
			UserName:                "",
			BirthDate:               time.Time{},
			AccountRepository:       p.AccountRepository,
			AccountAccessRepository: p.AccountAccessRepository,
			ClientRepository:        p.ClientRepository,
		}
	}))

	fmt.Println(Container.String())

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type ClientParamsHandler struct {
	dig.In
	AccountRepository       domain.AccountRepository
	AccountAccessRepository domain.AccountAccessRepository
	ClientRepository        domain.ClientRepository
}
