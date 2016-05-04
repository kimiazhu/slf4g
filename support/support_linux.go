// Copyright 2011 ZHU HAIHUA <kimiazhu@gmail.com>.
// All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// package support
// Author: ZHU HAIHUA
// Since: 2016-05-02 21:22
package support

import (
	"os"
	"syscall"
	"time"
)

func init() {
	_support = &supportLinux{}
}

type supportLinux struct{}

func (t *supportLinux) StatTimes(filepath string) (atime, ctime, mtime time.Time, err error) {
	fi, err := os.Stat(filepath)
	if err != nil {
		return
	}
	mtime = fi.ModTime()
	stat := fi.Sys().(*syscall.Stat_t)
	atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	ctime = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	return
}
