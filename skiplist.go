package main

import (
	"fmt"
	"math/rand"
)

const (
	SKIPLIST_MAXLEVEL = 32
	SKIPLIST_P        = 0.25
)

type SkipListLevel struct {
	Forward *SkipListNode
}

type SkipListNode struct {
	Key   int
	Value interface{}
	Level []SkipListLevel
}

func NewNode(level int, key int, value interface{}) *SkipListNode {
	return &SkipListNode{
		Key:   key,
		Value: value,
		Level: make([]SkipListLevel, level),
	}
}

type SkipList struct {
	Header *SkipListNode
	Level  int
}

func NewSkipList() *SkipList {
	return &SkipList{
		Header: NewNode(SKIPLIST_MAXLEVEL, 0, nil),
		Level:  1,
	}
}

func (this *SkipList) RandomLevel() int {
	level := 1
	for float64(rand.Int()&0xFFFF) < SKIPLIST_P*0xFFFF {
		level += 1
	}
	if level > SKIPLIST_MAXLEVEL {
		return SKIPLIST_MAXLEVEL
	}
	return level
}

func (this *SkipList) Insert(key int) {

	update := make([]*SkipListNode, SKIPLIST_MAXLEVEL)
	x := this.Header

	for i := this.Level - 1; i >= 0; i-- {
		for x.Level[i].Forward != nil &&
			x.Level[i].Forward.Key < key {
			x = x.Level[i].Forward
		}
		update[i] = x
	}

	level := this.RandomLevel()
	if level > this.Level {
		for i := this.Level; i < level; i++ {
			update[i] = this.Header
		}
		this.Level = level
	}

	x = NewNode(level, key, nil)

	for i := 0; i < level; i++ {
		x.Level[i].Forward = update[i].Level[i].Forward
		update[i].Level[i].Forward = x
	}
}

func (this *SkipList) Search(key int) *SkipListNode {

	x := this.Header
	for i := this.Level - 1; i >= 0; i-- {
		for x.Level[i].Forward != nil {
			if x.Level[i].Forward.Key == key {
				return x.Level[i].Forward
			}
			if x.Level[i].Forward.Key < key {
				x = x.Level[i].Forward
			} else {
				break
			}
		}
	}
	return nil
}

func (this *SkipList) Remove(key int) {

	update := make([]*SkipListNode, SKIPLIST_MAXLEVEL)
	x := this.Header
	for i := this.Level - 1; i >= 0; i-- {

		for x.Level[i].Forward != nil &&
			x.Level[i].Forward.Key < key {
			x = x.Level[i].Forward
		}

		update[i] = x
	}

	x = x.Level[0].Forward
	if x != nil && key == x.Key {

		for i := 0; i < this.Level; i++ {
			if update[i].Level[i].Forward == x {
				update[i].Level[i].Forward = x.Level[i].Forward
			}
		}

		for this.Level > 1 && this.Header.Level[this.Level-1].Forward == nil {
			this.Level--
		}
	}
}

func (this *SkipList) Range(f func(key int, value interface{})) {
	x := this.Header.Level[0].Forward
	for x != nil {
		f(x.Key, x.Value)
		x = x.Level[0].Forward
	}
}

func (this *SkipList) Print() {
	for i := SKIPLIST_MAXLEVEL - 1; i >= 0; i-- {
		fmt.Println("level:", i)
		x := this.Header.Level[i].Forward
		for x != nil {
			fmt.Printf("%d ", x.Key)
			x = x.Level[i].Forward
		}
		fmt.Println("\n-----------------------------------------------")
	}
	fmt.Println("MaxLevel:", this.Level)
}
