// Description: slf4g
// Author: ZHU HAIHUA
// Since: 2016-05-02 22:35
package slf4g

import (
)

type IAppender interface {

	DoAppend(record *Record)

	//Roll()

	//NeedRoll()
}

