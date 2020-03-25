package logger

import "github.com/sirupsen/logrus/hooks/test"

var hook *test.Hook
var testLoggerSetup bool

func SetupTestLogger() {
	if testLoggerSetup {
		return
	}
	logger, hook = test.NewNullLogger()
	testLoggerSetup = true
}

func Hook() *test.Hook {
	return hook
}
