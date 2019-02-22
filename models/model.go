package models

import (
	"fmt"
	"github.com/TopJohn/go-web-framework/pkg/conf"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

var db *gorm.DB

func Setup() {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		conf.DatabaseConfig.Host,
		conf.DatabaseConfig.Port,
		conf.DatabaseConfig.User,
		conf.DatabaseConfig.Name,
		conf.DatabaseConfig.Password))

	if err != nil {
		logs.Error("model.Setup error: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.DatabaseConfig.TablePrefix + defaultTableName;
	}
	// Disable table name's pluralization, if set to true, `User`'s table name will be `user`
	db.SingularTable(true)
}

func CloseDB() {
	defer db.Close()
}
