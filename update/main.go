package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID        uint   // 字段`ID`为默认主键
	CreatedBy string `json:"created_by"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedOn time.Time
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now())
	return nil
}

func main() {
	db, err := gorm.Open("mysql", "root:password@(localhost)/gorm-demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.LogMode(true) //打印SQL语句
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	defer db.Close()

	var user User
	db.Where("username=?", "chengang").First(&user)
	db.Model(&user).Update("created_by", "root")
	fmt.Println(user)

}
