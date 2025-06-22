package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/logging"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"moneyx.golang.framework/injection"
	"moneyx.golang.framework/logger"
	"moneyx.golang.framework/logger/logrepo"

	moneyxproto "moneyx.golang.framework/proto"
)

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID    uint `gorm:"primaryKey;"`
	Code  string
	Price uint
}

func customerMaker() (injection.BaseModel, error) {
	return &Customer{
		ID:   1,
		Name: "Mahdi",
	}, nil
}

func doWork(heartbeat chan<- struct{}) {
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 500)
		heartbeat <- struct{}{}
	}
}

func main() {
	y, _ := uuid.NewUUID()
	fmt.Printf("Correlation: %v", y)
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
	}
	i := injection.NewInjection()
	// initialize gofr object
	app := gofr.New(logger.NewGoFrLogger(logging.DEBUG, logrepo.NewMoneyxLogRepo()))
	dsn := "Initial Catalog=tempdb;MultipleActiveResultSets=true;Data Source=DESKTOP-J3OFINQ;User ID=sa;Password=mahdi1380;MultipleActiveResultSets=true;TrustServerCertificate=True;"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.NewGormLogger(logger.NewMoneyxLog(), logger.GormConfig{}, app.Metrics()),
	})
	injection.NewInjection().AddSingleton(app.Metrics())
	app.AddGorm(db)
	app.Metrics()
	moneyxproto.RegisterWhatsappServiceServerWithGofr(app, moneyxproto.NewWhatsappServiceGoFrServer(), moneyxproto.NewWhatsappServiceGoFrValidation())

	i.AddTransient(&Customer{}, customerMaker)
	//ctx = context.WithValue(ctx, db.txnCtxKey, tx)
	//txn, ok := ctx.Value(db.txnCtxKey).(Transaction)
	//app.UseMiddlewareWithContainer()
	app.POST("/customer/{name}", func(ctx *gofr.Context) (any, error) {
		name := ctx.PathParam("name")

		// Inserting a customer row in database using SQL
		_, err := ctx.SQL.ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)

		return nil, err
	})

	app.GET("/customer", func(ctx *gofr.Context) (any, error) {
		var customers []Customer

		// Getting the customer from the database using SQL
		//err := ctx.Gorm.WithContext(ctx).AutoMigrate(&Product{})
		customer, _, err := i.GetInstance(ctx, &Customer{})
		if err != nil {
			return nil, err
		}
		okCustomer, ok := customer.(*Customer)
		if !ok {
			return nil, &injection.InjectionError{
				Message: "Customer instance not found",
			}
		}
		customers = append(customers, *okCustomer)
		customer, _, err = i.GetInstance(ctx, &Customer{})
		if err != nil {
			return nil, err
		}
		okCustomer, ok = customer.(*Customer)
		if !ok {
			return nil, &injection.InjectionError{
				Message: "Customer instance not found",
			}
		}
		customers = append(customers, *okCustomer)
		return customers, nil
	})

	app.Run()
}
