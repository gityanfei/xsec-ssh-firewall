package util

import (
	"github.com/hpcloud/tail"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"regexp"
	"time"
	"xsec-ssh-firewall/settings"
)

func MonitorLog(path string) {
	t, err := tail.TailFile(path, tail.Config{Follow: true})
	if err != nil {
		panic("fail to open file！！" + err.Error())
	}
	for line := range t.Lines {
		CheckSSH(line)
	}
}

func CheckSSH(logContent *tail.Line) {
	content := logContent.Text
	// Multiple matching methods can be customized. There are multiple rejection logs in auth.log, so it is more complete to specify manually
	for _, s := range settings.SettingConfig.ErrorLogREGX {
		re, _ := regexp.Compile(s)
		ret := re.FindStringSubmatch(content)
		REGEX(ret)
	}
}
func REGEX(ret []string) {
	// If the matching error log is successful (with brute force cracking), the expression can match the attacker's public IP address
	if len(ret) > 0 {
		ip := ret[1]
		if _, ok := settings.WhiteIPlist[ip]; ok {
			zap.S().Infof("IP:%s Ignore this IP in the white list\n", ip)
			return
		}
		if c, ok := settings.Cache[ip]; ok {
			// 没有过期
			if _, ook := c.Get("count"); ook {
				count, _ := c.Get("count")
				t := count.(int)
				if t == settings.SettingConfig.MaxFailedCount {
					zap.S().Infof("Lock ip: %v\n", ip)
					AddPolicy(ip)
				}
				c.IncrementInt("count", 1)
			} else {
				//如果过期了，重新添加进入黑名单设置
				settings.Cache[ip].Set("count", 1, time.Duration(settings.SettingConfig.LockTime)*time.Second)
				if 1 == settings.SettingConfig.MaxFailedCount {
					zap.S().Infof("Lock ip: %v\n", ip)
					AddPolicy(ip)
				}
			}
		} else {
			//如果这个ip是第一次匹配到（第一次攻击）
			n := cache.New(time.Duration(settings.SettingConfig.LockTime)*time.Second, 10*time.Second)
			n.Set("count", 1, time.Duration(settings.SettingConfig.LockTime)*time.Second)
			if 1 == settings.SettingConfig.MaxFailedCount {
				zap.S().Infof("Lock ip: %v\n", ip)
				AddPolicy(ip)
			}
			settings.Cache[ip] = n
		}
	}
}
