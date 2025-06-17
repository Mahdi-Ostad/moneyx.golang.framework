package main

import (
	"path/filepath"

	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/logging"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"moneyx.golang.framework/injection"
	"moneyx.golang.framework/logger"
	"moneyx.golang.framework/logger/logrepo"
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

func main() {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
	}
	i := injection.NewInjection()
	// initialize gofr object
	app := gofr.New(logger.NewGoFrLogger(logging.INFO, logrepo.NewMoneyxLogRepo()))
	dsn := "Initial Catalog=tempdb;MultipleActiveResultSets=true;Data Source=DESKTOP-J3OFINQ;User ID=sa;Password=mahdi1380;MultipleActiveResultSets=true;TrustServerCertificate=True;"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.NewGormLogger(logger.NewMoneyxLog(), logger.GormConfig{}, app.Metrics()),
	})
	app.AddGorm(db)

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
