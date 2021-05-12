package dao

import (
	"context"
	"geak/libs/conf"
	"geak/libs/cache"
	"geak/libs/database"
	"github.com/go-redis/redis"
	sql "github.com/jmoiron/sqlx"
)

type Dao struct {
	c 		*conf.Config
	DB		*sql.DB
	RDB 	*redis.Client
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:   c,
		DB:  database.NewMySQL(c.DB),
		RDB: cache.NewRedis(c.Redis),
	}
	return
}

// Ping check service health.
func (d *Dao) Ping(c context.Context) error {
	return d.DB.PingContext(c)
}

// Close close sevice.
func (d *Dao) Close() {
	d.DB.Close()
}
