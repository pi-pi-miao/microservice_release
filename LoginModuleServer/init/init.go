package init

import (
	"LoginModuleServer/config"
	"LoginModuleServer/server"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	config.Conf.MysqlConf.Username = "root"
	config.Conf.MysqlConf.Password = "123456"
	config.Conf.MysqlConf.Host = "localhost"
	config.Conf.MysqlConf.Port = 3306
	config.Conf.MysqlConf.Database = "microservice_user"

}

func InitDb() error {
	d := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Conf.MysqlConf.Username,
		config.Conf.MysqlConf.Password, config.Conf.MysqlConf.Host, config.Conf.MysqlConf.Port,
		config.Conf.MysqlConf.Database)
	database, err := sqlx.Open("mysql", d)
	if err != nil {
		fmt.Println("sqlx open mysql err", err)
		return err
	}
	Db = database
	return nil
}

func Initialize() (err error) {

	err = InitDb()
	if err != nil {
		fmt.Println("init err")
		return
	}
	server.Init(Db)
	return nil
}