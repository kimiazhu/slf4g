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

	Decide(level Level) FilterReply
}