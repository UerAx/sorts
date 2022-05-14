/*
 * @Author: UerAx
 * @Date: 2022-05-14 23:32:56
 * @FilePath: \sorts\heapsort.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package sorts

type HeapSort struct {
	arr  []int
	size int
}

func (this *HeapSort) Sort(arr []int) {
	this.add(arr)
	this.adjustUp()
	for i := this.size - 1; i > 0; i-- {
		this.arr[i], this.arr[0] = this.arr[0], this.arr[i]
		this.adjust(0, i)
	}
}

func (this *HeapSort) add(arr []int) {
	this.arr = arr
	this.size = len(arr)
}

func (this *HeapSort) adjustUp() {
	for i := (this.size >> 1) - 1; i >= 0; i-- {
		this.adjust(i, this.size)
	}
}

func (this *HeapSort) adjust(idx, size int) {
	if idx > size || idx<<1 >= size-1 {
		return
	}

	son := (idx << 1) + 1
	if son+1 < size && this.arr[son] < this.arr[son+1] {
		son++
	}

	if this.arr[son] > this.arr[idx] {
		this.arr[son], this.arr[idx] = this.arr[idx], this.arr[son]
		this.adjust(son, size)
	}
}
