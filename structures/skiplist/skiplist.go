package skiplist

import (
	"github.com/zaoshi-studio/slime-go/types"
	"math/rand"
)

const (
	skipListMaxLevel = 32
	skipListP        = 0.25
)

type SkipList[K types.Comparable, V any] struct {
	header, tail *SkipListNode[K, V]
	length       int
	level        int
	comparator   types.Comparator[K]
}

func NewSkipList[K types.Comparable, V any](comparator types.Comparator[K]) *SkipList[K, V] {
	skipList := &SkipList[K, V]{
		level:      1,
		length:     0,
		comparator: comparator,
	}

	var (
		defaultK K
		defaultV V
	)

	skipList.header = NewSkipListNode[K, V](skipListMaxLevel, defaultK, defaultV)
	return skipList
}

func (sl *SkipList[K, V]) randomLevel() int {
	var level = 1
	for float64(rand.Int()&0xFFFF) < (skipListP * 0xFFFF) {
		level += 1
	}
	if level < skipListMaxLevel {
		return level
	}
	return skipListMaxLevel
}

// getKeyRank
//
//	 @receiver *SkipList[K, V]
//	 @param key
//	 @return int
//	 @Description:
//		get node rank from header
//		node rank start from 0
func (sl *SkipList[K, V]) getKeyRank(key K) int {
	x := sl.header
	rank := 0

	for level := sl.level - 1; level >= 0; level-- {
		for next := x.Forward(level); next != nil && (sl.comparator.Compare(next.key, key) || next.key.EQ(key)); next = x.Forward(level) {
			rank += x.Span(level)
			x = next
		}

		if x != sl.header && x.key.EQ(key) {
			return rank
		}
	}
	return 0
}

func (sl *SkipList[K, V]) getNodeByRank(rank int) *SkipListNode[K, V] {
	x := sl.header
	spanned := 0
	for level := sl.level - 1; level >= 0; level-- {
		for next := x.Forward(level); next != nil && spanned+x.Span(level) <= rank; next = x.Forward(level) {
			spanned += x.Span(level)
			x = next
		}
		if spanned == rank {
			return x
		}
	}
	return nil
}

func (sl *SkipList[K, V]) Insert(key K, value V) *SkipListNode[K, V] {
	update := make([]*SkipListNode[K, V], skipListMaxLevel)
	rank := make([]int, skipListMaxLevel)

	x := sl.header

	for level := sl.level - 1; level >= 0; level-- {
		if level == sl.level-1 {
			rank[level] = 0
		} else {
			rank[level] = rank[level+1]
		}

		for x.Forward(level) != nil && sl.comparator.Compare(x.Forward(level).key, key) {
			rank[level] += x.Span(level)
			x = x.Forward(level)
		}
		update[level] = x
	}

	randLevel := sl.randomLevel()

	if randLevel > sl.level {
		for level := sl.level; level < randLevel; level++ {
			rank[level] = 0
			update[level] = sl.header
			update[level].SetSpan(level, sl.length)
		}
		sl.level = randLevel
	}

	x = NewSkipListNode[K](randLevel, key, value)
	for level := 0; level < randLevel; level++ {

		x.SetForward(level, update[level].Forward(level))
		update[level].SetForward(level, x)

		x.level[level].span = update[level].Span(level) - (rank[0] - rank[level])
		update[level].SetSpan(level, rank[0]-rank[level]+1)
	}

	for i := randLevel; i < sl.level; i++ {
		update[i].level[i].span++
	}

	if update[0] == sl.header {
		x.backward = nil
	} else {
		x.backward = update[0]
	}

	if x.Forward(0) != nil {
		x.Forward(0).backward = x
	} else {
		sl.tail = x
	}

	sl.length++

	return x
}

func (sl *SkipList[K, V]) deleteNode(x *SkipListNode[K, V], update []*SkipListNode[K, V]) {
	for level := 0; level < sl.level; level++ {
		if update[level].Forward(level) == x {
			update[level].SetForward(level, x.Forward(level))
			update[level].SetSpan(level, update[level].Span(level)+x.Span(level)-1)

		} else {
			update[level].SetSpan(level, update[level].Span(level)-1)
		}
	}

	if x.Forward(0) != nil {
		x.Forward(0).backward = x.backward
	} else {
		sl.tail = x.backward
	}

	for sl.level > 1 && sl.header.Forward(sl.level-1) == nil {
		sl.level--
	}
	sl.length--
}

func (sl *SkipList[K, V]) Delete(key K) {
	update := make([]*SkipListNode[K, V], skipListMaxLevel)

	x := sl.header

	for level := sl.level - 1; level >= 0; level-- {
		for x.Forward(level) != nil && sl.comparator.Compare(x.Forward(level).key, key) {
			x = x.Forward(level)
		}

		update[level] = x
	}

	x = x.Forward(0)

	if x != nil && x.key.EQ(key) {
		sl.deleteNode(x, update)
	}
}

func (sl *SkipList[K, V]) Update(oldKey, newKey K, value V) {
	sl.Delete(oldKey)
	sl.Insert(newKey, value)
}

// GetKeyRank
//
//	 @receiver *SkipList[K, V]
//	 @param key
//	 @return int
//	 @Description:
//		get node rank from first node
//		the first node rank is 0
func (sl *SkipList[K, V]) GetKeyRank(key K) int {
	rank := sl.getKeyRank(key)
	return rank - 1
}

func (sl *SkipList[K, V]) GetNodeByKey(key K) *SkipListNode[K, V] {
	x := sl.header
	for level := sl.level - 1; level >= 0; level-- {
		for x.Forward(level) != nil && sl.comparator.Compare(x.Forward(level).key, key) {
			x = x.Forward(level)
		}
	}
	if x.Forward(0) != nil && x.Forward(0).key.EQ(key) {
		return x.Forward(0)
	}
	return nil
}

func (sl *SkipList[K, V]) GetNodeByRank(rank int) *SkipListNode[K, V] {
	return sl.getNodeByRank(rank + 1)
}

func (sl *SkipList[K, V]) GetNodeByRange(start, end int) []*SkipListNode[K, V] {
	if start < 0 {
		start = sl.length + start
	}

	if end < 0 {
		end = sl.length + end
	}

	if start < 0 {
		start = 0
	}

	if end >= sl.length {
		end = sl.length - 1
	}

	x := sl.getNodeByRank(start + 1)

	if x == nil {
		return nil
	}

	var ret []*SkipListNode[K, V]
	for idx := start; idx <= end && x != nil; idx++ {
		ret = append(ret, x)
		x = x.Forward(0)
	}
	return ret
}

type SkipListLevel[K types.Comparable, V any] struct {
	forward *SkipListNode[K, V]
	span    int
}

type SkipListNode[K types.Comparable, V any] struct {
	key      K
	value    V
	level    []*SkipListLevel[K, V]
	backward *SkipListNode[K, V]
}

func NewSkipListNode[K types.Comparable, V any](level int, key K, value V) *SkipListNode[K, V] {
	skipListNode := &SkipListNode[K, V]{
		key:   key,
		value: value,
		level: make([]*SkipListLevel[K, V], level),
	}

	for i := 0; i < level; i++ {
		skipListNode.level[i] = &SkipListLevel[K, V]{}
	}
	return skipListNode
}

func (node *SkipListNode[K, V]) Forward(level int) *SkipListNode[K, V] {
	return node.level[level].forward
}

func (node *SkipListNode[K, V]) SetForward(level int, forward *SkipListNode[K, V]) {
	node.level[level].forward = forward
}

func (node *SkipListNode[K, V]) Span(level int) int {
	return node.level[level].span
}

func (node *SkipListNode[K, V]) SetSpan(level, span int) {
	node.level[level].span = span
}
