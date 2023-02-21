package declare

type Config struct {
	Interface       string    `json:"interface"`
	LockTime        int64     `json:"lockTime"`
	MaxFailedCount  int       `json:"maxFailedCount"`
	WhiteIpList     []string  `json:"whiteIpList"`
	SshdLogPath     string    `json:"sshdLogPath"`
	ErrorLogREGX    []string  `json:"errorLogREGX"`
	UserDefineChain string    `json:"userDefineChain"`
	GlobalFlushTime int       `json:"globalFlushTime"`
	LogConfig       LogConfig `json:logConfig`
}
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}
