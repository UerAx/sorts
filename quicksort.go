package sorts

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {

	if left >= right {
		return
	}

	key := key(arr, left, right)

	l, r := left, right

	for l < r {
		for l < r && arr[r] >= key {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= key {
			l++
		}
		arr[r] = arr[l]
	}

	arr[l] = key
	quickSort(arr, left, l-1)
	quickSort(arr, l+1, right)

}

func key(arr []int, left, right int) int {
	mid := (left + right) >> 1
	arr[left], arr[mid] = arr[mid], arr[left]
	return arr[left]
}

func QuickSort2Way(arr []int) {

}

func QuickSort3Way(arr []int) {

}