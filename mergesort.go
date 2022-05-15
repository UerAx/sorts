/*
 * @Author: ww
 * @Date: 2022-05-15 04:03:34
 * @Description:
 * @FilePath: \sorts\mergesort.go
 */
package sorts

type MergeSort struct {
	arr  []int
	size int
}

func (this *MergeSort) Sort(arr []int) {
	this.add(arr)
	this.mergeSort(0, this.size-1)
}

func (this *MergeSort) add(arr []int) {
	this.arr = arr
	this.size = len(arr)
}

func (this *MergeSort) merge(left, right, mid int) {
	if right - left > 1 {
		this.merge(left, mid, (left+mid)>>1)
		this.merge(mid+1, right, (mid+right+1)>>1)
	}
	
	l1, l2 := left, mid+1
	tmp := make([]int, 0, right-left+1)
	for l1 <= mid || l2 <= right {
		if l1 > mid {
			tmp = append(tmp, this.arr[l2:right+1]...)
			break
		}
		if l2 > right {
			tmp = append(tmp, this.arr[l1:mid+1]...)
			break
		}

		if this.arr[l1] > this.arr[l2] {
			tmp = append(tmp, this.arr[l2])
			l2++
		} else {
			tmp = append(tmp, this.arr[l1])
			l1++
		}
	}

	for i, v := range tmp {
		this.arr[i+left] = v
	}

}

func (this *MergeSort) mergeSort(left, right int) {
	if right <= left{
		return
	}
	this.merge(left, right, (left+right)>>1)
}
