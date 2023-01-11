package models

import (
	"github.com/haimait/go-mindoc/pkg/component"
	db "github.com/haimait/go-mindoc/pkg/components/DB"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB(dsn string) {

	if dsn == "" {
		//dsn为空时，从配置文件中读取
		//加载yaml中的配置项目到传一的结构体中，如DB / Redis
		component.LoadConfigOption("user/rpc/etc/user.yaml", &db.Instance{})
		dsn = db.Get().DataSource
	}
	dbClient, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	//迁移表，在go-admin里做
	//err = db.AutoMigrate(&SysUser{})
	//if err != nil {
	//	log.Fatalln("[DB ERROR] : ", err)
	//}
	DB = dbClient
}
