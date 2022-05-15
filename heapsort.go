/*
 * @Author: UerAx
 * @Date: 2022-05-14 23:32:56
 * @FilePath: /sorts/heapsort.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package sorts

type HeapSort struct {
	arr  []int
	size int
}

func (p *HeapSort) Sort(arr []int) {
	p.add(arr)
	p.adjustUp()
	for i := p.size - 1; i > 0; i-- {
		p.arr[i], p.arr[0] = p.arr[0], p.arr[i]
		p.adjust(0, i)
	}
}

func (p *HeapSort) add(arr []int) {
	p.arr = arr
	p.size = len(arr)
}

func (p *HeapSort) adjustUp() {
	for i := (p.size >> 1) - 1; i >= 0; i-- {
		p.adjust(i, p.size)
	}
}

func (p *HeapSort) adjust(idx, size int) {
	if idx > size || idx<<1 >= size-1 {
		return
	}

	son := (idx << 1) + 1
	if son+1 < size && p.arr[son] < p.arr[son+1] {
		son++
	}

	if p.arr[son] > p.arr[idx] {
		p.arr[son], p.arr[idx] = p.arr[idx], p.arr[son]
		p.adjust(son, size)
	}
}
