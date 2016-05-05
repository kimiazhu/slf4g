// Copyright 2011 ZHU HAIHUA <kimiazhu@gmail.com>.
// All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// Description: slf4g
// Author: ZHU HAIHUA
// Since: 2016-05-02 22:49
package slf4g

type FilterReply int

const (
	DENY FilterReply = iota
	NEUTRAL
	ACCEPT
)

type IFilter interface {

	// Decide whether the log should be write or not
	Decide(level Level) FilterReply
}

type Filter struct {
	Lvl Level
}

type LevelFilter struct {
	*Filter
	onMatchAccept bool
	onMismatchAccept bool
}

func (f *LevelFilter) Decide(lvl Level) FilterReply {
	if (lvl == f.Lvl && f.onMatchAccept) || (lvl != f.Lvl && f.onMatchAccept) {
		return ACCEPT
	}
	return DENY
}

type ThresholdFilter struct {
	*Filter
}

func (f *ThresholdFilter) Decide(lvl Level) FilterReply {
	if lvl >= f.Lvl {
		return NEUTRAL
	} else {
		return DENY
	}
}