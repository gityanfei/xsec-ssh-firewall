# xsec-ssh-firweall

xsec-ssh-firewall，源项目为https://github.com/netxfly/xsec-ssh-firewall，这里对源项目做了一些更改
为一个简易的SSH密码防暴力破解程序，通过读取SSH的`auth.log`判断恶意IP，执行`iptables`进行封杀。本项目采用Go语言开发。

## 使用方法

设置好配置文件`conf/app.yaml`后直接启动`./xsec-ssh-firewall`即可。

建议使用`supvervisor`或`nohup`或screen将其跑在后台。

```yaml
interface: eth0
lockTime: 86400
maxFailedCount: 1
whiteIpList:
  - 10.0.16.15
  - 127.0.0.1
  - 8.8.8.8
# ubuntu使用这个配置
sshdLogPath: /var/log/auth.log
# centos使用这个配置
#sshdLogPath: /var/log/secure
errorLogREGX:
  - ^.*Invalid user.*from (.*) port.*$
  - ^.*Connection closed by authenticating user [a-zA-Z0-9]+ (.*) port.*$
userDefineChain: BLACKLIST
globalFlushTime: 600
logConfig:
  level: "debug"
  filename: "./run.log"
  maxsize: 200
  maxage: 7
  maxbackups: 10
```

### 配置文件说明：

1. interface表示对外暴露ssh端口绑定的IP，设置防火墙规则时需要用到
2. lockTime表示对暴力破解来源IP的封禁时间，单位为秒
3. whiteIpList为白名单，防止自己输错密码后无法登录服务器
4. sshdLogPath为ssh log的位置，ubuntu为：/var/log/auth.log centos为：/var/log/secure
5. errorLogREGXssh错误日志中有多种错误日志，通过不同的正则表达式去匹配IP(.*)中取的就是攻击方IP地址
