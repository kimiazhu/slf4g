// Copyright 2011 ZHU HAIHUA <kimiazhu@gmail.com>.
// All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// Description: slf4g
// Author: ZHU HAIHUA
// Since: 2016-05-02 21:22
package slf4g

import (
	//"log"
)
import (
	"time"
	"runtime"
	"fmt"
)

type ILogger interface {

	//Finest(arg0 interface{}, args ...interface{})
	//
	//Fine(arg0 interface{}, args ...interface{})
	//
	//Debug(arg0 interface{}, args ...interface{})
	//
	//Trace(arg0 interface{}, args ...interface{})
	//
	//Info(arg0 interface{}, args ...interface{})
	//
	//Warn(arg0 interface{}, args ...interface{})
	//
	//Error(arg0 interface{}, args ...interface{})
	//
	//Fatal(arg0 interface{}, args ...interface{})

	Level() Level

	Appenders() []IAppender
}


type Loggers []ILogger

var (
	Global Loggers
)

// Send a formatted log message internally
func (loggers Loggers) writef(lvl Level, format string, args ...interface{}) {
	valid := make(Loggers, 0)
	for i, _ := range loggers {
		if lvl >= loggers[i].Level() {
			valid = append(valid, loggers[i])
		}
	}

	if len(valid) <= 0 {
		return
	}

	// prepare log record
	pc, _, lineno, ok := runtime.Caller(2)
	src := ""
	if ok {
		src = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), lineno)
	}

	msg := format
	if len(args) > 0 {
		msg = fmt.Sprintf(format, args...)
	}

	// Make the log record
	rec := &Record{
		Level:   lvl,
		Created: time.Now(),
		Source:  src,
		Message: msg,
	}

	for i, _ := range valid {
		// Determine caller func
		appenders := loggers[i].Appenders()
		for j, _ := range appenders {
			appenders[j].DoAppend(rec)
		}
	}
}



// Send a closure log message internally
func (loggers Loggers) writec(lvl Level, closure func() string) {

}