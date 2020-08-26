package initialize

import (
	"github.com/gomodule/redigo/redis"
	"github.com/zhenghuajing/fresh_shop/global"
	"time"
)

func Redis() error {
	redisCfg := global.Config.Redis
	global.RedisPool = &redis.Pool{
		MaxIdle:     redisCfg.MaxIdle,
		MaxActive:   redisCfg.MaxActive,
		IdleTimeout: redisCfg.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisCfg.Host)
			if err != nil {
				return nil, err
			}
			if redisCfg.Password != "" {
				if _, err := c.Do("AUTH", redisCfg.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}
