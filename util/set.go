package util

import (
	"sync"
)

// 数组转集合

// Set 集合
type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

// New 创建一个集合
func New() *Set {
	return &Set{m: make(map[interface{}]bool)}
}

// Clear 集合清空
func (I *Set) Clear() {
	I.Lock()
	defer I.Unlock()
	I.m = make(map[interface{}]bool)
}

// List 集合转数组
func (I *Set) List() []interface{} {
	I.RLock()
	defer I.RUnlock()
	var list []interface{}
	for item := range I.m {
		list = append(list, item)
	}
	return list
}

// Len 集合长度
func (I *Set) Len() int {
	return len(I.m)
}

// Add 集合新增元素
func (I *Set) Add(item interface{}) {
	I.Lock()
	defer I.Unlock()
	I.m[item] = true
}

// Remove 集合去除元素
func (I *Set) Remove(item interface{}) {
	I.Lock()
	defer I.Unlock()
	delete(I.m, item)
}

// Has 集合存在
func (I *Set) Has(item interface{}) bool {
	I.RLock()
	defer I.RUnlock()
	_, ok := I.m[item]
	return ok
}

func (I *Set) union(m []interface{}) {
	I.Lock()
	defer I.Unlock()
	for _, item := range m {
		I.m[item] = true
	}
}

// Union 返回一个线程安全的并集对象
func (I *Set) Union(H *Set) *Set {
	// 先复制自己的数组
	retSet := New()
	retSet.union(I.List())
	retSet.union(H.List())
	return retSet
}

// mix 返回一个交集数组
func (I *Set) mix(m []interface{}) []interface{} {
	I.RLock()
	defer I.RUnlock()
	var retList []interface{}
	for _, item := range m {
		if _, ok := I.m[item]; ok {
			retList = append(retList, item)
		}
	}
	return retList
}

// Mix 返回一个线程安全的交集对象
func (I *Set) Mix(H *Set) *Set {
	retSet := New()
	// 获取h的列表到I中获取交集存放到新集合中
	retSet.union(I.mix(H.List()))
	return retSet
}
func (I *Set) diff(m []interface{}) []interface{} {
	I.RLock()
	defer I.RUnlock()
	var retList []interface{}
	for _, item := range m {
		if _, ok := I.m[item]; !ok {
			retList = append(retList, item)
		}
	}
	return retList
}

// Diff 差集
func (I *Set) Diff(H *Set) *Set {
	retSet := New()
	// 差集
	retSet.union(I.diff(H.List()))
	return retSet
}
