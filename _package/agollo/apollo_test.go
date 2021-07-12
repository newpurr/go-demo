package agollo

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"strings"
	"testing"
	"time"
)

type CustomChangeListener struct {
}

func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	fmt.Println(changeEvent)
}

func (c *CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	fmt.Println(event)
}

func TestApollo(t *testing.T) {
	c := &config.AppConfig{
		AppID:          "testApplication_yang",
		Cluster:        "dev",
		IP:             "http://106.54.227.205:8080",
		NamespaceName:  "dubbo,product.joe",
		IsBackupConfig: true,
		Secret:         "6ce3ff7e96a24335a9634fe9abca6d51",
	}
	agollo.SetLogger(&DefaultLogger{})

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}

	client.AddChangeListener(&CustomChangeListener{})

	split := strings.Split(c.NamespaceName, ",")
	for _, n := range split {
		checkKey(n, client)
	}

	time.Sleep(5 * time.Second)
}

func checkKey(namespace string, client *agollo.Client) {
	cache := client.GetConfigCache(namespace)
	count := 0
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value)
		count++
		return true
	})
	if count < 1 {
		panic("config key can not be null")
	}
}

type DefaultLogger struct {
}

func (d *DefaultLogger) Debugf(format string, params ...interface{}) {
	d.Debug(format, params)
}

func (d *DefaultLogger) Infof(format string, params ...interface{}) {
	d.Debug(format, params)
}

func (d *DefaultLogger) Warnf(format string, params ...interface{}) {
	d.Debug(format, params)
}

func (d *DefaultLogger) Errorf(format string, params ...interface{}) {
	d.Debug(format, params)
}

func (d *DefaultLogger) Debug(v ...interface{}) {
	//fmt.Println(v)
}
func (d *DefaultLogger) Info(v ...interface{}) {
	d.Debug(v)
}

func (d *DefaultLogger) Warn(v ...interface{}) {
	d.Debug(v)
}

func (d *DefaultLogger) Error(v ...interface{}) {
	d.Debug(v)
}
