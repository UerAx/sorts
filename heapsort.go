/*
 * @Author: UerAx
 * @Date: 2022-05-14 23:32:56
 * @FilePath: \sorts\heapsort.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package sorts

func HeapSort(arr []int) {
	heapSort(arr)
}

func heapSort(arr []int) {
	adjustUp(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		adjust(arr, 0)
	}
}

func adjustUp(arr []int) {
	for i := (len(arr) >> 1) - 1; i >= 0; i-- {
		adjust(arr, i)
	}
}

func adjust(arr []int, idx int) {
	size := len(arr)
	if idx > size || idx<<1 >= size-1 {
		return
	}

	son := (idx << 1) + 1
	if son+1 < size && arr[son] < arr[son+1] {
		son++
	}

	if arr[son] > arr[idx] {
		arr[son], arr[idx] = arr[idx], arr[son]
		adjust(arr, size)
	}
}
