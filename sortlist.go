package main

type SortNode struct {
	key   int
	value interface{}
}

type SortList struct {
	list []SortNode
}

func (this *SortList) Insert(key int) {
	/*
		this.list = append(this.list, SortNode{
			key:   key,
			value: len(this.list),
		})

		j := len(this.list) - 1

		for j > 0 && this.list[j].key < this.list[j-1].key {

			this.list[j], this.list[j-1] = this.list[j-1], this.list[j]

			j--
		}

		if j > 0 && this.list[j].key == this.list[j-1].key {
			this.list = append(this.list[:j-1], this.list[j:]...)
		}
	*/

	var (
		lsize = len(this.list)
		low   = 0
		high  = lsize - 1
		mid   int
	)

	if lsize == 0 || this.list[high].key < key {

		this.list = append(this.list, SortNode{
			key:   key,
			value: lsize,
		})

		return
	}

	for low <= high {
		mid = (low + high) >> 1

		node := &this.list[mid]

		if node.key == key {

			this.list[mid] = SortNode{
				key:   key,
				value: lsize,
			}

			return
		}

		if node.key > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	index := mid

	if this.list[mid].key < key {
		index = mid + 1
	}

	/*
		tmplist := append([]SortNode{}, this.list[index:]...)
		this.list = append(this.list[:index], SortNode{
			key:   key,
			value: lsize,
		})
		this.list = append(this.list, tmplist...)
	*/

	/*
		this.list = append(this.list, SortNode{})
		copy(this.list[index+1:], this.list[index:])
		this.list[index] = SortNode{
			key:   key,
			value: lsize,
		}
	*/

	this.list = append(this.list, SortNode{
		key:   key,
		value: lsize,
	})
	for i := lsize; i > index; i-- {
		this.list[i], this.list[i-1] = this.list[i-1], this.list[i]
	}
}

func (this *SortList) Load(key int) *SortNode {
	lsize := len(this.list)
	if lsize == 0 {
		return nil
	}
	var (
		low  = 0
		high = lsize - 1
		mid  int
	)
	for low <= high {
		mid = (low + high) >> 1
		node := &this.list[mid]
		if node.key == key {
			return node
		}
		if node.key > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return nil
}

func (this *SortList) Delete(key int) bool {
	lsize := len(this.list)
	if lsize == 0 {
		return false
	}
	var (
		low  = 0
		high = lsize - 1
		mid  int
	)
	for low <= high {
		mid = (low + high) >> 1
		node := &this.list[mid]
		if node.key == key {
			this.list = append(this.list[:mid], this.list[mid+1:]...)
			return true
		}
		if node.key > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func (this *SortList) Range(f func(key int, value interface{})) {
	for i := 0; i < len(this.list); i++ {
		node := this.list[i]
		f(node.key, node.value)
	}
}
