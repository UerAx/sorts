/*
 * @Author: ww
 * @Date: 2022-05-15 04:03:34
 * @Description:
 * @FilePath: \sorts\mergesort.go
 */
package sorts

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, left, right int) {
	if right <= left {
		return
	}
	mid := (right + left) >> 1
	if right-left > 1 {
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
	}

	l1, l2 := left, mid+1
	tmp := make([]int, 0, right-left+1)
	for l1 <= mid || l2 <= right {
		if l1 > mid {
			tmp = append(tmp, arr[l2:right+1]...)
			break
		}
		if l2 > right {
			tmp = append(tmp, arr[l1:mid+1]...)
			break
		}

		if arr[l1] > arr[l2] {
			tmp = append(tmp, arr[l2])
			l2++
		} else {
			tmp = append(tmp, arr[l1])
			l1++
		}
	}

	for i, v := range tmp {
		arr[i+left] = v
	}
}
