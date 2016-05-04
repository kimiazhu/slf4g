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
	_support = &supportDarwin{}
}

type supportDarwin struct{}

func (t *supportDarwin) StatTimes(filepath string) (atime, ctime, mtime time.Time, err error) {
	fi, err := os.Stat(filepath)
	if err != nil {
		return
	}
	mtime = fi.ModTime()
	stat := fi.Sys().(*syscall.Stat_t)
	atime = time.Unix(int64(stat.Atimespec.Sec), int64(stat.Atimespec.Nsec))
	ctime = time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec))
	return
}
