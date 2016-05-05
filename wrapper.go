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

func Finest(arg0 interface{}, args ...interface{}) {
	writeLog(FINEST, arg0, args)
}

func Fine(arg0 interface{}, args ...interface{}) {
	writeLog(FINE, arg0, args)
}

func Debug(arg0 interface{}, args ...interface{}) {
	writeLog(DEBUG, arg0, args)
}

func Info(arg0 interface{}, args ...interface{}) {
	writeLog(INFO, arg0, args)
}