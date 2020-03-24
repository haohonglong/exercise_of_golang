package redis

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
	"console"

)
var (
	Host string = "redis:6379"
	Pwd string = "123456"
	MaxIdle int = 8
	MaxActive int = 0
	IdleTimeout time.Duration

	pool *redis.Pool
    conn redis.Conn
)
type Redis struct {
	Conn redis.Conn
	Pool *redis.Pool
}



func InitPool(maxIdle, 
			maxActive int,
			idleTimeout time.Duration) (pool *redis.Pool){
	pool = &redis.Pool{
		MaxIdle : maxIdle,
		MaxActive : maxActive,
		IdleTimeout : idleTimeout,
		Dial : func() (redis.Conn, error) {
			r := &Redis{}
			return r.Connect()
		},
	}
	return
}

func GetPool() (pool *redis.Pool){
	return pool
}



func (this *Redis) Connect() (conn redis.Conn, err error){
	conn, err = redis.Dial("tcp",Host,redis.DialPassword(Pwd))
	if err != nil {
		fmt.Printf("%s redis err=%s",console.Log(), err)
		return
	}
	this.Conn = conn
	return 
}

func (this *Redis) Set(key, value string) (err error){
	_, err = this.Conn.Do("set", key, value)
	return 
}

func (this *Redis) Hset(key, field, value string) (err error){
	_, err = this.Conn.Do("hset", key, field, value)
	return
}