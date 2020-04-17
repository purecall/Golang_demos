package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()

	//_, err = c.Do("Set", "abc", 100)
	//_, err = c.Do("HSet", "books", "abc", 200)
	//_, err = c.Do("MSet", "xxx", 300, "yyy", 400)
	_, err = c.Do("lpush", "book_list", "abc", "mlf", 1000) // lpush+lpop -> stack // rpop 当作队列用
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	//r, err := redis.Int(c.Do("Get", "abc"))
	//r, err := redis.Int(c.Do("HGet", "books", "abc"))
	//r, err := redis.Ints(c.Do("MGet", "xxx", "yyy"))
	r, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)
}
