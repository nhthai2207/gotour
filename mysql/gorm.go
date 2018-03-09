package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"learn/mysql/model"
)

type ORM struct {
	UserName string
	PassWord string
	DbName   string
}


func (orm *ORM) getConnection() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", orm.UserName, orm.PassWord, orm.DbName))
	if err == nil{
		return db
	}
	return nil
}

func (orm *ORM) Migration() {
	db := orm.getConnection()
	defer db.Close()
	db.AutoMigrate(&model.User{})
}