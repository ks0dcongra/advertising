package configs

import (
	"fmt"
)

// add some function to get env setting
type config struct {
	DB    *dbConfig
	Redis *redisConfig
}

var (
	Cfg config
	ServiceInfo *commonConfig
)

func InitConfigs() error {

	var err error

	Cfg.DB, err = getDBConfig()
	if err != nil {
		fmt.Println("initialize the DB configuration failure")
		return err
	}

	Cfg.Redis, err = getRedisConfig()
	if err != nil {
		fmt.Println("initialize the DB configuration failure")
		return err
	}

	return nil
}

func InitCommonConf() (err error) {
	ServiceInfo, err = getCommonConfig()
	if err != nil {
		fmt.Println("initialize the common configuration failure")
		return err
	}
	return nil
}
