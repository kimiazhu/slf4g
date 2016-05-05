// Copyright 2011 ZHU HAIHUA <kimiazhu@gmail.com>.
// All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// Description: slf4g
// Author: ZHU HAIHUA
// Since: 2016-05-02 21:35
package slf4g

import (
	"fmt"
	"strings"
	"errors"
	. "github.com/kimiazhu/golib/stack"
)

func init() {

}

func writeLog(lvl Level, arg0 interface{}, args ...interface{}) {
	switch first := arg0.(type) {
	case string:
		// Use the string as a format string
		Global.writef(lvl, first, args...)
	case func() string:
		// Log the closure (no other arguments used)
		Global.writec(lvl, first)
	default:
		// Build a format string so that it will be similar to Sprint
		Global.writef(lvl, fmt.Sprint(arg0)+strings.Repeat(" %v", len(args)), args...)
	}
}

func writeAndReturnErr(lvl Level, arg0 interface{}, args ...interface{}) error {
	switch first := arg0.(type) {
	case string:
		// Use the string as a format string
		Global.writef(lvl, first, args...)
		return errors.New(fmt.Sprintf(first, args...))
	case func() string:
		// Log the closure (no other arguments used)
		str := first()
		Global.writef(lvl, "%s", str)
		return errors.New(str)
	default:
		// Build a format string so that it will be similar to Sprint
		Global.writef(lvl, fmt.Sprint(first)+strings.Repeat(" %v", len(args)), args...)
		return errors.New(fmt.Sprint(first) + fmt.Sprintf(strings.Repeat(" %v", len(args)), args...))
	}
}

func Finest(arg0 interface{}, args ...interface{}) {
	writeLog(FINEST, arg0, args)
}

func Fine(arg0 interface{}, args ...interface{}) {
	writeLog(FINE, arg0, args)
}

func Trace(arg0 interface{}, args ...interface{}) {
	writeLog(TRACE, arg0, args)
}

func Debug(arg0 interface{}, args ...interface{}) {
	writeLog(DEBUG, arg0, args)
}

func Info(arg0 interface{}, args ...interface{}) {
	writeLog(INFO, arg0, args)
}

func Warn(arg0 interface{}, args ...interface{}) error {
	return writeAndReturnErr(WARNING, arg0, args)
}

func Error(arg0 interface{}, args ...interface{}) error {
	return writeAndReturnErr(ERROR, arg0, args)
}

func Critical(arg0 interface{}, args ...interface{}) error {
	const lvl = CRITICAL
	switch first := arg0.(type) {
	case string:
		// Use the string as a format string
		msg := fmt.Sprintf("%s\n%s", fmt.Sprintf(first, args...), CallStack(3))
		Global.writef(lvl, msg)
		//Global.intLogf(lvl, "%s", CallStack(3))
		return errors.New(fmt.Sprintf(first, args...))
	case func() string:
		// Log the closure (no other arguments used)
		str := first()
		Global.writef(lvl, "%s\n%s", str, CallStack(3))
		//Global.intLogf(lvl, "%s", CallStack(3))
		return errors.New(str)
	case func(interface{}) string:
		str := first(args[0])
		Global.writef(lvl, "%s\n%s", str, CallStack(3))
		return errors.New(str)
	default:
		// Build a format string so that it will be similar to Sprint
		msg := fmt.Sprintf("%s\n%s", fmt.Sprint(first) + fmt.Sprintf(strings.Repeat(" %v", len(args)), args...), CallStack(3))
		Global.writef(lvl, msg)
		return errors.New(fmt.Sprint(first) + fmt.Sprintf(strings.Repeat(" %v", len(args)), args...))
	}
	return nil
}

// Recover used to log the stack when panic occur.
// usage: defer log4go.Recover("this is a msg: %v", "msg")
// or:
//      defer log4go.Recover(func(err interface{}) string {
//          // ... your code here, return the error message
//          return fmt.Sprintf("recover..v1=%v;v2=%v;err=%v", 1, 2, err)
//      })
func Recover(arg0 interface{}, args ...interface{}) {
	if err := recover(); err != nil {
		switch a := arg0.(type) {
		case func(interface{}) string:
			// the recovered err will pass to this func
			Critical(arg0, append([]interface{}{err}, args)...)
		case string:
			Critical(a+"\n%v", append(args, err)...)
		default:
			Critical(arg0, append(args, err)...)
		}
	}
}

func IsFinestEnabled() bool {
	return isLevelEnabled(FINEST)
}

func IsFineEnabled() bool {
	return isLevelEnabled(FINE)
}

func IsDebugEnabled() bool {
	return isLevelEnabled(DEBUG)
}

func IsTraceEnabled() bool {
	return isLevelEnabled(TRACE)
}

func IsInfoEnabled() bool {
	return isLevelEnabled(INFO)
}

func IsWarnEnabled() bool {
	return isLevelEnabled(WARNING)
}

func IsErrorEnabled() bool {
	return isLevelEnabled(ERROR)
}

func isLevelEnabled(lvl Level) bool {
	return len(Global.findValid(lvl)) > 0
}