package database

import (
	"database/sql"
	"fmt"

	"github.com/fernandocastrotelco/gobackend/pasarela/internal/config"
	"github.com/fernandocastrotelco/gobackend/pasarela/internal/logs"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MySQLClient contiene la instancia de base de datos
type MySQLClient struct {
	*gorm.DB
}

// NewMySQLClient cliente de la base de datos en MySql
func NewMySQLClient() *MySQLClient {
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.USER, config.PASSW, config.HOST, config.PORT, config.DB))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		logs.Error("cannot create mysql client")
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		logs.Error("cannot connect database" + err.Error())
	}

	return &MySQLClient{gormDB}
}

func NewUsuariosClient() *MySQLClient {
	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.USER, config.PASSW, config.HOST, config.PORT, "loginservice"))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		logs.Error("cannot create loginservice sql client")
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		logs.Error("cannot connect database" + err.Error())
	}

	return &MySQLClient{gormDB}
}
