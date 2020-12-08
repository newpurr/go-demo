package main

import (
	"log"
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

var (
	addr1     = "http://127.0.0.1:2379"
	lockName2 = "my-test-lock"
)

func TestMutex(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// etcd地址
	endpoints := strings.Split(addr1, ",")
	// 生成一个etcd client
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	// 生成一个etcd client
	cli2, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		useLock(cli, 1) // 测试锁
		wg.Done()
	}()
	go func() {
		useLock(cli2, 2) // 测试锁
		wg.Done()
	}()

	wg.Wait()
}

func useLock(cli *clientv3.Client, num int) {
	// 为锁生成session
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	// 得到一个分布式锁
	locker := concurrency.NewLocker(s1, lockName2)

	// 请求锁
	log.Println("acquiring lock", num)
	locker.Lock()
	log.Println("acquired lock", num)

	// 等待一段时间
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	locker.Unlock() // 释放锁

	log.Println("released lock", num)
}
