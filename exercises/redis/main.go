package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"lib/myredis"
)

func main() {
	conn, err := redis.Dial("tcp","redis:6379",redis.DialPassword("123456"))
	myRedis := myredis.MyRedis{}
	if err != nil {
		fmt.Println("redis err", err)
		return
	}
	defer conn.Close()

	if myRedis.Set(conn,"name","haohonglong777") {
		fmt.Println("set success")
	}

	if myRedis.Hset(conn,"001","name","88888") {
		fmt.Println("hset success")
	}



	


}