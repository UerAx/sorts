/*
 * @Author: UerAx
 * @Date: 2022-05-14 23:32:56
 * @FilePath: /sorts/sort_test.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package sorts

import "testing"

func TestSort(t *testing.T) {
	s := new(HeapSort)
	arr := []int{5, 2, 6}
	s.Sort(arr)
	t.Log(s)
	s1 := new(QuickSort)
	arr1 := []int{}
	s1.Sort(arr1)
	t.Logf("%v", s1)
}
