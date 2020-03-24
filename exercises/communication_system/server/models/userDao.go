package models

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"console"
)
var (
	MyUserDao *UserDao
)

type UserDao struct {
	Pool *redis.Pool
	Conn redis.Conn
}
//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		Pool : pool,
		Conn : pool.Get(),
	}
	return
}

//根据用户id返回一个User实例
func (this *UserDao) getUserById(id int) (user *User, err error){
	//通过给定id去redis查询这个用户
	res, err := redis.String(this.Conn.Do("hget", "users", id))
	if err != nil {
		if err == redis.ErrNil {//user hash 中没有找到对于的id
			err = ERROR_USER_NOTEXIST
		}
		return
	}

	user = &User{}
	//这里把res 反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Printf("%s json.Unmarshal error=%s\n",console.Log(), err)
	}
	return


}
//完成登录的校验
//1.Login 完成对用户的验证
//2.如果用户的id 和pwd都正确，则返回一个user实例
//3.如果用户的id或pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.Pool.Get()
	defer conn.Close()
	this.Conn = conn
	user, err = this.getUserById(userId)
	if err != nil {
		fmt.Printf("%s this.getUserById\n",console.Log(), err)
		return
	}

	if user.UserPwd != userPwd {//密码错误
		err = ERROR_USER_PWD
		return
	}
	return
}
