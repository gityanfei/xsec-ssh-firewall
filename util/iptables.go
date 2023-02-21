package util

import (
	"go.uber.org/zap"
	"os/exec"
	"xsec-ssh-firewall/settings"
)

var newChain = settings.SettingConfig.UserDefineChain + "_NEW"

func SetNewIptablesChain() {
	exec.Command("/sbin/iptables", "-t", "filter", "-N", newChain).Run()
	for ip, ipCache := range settings.Cache {
		count, ok := ipCache.Get("count")
		if ok {
			if count.(int) >= settings.SettingConfig.MaxFailedCount {
				exec.Command("/sbin/iptables", "-t", "filter", "-A", newChain, "-i", settings.SettingConfig.Interface, "-s", ip, "-j", "DROP").Output()
			}
		} else {
			delete(settings.Cache, ip)
			zap.S().Infof("将IP:%s移出%s\n", ip, settings.SettingConfig.UserDefineChain)
		}
	}
}

// Init iptables policy
func RenameChain() {
	// rename _new to old
	exec.Command("/sbin/iptables", "-t", "filter", "--rename-chain", newChain, settings.SettingConfig.UserDefineChain).Run()
	// set white list chain in filter table
	exec.Command("/sbin/iptables", "-t", "filter", "-A", "INPUT", "-j", settings.SettingConfig.UserDefineChain).Run()
}

// AddPolicy Policy
func AddPolicy(ipAddr string) {
	// Add rule
	exec.Command("/sbin/iptables", "-t", "filter", "-A", settings.SettingConfig.UserDefineChain, "-i", settings.SettingConfig.Interface, "-s", ipAddr, "-j", "DROP").Output()

}

// Delete Policy
func FlushAndDeleteOldChain() {
	// Flush rule
	exec.Command("/sbin/iptables", "-t", "filter", "-D", "INPUT", "-j", settings.SettingConfig.UserDefineChain).Run()
	// delete rules.if not: iptables: Directory not empty.
	exec.Command("/sbin/iptables", "-t", "filter", "-F", settings.SettingConfig.UserDefineChain).Run()
	// delete chain
	exec.Command("/sbin/iptables", "-t", "filter", "-X", settings.SettingConfig.UserDefineChain).Run()
}

func RefreshPolicy() {
	SetNewIptablesChain()
	FlushAndDeleteOldChain()
	RenameChain()
}
