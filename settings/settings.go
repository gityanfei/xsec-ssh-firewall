package settings

import (
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"xsec-ssh-firewall/declare"
)

var (
	SettingConfig = declare.Config{}
	// The map will not change after initialization, so there is no need to make a structure to add locks
	WhiteIPlist = make(map[string]bool)
	Cache       = make(map[string]*cache.Cache)
)

func GetYamlConfig() (err error) {
	vip := viper.New()
	vip.AddConfigPath("./")     //设置读取的文件路径
	vip.SetConfigName("config") //设置读取的文件名
	vip.SetConfigType("yaml")   //设置文件的类型
	//尝试进行配置读取
	if err = vip.ReadInConfig(); err != nil {
		return err
	}
	err = vip.Unmarshal(&SettingConfig)
	if err != nil {
		return err
	}
	return
}
func init() {

	err := GetYamlConfig()
	if err != nil {
		panic("解析yaml报错:" + err.Error())
	}
	zap.S().Info(SettingConfig)
	for _, ip := range SettingConfig.WhiteIpList {
		WhiteIPlist[ip] = true
	}
	zap.S().Info("读取配置成功%#v\n", SettingConfig)
}
