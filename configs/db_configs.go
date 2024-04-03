package configs

import (
	"advertising/models"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


var DbConn *gorm.DB

type dbConfig struct {
	Host             string `json:"host"`
	Port             int    `json:"port"`
	Name             string `json:"name"`
	Username         string `json:"username"`
	Password         string `json:"pasword"`
	MaxIdle          int    `json:"maxidleconn"`
	MaxOpen          int    `json:"maxopenconn"`
	SSLMode          string `json:"sslmode"`
	ConnMaxLifeTime  int64  `json:"connmaxlifetime"`
	MigrationEnabled string `json:"migrationenabled"`
	MigrationMark    string `json:"migrationmark"`
}

func DBsetup() error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		Cfg.DB.Host,
		Cfg.DB.Username,
		Cfg.DB.Password,
		Cfg.DB.Name,
		Cfg.DB.Port,
		Cfg.DB.SSLMode,
	)

	DbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	DbConn.AutoMigrate(models.MigrationModel{})

	sqlDB, err := DbConn.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(Cfg.DB.MaxIdle)
	sqlDB.SetMaxOpenConns(Cfg.DB.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(Cfg.DB.ConnMaxLifeTime) * time.Hour)

	if err = sqlDB.Ping(); err != nil {
		return err
	}

	fmt.Println("Database connected ...")

	return nil
}

func getDBConfig() (config *dbConfig, err error) {

	var (
		host               string
		portStr            string
		port               int
		name               string
		username           string
		password           string
		maxIdleStr         string
		maxIdle            int
		maxOpenStr         string
		maxOpen            int
		sslMode            string
		connMaxLifeTime    int64
		connMaxLifeTimeStr string
	)

	host = os.Getenv("DB_HOST")
	if len(host) == 0 {
		return nil, errors.New("cannot get current DB host from env")
	}

	portStr = os.Getenv("DB_PORT")
	if len(portStr) == 0 {
		return nil, errors.New("cannot get current DB port from env")
	}

	port, err = strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.New("cannot get current DB port" + " : " + err.Error())
	}

	name = os.Getenv("DB_NAME")
	if len(name) == 0 {
		return nil, errors.New("cannot get current DB name from env")
	}

	username = os.Getenv("DB_USER")
	if len(username) == 0 {
		return nil, errors.New("cannot get current DB username from env")
	}

	password = os.Getenv("DB_PASSWORD")
	if len(password) == 0 {
		return nil, errors.New("cannot get current DB password from env")
	}

	maxIdleStr = os.Getenv("DB_MAXIDLE")
	if len(maxIdleStr) == 0 {
		return nil, errors.New("cannot get current DB max idle from env")
	}
	maxIdle, err = strconv.Atoi(maxIdleStr)
	if err != nil {
		return nil, errors.New("cannot get current DB max idle" + " : " + err.Error())
	}

	maxOpenStr = os.Getenv("DB_MAXOPEN")
	if len(maxOpenStr) == 0 {
		return nil, errors.New("cannot get current DB max open from env")
	}
	maxOpen, err = strconv.Atoi(maxOpenStr)
	if err != nil {
		return nil, errors.New("cannot get current DB max open" + " : " + err.Error())
	}

	sslMode = os.Getenv("DB_SSLMODE")
	if len(sslMode) == 0 {
		return nil, errors.New("cannot get current DB ssl mode from env")
	}

	connMaxLifeTimeStr = os.Getenv("DB_CONNMAXLIFETIME")
	if len(connMaxLifeTimeStr) == 0 {
		return nil, errors.New("cannot get current DB max life time from env")
	}
	connMaxLifeTime, err = strconv.ParseInt(connMaxLifeTimeStr, 10, 64)
	if err != nil {
		return nil, errors.New("cannot get current DB max life time" + " : " + err.Error())
	}

	return &dbConfig{
		Host:            host,
		Port:            port,
		Name:            name,
		Username:        username,
		Password:        password,
		MaxIdle:         maxIdle,
		MaxOpen:         maxOpen,
		SSLMode:         sslMode,
		ConnMaxLifeTime: connMaxLifeTime,
	}, nil
}
