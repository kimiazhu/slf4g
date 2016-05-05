// Copyright 2011 ZHU HAIHUA <kimiazhu@gmail.com>.
// All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// Description: slf4g
// Author: ZHU HAIHUA
// Since: 2016-05-02 21:22
package slf4g

import (
	"time"
	"runtime"
	"fmt"
)

type ILogger interface {

	Filter() IFilter

	Level() Level

	Appenders() []IAppender
}


type Loggers []ILogger

var (
	Global Loggers
)

type DefaultLogger struct {
	Filter IFilter
	Lvl Level
	Appenders []IAppender
}

func (l *DefaultLogger) Filter() IFilter {
	return l.Filter
}

func (l *DefaultLogger) Level() Level {
	return l.Lvl
}

func (l *DefaultLogger) Appenders() []IAppender {
	return l.Appenders
}

func (loggers Loggers) findValid(lvl Level) Loggers {
	valid := make(Loggers, 0)
	for i, _ := range loggers {
		switch loggers[i].Filter().Decide(lvl) {
		case ACCEPT, NEUTRAL:
			valid = append(valid, loggers[i])
		}
	}
	return valid
}

// Send a formatted log message internally
func (loggers Loggers) writef(lvl Level, format string, args ...interface{}) {
	valid := loggers.findValid(lvl)

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
	valid := loggers.findValid(lvl)

	if len(valid) <= 0 {
		return
	}

	// prepare log record
	pc, _, lineno, ok := runtime.Caller(2)
	src := ""
	if ok {
		src = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), lineno)
	}

	// Make the log record
	rec := &Record{
		Level:   lvl,
		Created: time.Now(),
		Source:  src,
		Message: closure(),
	}

	for i, _ := range valid {
		// Determine caller func
		appenders := loggers[i].Appenders()
		for j, _ := range appenders {
			appenders[j].DoAppend(rec)
		}
	}
}