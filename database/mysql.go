package database

import (
	"database/sql"
	"fmt"
	"time"

	//"github.com/go-sql-driver/mysql"
	//_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	//"github.com/go-sql-driver/mysql"
	//_ "github.com/go-sql-driver/mysql" //加载mysql驱动
	"gorm.io/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
mysqlUrl ="root:12345678@tcp(127.0.0.1:3306)/golang-demo?charset=utf8&parseTime=True&loc=Local"
DB *gorm.DB
)
func InitDb()   *sql.DB{
	var err,err2 error
	DB, err = gorm.Open(mysql.Open(mysqlUrl),&gorm.Config{

	} )


	//GORM使用数据库/sql维护连接池
	sqlDB, err2 := DB.DB()
	// SetMaxIdleConns设置空闲连接池中的最大连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns设置数据库的最大打开连接数。
	sqlDB.SetMaxOpenConns(100)// SetConnMaxLifetime设置连接可能被重用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute)
	//sqlDB.SetConnMaxIdleTime(0)
	//DB.Set("ConnMaxLifetime",time.Minute * 4)
	//defer sqlDB.Close()

	if err != nil {
		fmt.Printf("mysql connect error %v", err.Error())
		//panic("连接数据库失败")
	}

	if err2 != nil {
		fmt.Printf("mysql connect error2 %v", err2)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
	return sqlDB
}
