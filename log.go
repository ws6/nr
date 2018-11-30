package nr

import (
	"fmt"
)

const (
	_ = iota
	_INFO
	_WARN
)

var LOG_LEVEL = _INFO

func SetLogLevel(l int) {
	LOG_LEVEL = l
}

func info(s string, a ...interface{}) {
	if LOG_LEVEL < _INFO {
		return
	}
	fmt.Printf(s+"\n", a...)
}

func warn(s string, a ...interface{}) {
	if LOG_LEVEL < _WARN {
		return
	}
	fmt.Printf(s+"\n", a...)
}
