/*
 * @Author: UerAx
 * @Date: 2022-05-14 23:32:56
 * @FilePath: /sorts/heapsort_test.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package sorts

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{5, 5, 5, 5, 1}
	HeapSort(arr)
	fmt.Println(arr)
}
