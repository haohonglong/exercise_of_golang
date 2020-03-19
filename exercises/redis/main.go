package main

import (
	"fmt"
	"redis"
	"console"
)

func main() {
	myRedis := redis.Redis{}
	conn, err := myRedis.Connection()
	if err != nil {
		fmt.Printf("%s redis err=\n",console.Log(), err)
		return
	}
	defer conn.Close()

	if err := myRedis.Set("name","haohonglong777"); err != nil {
		fmt.Printf("%s set success\n",console.Log())
	}

	if err := myRedis.Hset("001","name","999"); err != nil {
		fmt.Printf("%s hset success\n",console.Log())
	}



	


}