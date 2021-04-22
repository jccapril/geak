package dao

import (
	"geak/cache"
	"geak/database"
	"geak/job/conf"
	"github.com/go-redis/redis"
	sql "github.com/jmoiron/sqlx"
)

type Dao struct {
	c 		*conf.Config
	db		*sql.DB
	rdb 	*redis.Client
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:  c,
		db: database.NewMySQL(c.DB),
		rdb: cache.NewRedis(c.Redis),
	}
	return
}

// Ping check service health.
func (d *Dao) Ping() error {
	return d.db.Ping()
}

// Close close sevice.
func (d *Dao) Close() {
	d.db.Close()
}
