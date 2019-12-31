package models

import (
	"database/sql"
	"fmt"
	"go-trading/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	tableNameSignalEvents = "signal_events"
)

var DbConnection *sql.DB

func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}

func init() {
	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver,config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	cmd := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            time datetime PRIMARY KEY NOT NULL,
            product_code varchar(10),
            side varchar(10),
            price float,
            size float)`, tableNameSignalEvents)
	if _,err := DbConnection.Exec(cmd);err!=nil{
		log.Fatalln(err)
	}

	for _, duration := range config.Config.Durations {
		tableName := GetCandleTableName(config.Config.ProductCode, duration)
		c := fmt.Sprintf(`
            CREATE TABLE IF NOT EXISTS %s (
            time datetime PRIMARY KEY NOT NULL,
            open float,
            close float,
            high float,
            low float,
			volume float)`, tableName)
		if _,err := DbConnection.Exec(c);err!=nil{
			log.Fatalln(err)
		}
	}
}
