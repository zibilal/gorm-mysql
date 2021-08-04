package dbengine

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbEngine struct {
	Db *gorm.DB
}

func InitDbEngine(hostname, user, password, dbname string, port int)(*DbEngine, error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, hostname, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	ctx := &DbEngine{}
	ctx.Db = db

	return ctx, err
}

