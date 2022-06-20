/*
 * @Author: UerAx
 * @Date: 2022-06-21 01:12:31
 * @FilePath: \sorts\quicksort_test.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package sorts

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{5, 5, 5, 5, 1}
	QuickSort(arr)
	fmt.Println(arr)
}
