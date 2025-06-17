package logrepo

import (
	"database/sql"
	"fmt"
	"log"
	"runtime/debug"
	"sync"
	"time"

	//mssql "github.com/denisenkom/go-mssqldb"
	"gorm.io/driver/sqlserver"
	"moneyx.golang.framework/config"
)

var tableName string

type MoneyxLogLevel int

const DELAY_READ time.Duration = 10
const DELAY_REMOVE time.Duration = 1

const (
	// DEBUG debug log level
	DEBUG MoneyxLogLevel = iota
	// INFO info log level
	INFO
	// Warn warn log level
	WARN
	// ERROR error log level
	ERROR
	// FATAL fatal log level
	FATAL
)

func (lvl MoneyxLogLevel) String() string {
	names := []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	if int(lvl) >= 0 && int(lvl) < len(names) {
		return names[lvl]
	}
	return "UNKNOWN"
}

func ParseMoneyxLogLevel(level string) MoneyxLogLevel {
	switch level {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return INFO // Default to INFO if unknown level
	}
}

type LogMessage struct {
	level      string
	message    string
	timestamp  string
	stacktrace string
}

type MoneyxLogRepo struct {
	db       *sql.DB
	Messages chan LogMessage
}

var NewMoneyxLogRepo = sync.OnceValue(getNewMoneyxLogRepo)

func getNewMoneyxLogRepo() *MoneyxLogRepo {
	tableName = config.GetConfig("LOGGER_TABLE_NAME")
	if tableName == "" {
		fmt.Printf("LoggerTableName is not set in config, using default 'logs'\n")
		return nil
	}
	loggerConnection := config.GetConfig("LOGGER_CONNECTION")
	if loggerConnection == "" {
		fmt.Printf("LoggerConnection is not set in config, using default 'logs'\n")
		return nil
	}
	db, err := sql.Open("sqlserver", loggerConnection)
	if err != nil {
		fmt.Printf("Error opening database connection: %s\n", err.Error())
		return nil
	}
	maxConnection := config.GetIntConfigOrDefault("MAX_CONNECTIONS", 1)
	db.SetMaxOpenConns(maxConnection)
	db.SetMaxIdleConns(maxConnection)
	//_, err = db.Exec("CREATE TABLE IF NOT EXISTS [enqueued_messages] (id INTEGER PRIMARY KEY AUTOINCREMENT, manager_id TEXT NOT NULL, message TEXT NOT NULL)")

	_, err = db.Exec(fmt.Sprintf(`
		IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = '%s')
		BEGIN
			CREATE TABLE [dbo].[%s] ( 
    			[id] bigint IDENTITY(1, 1) NOT NULL, 
    			[level] varchar(50) NOT NULL, 
    			[message] nvarchar(max) NOT NULL, 
				[timestamp] datetime NOT NULL,
				[stacktrace] varchar(500) NOT NULL,
    			CONSTRAINT [PK_%s] PRIMARY KEY ([id])
			)
		END
	`, tableName, tableName, tableName))
	if err != nil {
		fmt.Printf("Error creating table %s: %s\n", tableName, err.Error())
		return nil
	}
	logRepo := &MoneyxLogRepo{db: db, Messages: make(chan LogMessage, 1000)}
	go logRepo.log()
	go logRepo.cleanLog()
	return logRepo
}

func (logRepo *MoneyxLogRepo) log() {
	defer func() {
		if value := recover(); value != nil {
			err, ok := value.(error)
			if ok {
				fmt.Println("Error", err.Error()+":"+string(debug.Stack()))
			} else {
				fmt.Println("SayHello Recovered:", err)
			}
		}
	}()
	for {
		time.Sleep(time.Second * DELAY_READ)
		messages := []LogMessage{}
	L:
		for {
			select {
			case msg, ok := <-logRepo.Messages:
				if !ok {
					// Channel closed, exit loop
					break L
				}
				messages = append(messages, msg)
			default:
				// No more messages, exit loop
				break L
			}
		}
		isFatal, err := logRepo.bulkInsertDb(messages)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		if isFatal {
			log.Fatalf("")
		}
	}
}

func (logRepo *MoneyxLogRepo) cleanLog() {
	defer func() {
		if value := recover(); value != nil {
			err, ok := value.(error)
			if ok {
				fmt.Println("Error", err.Error()+":"+string(debug.Stack()))
			} else {
				fmt.Println("SayHello Recovered:", err)
			}
		}
	}()
	for {
		time.Sleep(time.Hour * DELAY_REMOVE)
		err := logRepo.cleanDb()
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}
}

func (r *MoneyxLogRepo) bulkInsertDb(messages []LogMessage) (isFatal bool, err error) {
	if r == nil {
		return false, nil
	}
	if r.db == nil {
		return false, nil
	}
	hasFatal := false
	for _, message := range messages {
		if message.level == "Fatal Error" {
			hasFatal = true
			break
		}
	}
	//bulkimportStr := mssql.CopyIn(tableName, mssql.BulkOptions{}, "level", "message", "timestamp", "stacktrace")
	bulkimportStr := ""
	stmt, err := r.db.Prepare(bulkimportStr)
	if err != nil {
		return isFatal, err
	}
	for _, message := range messages {
		tm, _ := time.Parse("2006-01-02 15:04:05", message.timestamp)
		stmt.Exec(message.level, message.message, tm, message.stacktrace)
	}
	_, err = stmt.Exec()
	return hasFatal, err
}

func (r *MoneyxLogRepo) cleanDb() error {
	if r == nil {
		return nil
	}
	if r.db == nil {
		return nil
	}
	normalLogs := time.Now().Add(-3 * time.Hour).Format("2006-01-02 15:04:05")
	errorLogs := time.Now().Add(-24 * time.Hour * 2).Format("2006-01-02 15:04:05")
	_, err := r.db.Exec(fmt.Sprintf("DELETE FROM [%s] WHERE stacktrace='' AND timestamp<@p1", tableName), normalLogs)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(fmt.Sprintf("DELETE FROM [%s] WHERE stacktrace!='' AND timestamp<@p1", tableName), errorLogs)
	if err != nil {
		return err
	}
	return nil
}

func (r *MoneyxLogRepo) InsertMessage(level MoneyxLogLevel, message string, timestamp string, stacktrace string) {
	if r == nil {
		return
	}
	if r.db == nil {
		return
	}
	select {
	case r.Messages <- LogMessage{level: level.String(), message: message, timestamp: timestamp, stacktrace: stacktrace}:
		break
	default:
		break
	}

}

func (r *MoneyxLogRepo) Close() {
	if r == nil {
		return
	}
	if r.db == nil {
		return
	}
	r.db.Close()
}
