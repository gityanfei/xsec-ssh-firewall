package util

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func SignalHandle() {
	ch := make(chan os.Signal, 1)
	// SIGHUP:1 SIGINT:2 SIGQUIT:3 SIGKILL:9 SIGTERM:15
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM)
	for {
		s := <-ch
		switch s {
		case syscall.SIGHUP:
			zap.S().Infof("\nreceive signal SIGHUP, program will exit,will clean iptables rules\n")
			FlushAndDeleteOldChain()
			os.Exit(1)
		case syscall.SIGINT:
			zap.S().Infof("\nreceive signal SIGINT, program will exit,will clean iptables rules\n")
			FlushAndDeleteOldChain()
			os.Exit(1)
		case syscall.SIGQUIT:
			zap.S().Infof("\nreceive signal SIGQUIT, program will exit,will clean iptables rules\n")
			FlushAndDeleteOldChain()
			os.Exit(1)
		case syscall.SIGKILL:
			zap.S().Infof("\nreceive signal SIGKILL, program will exit,will clean iptables rules\n")
			FlushAndDeleteOldChain()
			os.Exit(1)
		case syscall.SIGTERM:
			zap.S().Infof("\nreceive signal SIGTERM, program will exit,will clean iptables rules\n")
			FlushAndDeleteOldChain()
			os.Exit(0)
		default:
			zap.S().Infof("\nreceive unknown signal, ignore it\n")
		}
	}
}
