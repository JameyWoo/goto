/**
 * @Author: JameyWoo
 * @Email: 2622075127wjh@gmail.com
 * @Date: 2020/12/15
 * @Desc: tips
 * @Copyright (c) 2020, JameyWoo. All rights reserved.
 */

package tips

import (
	"fmt"
	"testing"
)

type SortElem struct {
	a, b int
}

func (e *SortElem) Less(elem SortElem) bool {
	return (*e).a < (elem).a
}

func TestSortStruct(t *testing.T) {
	x := SortElem{a: 1, b: 3}
	y := SortElem{a: 2, b: 2}
	fmt.Println(x == y)
	fmt.Println(x.Less(y))
	fmt.Println(y.Less(x))
}
