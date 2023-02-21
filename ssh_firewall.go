package main

import (
	"go.uber.org/zap"
	"time"
	"xsec-ssh-firewall/settings"
	"xsec-ssh-firewall/util"
)

func main() {
	err := util.InitLogger()
	if err != nil {
		zap.S().Error("FAILED")
	}
	go util.MonitorLog(settings.SettingConfig.SshdLogPath)
	go util.SignalHandle()
	util.Schedule(time.Duration(settings.SettingConfig.GlobalFlushTime))

}
