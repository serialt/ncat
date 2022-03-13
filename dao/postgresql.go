package dao

import (
	"fmt"
	"time"

	"github.com/serialt/cli/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgreSQLGormDB(mydb *config.Database) *gorm.DB {
	var myGormDB *gorm.DB
	var err error
	// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		mydb.Addr,
		mydb.Username,
		mydb.Password,
		mydb.Addr,
		mydb.Port,
	)
	for {
		myGormDB, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage,
		}), &gorm.Config{Logger: GormLogger})
		if err != nil {
			logSugar.Infof("gorm open failed: %v\n", err)
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	return myGormDB
}
