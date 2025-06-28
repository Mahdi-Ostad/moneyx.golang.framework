package moneyx

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/joho/godotenv"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/datasource/mongo"
	"gofr.dev/pkg/gofr/logging"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
	"gorm.io/plugin/prometheus"
	"moneyx.golang.framework/config"
	"moneyx.golang.framework/logger"
	"moneyx.golang.framework/logger/logrepo"
	MoneyxMetrics "moneyx.golang.framework/moneyxmetrics"
	"moneyx.golang.framework/rabbitmq"

	gormopentracing "gorm.io/plugin/opentracing"
)

// type Customer struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// type Product struct {
// 	ID    uint `gorm:"primaryKey;"`
// 	Code  string
// 	Price uint
// }

// func customerMaker() (injection.BaseModel, error) {
// 	return &Customer{
// 		ID:   1,
// 		Name: "Mahdi",
// 	}, nil
// }

// func doWork(heartbeat chan<- struct{}) {
// 	time.Sleep(time.Second)
// 	for i := 0; i < 10; i++ {
// 		time.Sleep(time.Millisecond * 500)
// 		heartbeat <- struct{}{}
// 	}
// }

type MoneyxApp struct {
	app                *gofr.App
	backgroundServices []func(context.Context, *logger.MoneyxLog, func(context.Context) *gofr.Context)
	subscribers        []*IntegratedMessageSubscriber
}

func (m *MoneyxApp) createGoFrContext(ctx context.Context) *gofr.Context {
	return &gofr.Context{
		Container:     m.app.GetContainer(),
		Context:       ctx,
		ContextLogger: *logging.NewContextLogger(ctx, m.app.GetContainer().Logger),
	}
}

func (m *MoneyxApp) AddBackgroundService(fn func(context.Context, *logger.MoneyxLog, func(context.Context) *gofr.Context)) {
	m.backgroundServices = append(m.backgroundServices, fn)
}

func (m *MoneyxApp) runBackgroundService(fn func(context.Context, *logger.MoneyxLog, func(context.Context) *gofr.Context)) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	fn(ctx, logger.NewMoneyxLog(), m.createGoFrContext)
}

func (m *MoneyxApp) BuildAriyana(srv, val any, rpcFunc func(*gofr.App, any, any)) {
	moneyxLogger := logger.NewMoneyxLog()
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	// initialize gofr object
	m.app = gofr.New(logger.NewGoFrLogger(logging.DEBUG, logrepo.NewMoneyxLogRepo()))
	microserviceName := config.GetConfig("APP_NAME")
	dialect := config.GetConfig("DB_DIALECT")
	dsn := config.GetConfig("DB_URL")
	var db *gorm.DB
	if dialect == "mssql" {
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
			Logger: logger.NewGormLogger(logger.NewMoneyxLog(), logger.GormConfig{}, m.app.Metrics()),
		})
	} else if dialect == "postgres" {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.NewGormLogger(logger.NewMoneyxLog(), logger.GormConfig{}, m.app.Metrics()),
		})
	}
	if err != nil {
		log.Fatalf(err.Error())
	}
	//injection.NewInjection().AddSingleton(app.Metrics())
	if db != nil {
		if err := db.Use(tracing.NewPlugin()); err != nil {
			moneyxLogger.Errorf("failed to add tracing plugin to gorm: %v", err)
		}
		if dialect == "postgres" {
			db.Use(prometheus.New(prometheus.Config{
				DBName:          db.Config.Name(),
				RefreshInterval: 15,
				MetricsCollector: []prometheus.MetricsCollector{
					&prometheus.Postgres{VariableNames: []string{"Threads_running"}},
				},
			}))
		}
		db.Use(gormopentracing.New())
		m.app.AddGorm(db)
	}
	mongoUrl := config.GetConfig("MONGO_URL")
	if mongoUrl != "" {
		m.app.AddMongo(mongo.New(mongo.Config{
			URI: mongoUrl,
		}))
	}

	MoneyxMetrics.SetGlobalMetrics(m.app.Metrics())
	rabbitmqUrl := config.GetConfig("RABBIT_URL")
	if rabbitmqUrl != "" {
		m.app.AddPubSub(rabbitmq.NewRabbitMQPubSub(rabbitmqUrl, microserviceName))
	}
	for _, subscriber := range m.subscribers {
		m.app.Subscribe(subscriber.topic, subscriber.handler)
	}
	for _, fn := range m.backgroundServices {
		go m.runBackgroundService(fn)
	}
	rpcFunc(m.app, srv, val)

	//ctx = context.WithValue(ctx, db.txnCtxKey, tx)
	//txn, ok := ctx.Value(db.txnCtxKey).(Transaction)
	//app.UseMiddlewareWithContainer()
	// app.POST("/customer/{name}", func(ctx *gofr.Context) (any, error) {
	// 	name := ctx.PathParam("name")
	// 	outbox.Publish(ctx, &ActionLog.AccountActionLogIntegratedCommand{})

	// 	// Inserting a customer row in database using SQL
	// 	_, err := ctx.SQL.ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)

	// 	return nil, err
	// })

	// app.GET("/customer", func(ctx *gofr.Context) (any, error) {
	// 	var customers []Customer

	// 	// Getting the customer from the database using SQL
	// 	//err := ctx.Gorm.WithContext(ctx).AutoMigrate(&Product{})
	// 	customer, _, err := i.GetInstance(ctx, &Customer{})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	okCustomer, ok := customer.(*Customer)
	// 	if !ok {
	// 		return nil, &injection.InjectionError{
	// 			Message: "Customer instance not found",
	// 		}
	// 	}
	// 	customers = append(customers, *okCustomer)
	// 	customer, _, err = i.GetInstance(ctx, &Customer{})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	okCustomer, ok = customer.(*Customer)
	// 	if !ok {
	// 		return nil, &injection.InjectionError{
	// 			Message: "Customer instance not found",
	// 		}
	// 	}
	// 	customers = append(customers, *okCustomer)
	// 	return customers, nil
	// })

	m.app.Run()
}
