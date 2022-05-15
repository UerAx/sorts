/*
 * @Author: ww
 * @Date: 2022-05-15 04:03:34
 * @Description:
 * @FilePath: /sorts/mergesort.go
 */
package sorts

type MergeSort struct {
	arr  []int
	size int
}

func (p *MergeSort) Sort(arr []int) {
	p.add(arr)
	p.mergeSort(0, p.size-1)
}

func (p *MergeSort) add(arr []int) {
	p.arr = arr
	p.size = len(arr)
}

func (p *MergeSort) merge(left, right, mid int) {
	if right-left > 1 {
		p.merge(left, mid, (left+mid)>>1)
		p.merge(mid+1, right, (mid+right+1)>>1)
	}

	l1, l2 := left, mid+1
	tmp := make([]int, 0, right-left+1)
	for l1 <= mid || l2 <= right {
		if l1 > mid {
			tmp = append(tmp, p.arr[l2:right+1]...)
			break
		}
		if l2 > right {
			tmp = append(tmp, p.arr[l1:mid+1]...)
			break
		}

		if p.arr[l1] > p.arr[l2] {
			tmp = append(tmp, p.arr[l2])
			l2++
		} else {
			tmp = append(tmp, p.arr[l1])
			l1++
		}
	}

	for i, v := range tmp {
		p.arr[i+left] = v
	}

}

func (p *MergeSort) mergeSort(left, right int) {
	if right <= left {
		return
	}
	p.merge(left, right, (left+right)>>1)
}
