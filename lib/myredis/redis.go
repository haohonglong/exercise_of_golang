package myredis

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

type MyRedis struct {}
/*
func GetRedis() (redis.Conn, error){
	conn, err := redis.Dial("tcp","redis:6379",redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("redis err", err)
		return err
	}
	return conn
}
*/
func (r MyRedis) Set(conn redis.Conn, key string, value string) bool{
	_, err := conn.Do("set",key,value)
	if err != nil {
		fmt.Println("set err=", err)
		return false
	}

	return true
}

func (r MyRedis) Hset(conn redis.Conn, key string,field string, value string) bool{
	_, err := conn.Do("hset",key,field,value)
	if err != nil {
		fmt.Println("hset err=", err)
		return false
	}

	return true

}