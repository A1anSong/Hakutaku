package model

import (
	"fmt"
	"server/config"
	"server/hook"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	hook.Register(hook.LoadConfigHook, "db", initDB)
}

func initDB(v interface{}) (err error) {
	cfg := v.(config.Config)
	switch cfg.DB.Driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(cfg.DB.Uri))
	case "postgresql":
		err = fmt.Errorf("unsupport %s driver", cfg.DB.Driver)
	default:
		err = fmt.Errorf("unsupport %s driver", cfg.DB.Driver)
	}
	if err == nil {
		hook.OnEvent(hook.LoadDBHook, db)
	}
	return
}
