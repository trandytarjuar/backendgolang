package config
import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	host := "localhost"
    port := "3306"
    dbname := "teskerjagolang"
    username := "root"
    password := ""
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
    if err != nil {
        panic("Cannot connect database")
    }
    DB.AutoMigrate()
}