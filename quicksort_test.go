/*
 * @Author: UerAx
 * @Date: 2022-06-21 01:12:31
 * @FilePath: /sorts/quicksort_test.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package sorts

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{2, 5, 5, 3, 1}
	QuickSort(arr)
	fmt.Println(arr)
}

func TestQuickSort2Way(t *testing.T) {
	arr := []int{5,1,1,2,0,0}
	QuickSort2Way(arr)
	fmt.Println(arr)
}