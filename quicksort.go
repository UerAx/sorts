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
		arr[r], arr[l] = arr[l], arr[r]
		for l < r && arr[l] < key {
			l++
		}
		arr[r], arr[l] = arr[l], arr[r]
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
	quickSort2Way(arr, 0, len(arr) - 1)
}

func quickSort2Way(arr []int, left, right int) {
	if left >= right {
		return
	}

	key := key(arr, left, right)
	l, r := left+1, right

	for {
		for l <= right && arr[l] < key {
			l++
		}
		for r > left && arr[r] > key {
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		} else {
			break
		}
	}
	arr[r], arr[left] = arr[left], arr[r]
	quickSort2Way(arr, left, r - 1)
	quickSort2Way(arr, r + 1, right)

}

func QuickSort3Way(arr []int) {

}