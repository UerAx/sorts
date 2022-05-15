/*
 * @Author: UerAx
 * @Date: 2022-05-14 23:32:56
 * @FilePath: \sorts\sort_test.go
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
	s2 := new(MergeSort)
	arr2 := []int{1, 3, 5, 7, 9, 2, 4, 8, 6, 4}
	s2.Sort(arr2)
	t.Logf("%v", s2)

}
