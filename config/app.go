package config

import (
	"github.com/rnikrozoft/himymonsters-master-backend/service"
	"github.com/uptrace/bun"
)

type Config struct {
	App  App                 `mapstructure:"app"`
	Shop service.ShopService `mapstructure:"shop"`
}

type App struct {
	Database string `mapstructure:"database"`
}

var Database *bun.DB
