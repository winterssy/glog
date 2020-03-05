package glog_test

import (
	"os"

	"github.com/winterssy/glog"
)

func ExampleReplaceGlobal() {
	glog.ReplaceGlobal(glog.New(os.Stderr, "", glog.LstdFlags, glog.Ldebug))
}

func ExampleInfo() {
	glog.Info("hello world")
}

func ExampleInfof() {
	glog.Infof("hello %s", "world")
}
