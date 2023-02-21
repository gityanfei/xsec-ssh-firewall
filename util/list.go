package util

import (
	"fmt"
	"time"
	"xsec-ssh-firewall/settings"
)

func CacheList(t time.Duration) {
	for {
		fmt.Println("【当前缓存中的黑名单IP为】")
		for k, v := range settings.Cache {
			a, _ := v.Get("count")
			fmt.Printf("IP:%s,次数:%v\n", k, a.(int))
		}
		time.Sleep(t * time.Second)
	}
}
