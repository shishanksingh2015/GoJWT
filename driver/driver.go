package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shishanksingh2015/GoJWT/config"
	"github.com/spf13/viper"
	"log"
)

func ConnectDb() *sql.DB {
	var configuration config.Configurations
	if err :=viper.ReadInConfig();err!=nil{
		fmt.Printf("Error Reading file %s",err)
	}
	err:= viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",configuration.Database.DBUser,configuration.Database.DBPassword,configuration.Database.DBName))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil
	}
	return db
}
