/*
 * @Author: UerAx
 * @Date: 2022-05-14 23:32:56
 * @FilePath: /sorts/heapsort.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package sorts

func HeapSort(arr []int) {
	heapSort(arr)
}

func heapSort(arr []int) {
	adjustAll(arr)
	size := len(arr)
	for i := 1; i < size; i++ {
		arr[0], arr[size - i] = arr[size - i], arr[0]
		adjust(arr[:size-i], 0)
	}
}

func adjustAll(arr []int) {
	for i := len(arr) >> 1; i >= 0; i-- {
		adjust(arr, i)
	}
}

func adjust(arr []int, index int) {
	if index < 0 || index >= len(arr) {
		return
	}
	son := (index << 1) + 1
	if son < len(arr) {
		if son + 1 < len(arr) && arr[son] < arr[son+1] {
			son++
		}
		if arr[son] > arr[index] {
			arr[son], arr[index] = arr[index], arr[son]
			adjust(arr, son)
		}
	}
}

