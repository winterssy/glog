package glog

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Rdate         = `[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]`
	Rtime         = `[0-9][0-9]:[0-9][0-9]:[0-9][0-9]`
	Rmilliseconds = `\,[0-9][0-9][0-9]`
	Rlevel        = `\[(DEBUG|INFO|WARN|ERROR|PANIC|FATAL)\]`
	Rmessage      = `(.*?)`
)

var RstdFlags = regexp.MustCompile("^" + Rdate + " " + Rtime + Rmilliseconds + " " + Rlevel + " " + Rmessage + "\n$")

func TestLogger_Output(t *testing.T) {
	var sb strings.Builder
	logger := New(&sb, "", LstdFlags, Ldebug)

	var (
		msg       = "hello world"
		format    = "hello %s"
		arg       = "world"
		formatMsg = fmt.Sprintf(format, arg)
	)

	tests := []struct {
		level     int
		output    func(v ...interface{})
		outputF   func(format string, v ...interface{})
		needPanic bool
	}{
		{
			level:   Ldebug,
			output:  logger.Debug,
			outputF: logger.Debugf,
		},
		{
			level:   Linfo,
			output:  logger.Info,
			outputF: logger.Infof,
		},
		{
			level:   Lwarn,
			output:  logger.Warn,
			outputF: logger.Warnf,
		},
		{
			level:   Lerror,
			output:  logger.Error,
			outputF: logger.Errorf,
		},
		{
			level:     Lpanic,
			output:    logger.Panic,
			outputF:   logger.Panicf,
			needPanic: true,
		},
	}

	for _, test := range tests {
		sb.Reset()
		if test.needPanic {
			assert.Panics(t, func() {
				test.output(msg)
			})
		} else {
			test.output(msg)
		}
		vs := RstdFlags.FindStringSubmatch(sb.String())
		if assert.Len(t, vs, 3) {
			assert.Equal(t, levels[test.level], vs[1])
			assert.Equal(t, msg, vs[2])
		}

		sb.Reset()
		if test.needPanic {
			assert.Panics(t, func() {
				test.outputF(format, arg)
			})
		} else {
			test.outputF(format, arg)
		}
		vs = RstdFlags.FindStringSubmatch(sb.String())
		if assert.Len(t, vs, 3) {
			assert.Equal(t, levels[test.level], vs[1])
			assert.Equal(t, formatMsg, vs[2])
		}
	}
}

func BenchmarkInfo(b *testing.B) {
	const dummyData = "hello world"
	logger := New(ioutil.Discard, "", LstdFlags, Ldebug)
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			logger.Info(dummyData)
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkStdPrintln(b *testing.B) {
	const dummyData = "hello world"
	logger := log.New(ioutil.Discard, "", log.LstdFlags)
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			logger.Println(dummyData)
			wg.Done()
		}()
	}
	wg.Wait()
}
