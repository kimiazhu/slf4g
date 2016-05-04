// Copyright 2011 ZHU HAIHUA <kimiazhu@gmail.com>.
// All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

// Description: slf4g
// Author: ZHU HAIHUA
// Since: 2016-05-02 22:20
package slf4g

type Configuration struct {

	// Auto detect the change of the config file
	Scan bool

	// Scan every period interval
	ScanPeriod int

	// Whether print the debug messages of slf4g itself
	Debug bool
}
