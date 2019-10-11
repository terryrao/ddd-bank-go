package di

import (
	"ddd-bank-go/domain"
	"ddd-bank-go/infrastructure/repository"
	"ddd-bank-go/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go.uber.org/dig"
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
	err := Container.Provide(func() (*xorm.Engine, error) {
		engine, err := xorm.NewEngine("mysql", dbusername+":"+dbpassword+"@tcp("+dbhostip+")/"+dbname+"?charset=utf8")
		checkErr(err)
		engine.ShowSQL(true)

		return engine, nil
	})
	checkErr(err)

	err = Container.Provide(func(engine *xorm.Engine) domain.ClientRepository {
		return &repository.ClientRepositoryImpl{Engine: engine}
	})
	checkErr(err)

	err = Container.Provide(func(clientRepository domain.ClientRepository) service.BankService {
		return &service.BankServiceImpl{ClientRepository: clientRepository}
	})
	checkErr(err)

	//Container.Invoke(func(clientRepository domain_repository.ClientRepository,
	//	accessRepository domain_repository.AccountAccessRepository,
	//	accountRepository domain_repository.AccountRepository) *domain.Client {
	//})

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
