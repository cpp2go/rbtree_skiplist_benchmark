package main

import (
	"math/rand"
	"testing"
	"time"
)

var (
	rbmap    *RbTree
	skiplist *SkipList
	gomap    map[uint32]*int
	sortlist SortList
)

func init() {

	rand.Seed(time.Now().Unix())

	rbmap = NewRbTree()
	for i := 0; i < 1000; i++ {
		rbmap.Store(ngx_rbtree_key_t(i), &i)
	}

	skiplist = NewSkipList()
	for i := 0; i < 1000; i++ {
		skiplist.Insert(i)
	}

	gomap = make(map[uint32]*int)
	for i := 0; i < 1000; i++ {
		gomap[uint32(i)] = &i
	}

	for i := 0; i < 1000; i++ {
		sortlist.Insert(i)
	}
}

func Benchmark_RbTreeNew(b *testing.B) {

	for i := 0; i < b.N; i++ {

		rbmap := NewRbTree()
		_ = &rbmap
	}
}

func Benchmark_SkipListNew(b *testing.B) {

	for i := 0; i < b.N; i++ {

		skiplist := NewSkipList()
		_ = &skiplist
	}
}

func Benchmark_MapNew(b *testing.B) {

	for i := 0; i < b.N; i++ {

		gomap := make(map[uint32]*int)
		_ = &gomap
	}
}

func Benchmark_RbTreeInsert(b *testing.B) {

	rbmap := NewRbTree()

	for i := 0; i < b.N; i++ {

		rbmap.Store(ngx_rbtree_key_t(i), &i)
	}
}

func Benchmark_SkipListInsert(b *testing.B) {

	skiplist := NewSkipList()

	for i := 0; i < b.N; i++ {

		skiplist.Insert(i)
	}
}

func Benchmark_MapInsert(b *testing.B) {

	gomap := make(map[uint32]*int)

	for i := 0; i < b.N; i++ {

		gomap[uint32(i)] = &i
	}
}

func Benchmark_SortListInsert(b *testing.B) {

	sortlist := SortList{}

	for i := 0; i < b.N; i++ {

		sortlist.Insert(i)
	}
}

func Benchmark_RbTreeLoad(b *testing.B) {

	for i := 0; i < b.N; i++ {

		_, ok := rbmap.Load(ngx_rbtree_key_t(i))
		if !ok {

		}
	}
}

func Benchmark_SkipListLoad(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = skiplist.Search(i)
	}
}

func Benchmark_MapLoad(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, ok := gomap[uint32(i)]
		if !ok {

		}
	}
}

func Benchmark_SortListLoad(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = sortlist.Load(i)
	}
}

func Benchmark_RbTreeRange(b *testing.B) {

	for i := 0; i < b.N; i++ {

		rbmap.Range(func(key ngx_rbtree_key_t, data interface{}) {
			//fmt.Printf("key = %d value = %v\n", key, data)
			_ = key
		})

	}
}

func Benchmark_SkipListRange(b *testing.B) {

	for i := 0; i < b.N; i++ {
		skiplist.Range(func(key int, data interface{}) {
			//fmt.Printf("key = %d value = %v\n", key, data)
			_ = key
		})
	}
}

func Benchmark_MapRange(b *testing.B) {

	for i := 0; i < b.N; i++ {

		for k, _ := range gomap {
			//fmt.Printf("key = %d value = %v\n", k, v)
			_ = k
		}
	}
}

func Benchmark_SortListRange(b *testing.B) {

	for i := 0; i < b.N; i++ {
		/*
			sortlist.Range(func(key int, value interface{}) {
				//fmt.Printf("key = %d value = %v\n", key, data)
				_ = key
			})
		*/

		for i := 0; i < len(sortlist.list); i++ {
			node := sortlist.list[i]
			_ = node.key
		}

	}
}
