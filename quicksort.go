package sorts

type QuickSort struct {
	arr  []int
	size int
}

func (this *QuickSort) Sort(arr []int) {
	this.add(arr)
	this.quickSort(0, this.size - 1)
}

func (this *QuickSort) add(arr []int) {
	this.arr = arr
	this.size = len(arr)
}

func (this *QuickSort) quickSort(left, right int) {
	if right <= left {
		return
	}
	
	key := this.getKey(left, right)

	l,r := left,right

	for(l < r) {
		for l < r && this.arr[r] >= key {r--}
		this.arr[l] = this.arr[r]
		for l < r && this.arr[l] <= key {l++}
		this.arr[r] = this.arr[l]
	}

	this.arr[l] = key
	this.quickSort(left, l-1)
	this.quickSort(l+1, right)
	

}

func (this *QuickSort) getKey(left, right int) int {
	mid := (left+right)>>1
	this.arr[left], this.arr[mid] = this.arr[mid], this.arr[left]
	return this.arr[left]
}

