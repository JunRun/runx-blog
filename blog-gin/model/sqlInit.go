/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-23 15:10
 */
package model

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/JunRun/blog-gin/conf"
)

var (
	host     = conf.Config.GetString("database.host")
	user     = conf.Config.GetString("database.user")
	password = conf.Config.GetString("database.password")
	dbname   = conf.Config.GetString("database.dbname")
)
var (
	err    error
	db     *gorm.DB
	config = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbname, password)
	sy     sync.Once
)

func GetConnect() *gorm.DB {

	if db == nil {
		sy.Do(func() {
			if db, err = gorm.Open("postgres", config); err != nil {
				panic(err)
			}
		})
		db.DB().SetMaxOpenConns(100)
		db.DB().SetMaxIdleConns(100)
	}

	return db
}

func init() {
	db := GetConnect()
	if !db.HasTable(&UserModel{}) {
		db.CreateTable(&UserModel{})
	}
}
