// Copyright 2011 ZHU HAIHUA <kimiazhu@gmail.com>.
// All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// Description: slf4g
// Author: ZHU HAIHUA
// Since: 2016-05-02 23:10
package slf4g

import "time"

type Level int

const (
	FINEST Level = iota
	FINE
	TRACE
	DEBUG
	INFO
	WARNING
	ERROR
	CRITICAL
)

// Logging level strings
var (
	levelStrings = [...]string{"FNST", "FINE", "TRAC", "DEBG", "INFO", "WARN", "EROR", "CRIT"}
)

// A Record contains all of the pertinent information for each message
type Record struct {
	Level   Level     // The log level
	Created time.Time // The time at which the log message was created (nanoseconds)
	Source  string    // The message source
	Message string    // The log message
}