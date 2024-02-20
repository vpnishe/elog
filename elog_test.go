package elog_test

import (
	"testing"
	"time"

	"github.com/vpnishe/elog"
)

var tlog *elog.EasyLogger

func TestGetLogger(t *testing.T) {
	tlog = elog.GetLogger()
	defer tlog.Flush()

	tlog.Info("Hello")
	tlog.Info("Hello")
}

func TestElog(t *testing.T) {
	defer elog.Flush()

	elog.Info("Hello")
	elog.Info("Hello")
	elog.Infof("%d,%v,%s", 1, "xx", "xxxxxxx")

	time.Sleep(time.Second * 10)
}
