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
		mid = (low + high) / 2

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

	tmplist := append([]SortNode{}, this.list[index:]...)

	this.list = append(this.list[:index], SortNode{
		key:   key,
		value: lsize,
	})

	this.list = append(this.list, tmplist...)

}

func (this *SortList) Load(key int) *SortNode {
	var (
		lsize = len(this.list)
		low   = 0
		high  = lsize - 1
	)
	for low <= high {
		mid := (low + high) / 2
		if mid >= lsize {
			return nil
		}
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

func (this *SortList) Range(f func(key int, value interface{})) {
	for i := 0; i < len(this.list); i++ {
		node := this.list[i]
		f(node.key, node.value)
	}
}
