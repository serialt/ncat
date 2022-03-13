package dao

import (
	"fmt"
	"time"

	"github.com/serialt/cli/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMysqlGormDB(mydb *config.Database) *gorm.DB {
	var myGormDB *gorm.DB
	var err error
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mydb.Username,
		mydb.Password,
		mydb.Addr,
		mydb.Port,
		mydb.DBName,
	)
	for {
		myGormDB, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{Logger: GormLogger})
		if err != nil {
			logSugar.Infof("grom open failed: %v\n", err)
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	return myGormDB
}
