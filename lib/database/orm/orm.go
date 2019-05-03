package orm

import (
	"github.com/davyxu/golog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// Config mysql config.
type Config struct {
	DSN         string        // data source name.
	Active      int           // pool
	Idle        int           // pool
	IdleTimeout time.Duration // connect max life time.
}

var log = golog.New("goorm")

type ormLog struct{}

func (l ormLog) Print(v ...interface{}) {
	log.Errorln(v)
}

func NewMySQL(c *Config) (db *gorm.DB) {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		log.Errorln("db dsn(%s) error: %v", c.DSN, err)
		panic(err)
	}
	/*db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)*/
	db.SetLogger(ormLog{})
	return
}
