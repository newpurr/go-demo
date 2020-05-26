package redis

import (
	"fmt"
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	c, err := redis.Dial("tcp", "192.168.0.212:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer func() {
		_ = c.Close()
	}()
	fmt.Println("redis conn success")

	t.Run("SET AND GET", func(t *testing.T) {
		i := 100
		_, err = c.Do("Set", "abc", i)
		if err != nil {
			fmt.Println(err)
			return
		}

		r, err := redis.Int(c.Do("Get", "abc"))
		if err != nil {
			fmt.Println("get abc failed,", err)
			return
		}

		assert.True(t, r == i)
	})

	t.Run("MSET AND MGET", func(t *testing.T) {
		_, err = c.Do("MSet", "abc", 100, "efg", 300)
		if err != nil {
			fmt.Println(err)
			return
		}

		r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
		if err != nil {
			fmt.Println("get abc failed,", err)
			return
		}

		for i, v := range r {
			fmt.Println(i, v)
		}
	})

	t.Run("Expire", func(t *testing.T) {
		_, err = c.Do("expire", "abc", 10)
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	t.Run("list 操作", func(t *testing.T) {
		_, err = c.Do("lpush", "book_list", "abc", "ceg", 300)
		if err != nil {
			fmt.Println(err)
			return
		}

		r, err := redis.String(c.Do("lpop", "book_list"))
		if err != nil {
			fmt.Println("get abc failed,", err)
			return
		}

		fmt.Println(r)
	})

	t.Run("hash", func(t *testing.T) {
		_, err = c.Do("HSet", "books", "abc", 100)
		if err != nil {
			fmt.Println(err)
			return
		}

		r, err := redis.Int(c.Do("HGet", "books", "abc"))
		if err != nil {
			fmt.Println("get abc failed,", err)
			return
		}

		fmt.Println(r)
	})

	t.Run("Redis Pool", func(t *testing.T) {
		pool := &redis.Pool{ // 实例化一个连接池
			MaxIdle: 16, // 最初的连接数量
			// MaxActive:1000000,    //最大连接数量
			MaxActive:   0,   // 连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
			IdleTimeout: 300, // 连接关闭时间 300秒 （300秒不使用自动关闭）
			Dial: func() (redis.Conn, error) { // 要连接的redis数据库
				return redis.Dial("tcp", "192.168.0.212:6379")
			},
		}
		defer func() {
			// 关闭连接池
			_ = pool.Close()
		}()

		c := pool.Get() // 从连接池，取一个链接
		defer func() {
			// 函数运行结束 ，把连接放回连接池
			_ = c.Close()
		}()

		_, err := c.Do("Set", "abc", 200)
		if err != nil {
			fmt.Println(err)
			return
		}

		r, err := redis.Int(c.Do("Get", "abc"))
		if err != nil {
			fmt.Println("get abc faild :", err)
			return
		}
		fmt.Println(r)
	})
}
