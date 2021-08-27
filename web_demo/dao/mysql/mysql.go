package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	set "turan.com/web_demo/settings"
)
import "github.com/jinzhu/gorm"

var mysql *gorm.DB

func Init() (err error) {
	dst := set.MysqlC.User + ":" + set.MysqlC.Password + "@(" + set.MysqlC.Host + ":" + fmt.Sprintf("%v", set.MysqlC.Port) + ")/" + set.MysqlC.Dbname + "?charset=utf8&parseTime=True&loc=Local"
	log.Println(dst)
	mysql, err = gorm.Open("mysql", dst)
	if err != nil {
		panic(err)
		return
	}
	return
}

func Close() {
	_ = mysql.Close()
}
