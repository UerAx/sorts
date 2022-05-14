package sorts


type HeapSort struct {
	arr []int
}

func (this *HeapSort) Sort(arr []int) {
	this.add(arr)
	this.adjustUp()
	for i := len(this.arr) - 1; i > 0; i-- {
		this.arr[0], this.arr[i] = this.arr[i], this.arr[0]
		this.adjustDown(0, i)
	}
}

func (this *HeapSort) add(arr []int) {
	this.arr = arr;
}

func (this *HeapSort) adjustUp() {
	for i := len(this.arr) - 1; i > 0; i--{
		if this.arr[i] > this.arr[(i-1)/2] {
			this.arr[i], this.arr[(i-1)/2] = this.arr[(i-1)/2], this.arr[i]
			this.adjustDown(i, len(this.arr) - 1)
		}
	}
}

func (this *HeapSort) adjustDown(index, size int) {
	if index >= size || index * 2 + 1 >= len(this.arr) {
		return
	}
	son := index * 2 + 1
	if (son > size) {
		return
	}
	if son + 1 < size && son + 1 < len(this.arr) && this.arr[son] < this.arr[son+1] {
		son++
	}
	if this.arr[son] > this.arr[index] {
		this.arr[son], this.arr[index] = this.arr[index], this.arr[son]
		this.adjustDown(son, size)
	}
}