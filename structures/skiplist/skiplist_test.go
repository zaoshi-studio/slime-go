package skiplist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zaoshi-studio/slime-go/types"
	"math/rand"
	"sort"
	"testing"
)

type testValue struct {
	value int
}

func (v1 *testValue) EQ(c types.Comparable) bool {
	v2, ok := c.(*testValue)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *testValue) NE(c types.Comparable) bool {
	v2, ok := c.(*testValue)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *testValue) LT(c types.Comparable) bool {
	v2, ok := c.(*testValue)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *testValue) LE(c types.Comparable) bool {
	v2, ok := c.(*testValue)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *testValue) GT(c types.Comparable) bool {
	v2, ok := c.(*testValue)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *testValue) GE(c types.Comparable) bool {
	v2, ok := c.(*testValue)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

type testComparator[T types.Comparable] struct{}

func (cmp *testComparator[T]) Compare(t1 T, t2 T) bool {
	return t1.LT(t2)
}

func debugPrintSkipList[K types.Comparable, V any](t *testing.T, sl *SkipList[K, V]) {
	fmt.Println("debug print skip list")
	for level := sl.level - 1; level >= 0; level-- {
		fmt.Printf("Level %d: ", level)
		node := sl.header.Forward(level)
		for node != nil {
			fmt.Printf("%v (span: %d) -> ", node.key, node.level[level].span)
			node = node.Forward(level)
		}
		fmt.Println("nil")
	}
}

func debugPrintTestCases(t *testing.T, testCases []*testCase) {
	fmt.Println("debug print test cases")
	for _, v := range testCases {
		fmt.Printf("%v -> ", v.key)
	}
	fmt.Println("nil")
}

type testCase struct {
	value      float64
	key        *testValue
	updatedKey *testValue
}

const (
	n    = 10
	seed = 1718853930223332400
)

var (
	r = rand.New(rand.NewSource(seed))
)

func TestSkipList_Get(t *testing.T) {
	comparator := &testComparator[*testValue]{}

	sl := NewSkipList[*testValue, float64](comparator)

	node1 := NewSkipListNode[*testValue, float64](1, &testValue{value: 1}, 1)
	node2 := NewSkipListNode[*testValue, float64](2, &testValue{value: 2}, 2)
	node5 := NewSkipListNode[*testValue, float64](1, &testValue{value: 5}, 5)
	node6 := NewSkipListNode[*testValue, float64](3, &testValue{value: 6}, 6)
	node7 := NewSkipListNode[*testValue, float64](1, &testValue{value: 7}, 7)
	node8 := NewSkipListNode[*testValue, float64](2, &testValue{value: 8}, 8)
	node9 := NewSkipListNode[*testValue, float64](1, &testValue{value: 9}, 9)
	node10 := NewSkipListNode[*testValue, float64](4, &testValue{value: 10}, 10)

	node1.SetForward(0, node2)
	node1.SetSpan(0, 1)

	node2.SetForward(0, node5)
	node2.SetSpan(0, 1)
	node2.SetForward(1, node6)
	node2.SetSpan(1, 2)

	node5.SetForward(0, node6)
	node5.SetSpan(0, 1)

	node6.SetForward(0, node7)
	node6.SetSpan(0, 1)
	node6.SetForward(1, node8)
	node6.SetSpan(1, 2)
	node6.SetForward(2, node10)
	node6.SetSpan(2, 4)

	node7.SetForward(0, node8)
	node7.SetSpan(0, 1)

	node8.SetForward(0, node9)
	node8.SetSpan(0, 1)
	node8.SetForward(1, node10)
	node8.SetSpan(1, 2)

	node9.SetForward(0, node10)
	node9.SetSpan(0, 1)

	node10.SetForward(0, nil)
	node10.SetSpan(0, 0)
	node10.SetForward(1, nil)
	node10.SetSpan(1, 0)
	node10.SetForward(2, nil)
	node10.SetSpan(2, 0)
	node10.SetForward(3, nil)
	node10.SetSpan(3, 0)

	sl.header.SetForward(0, node1)
	sl.header.SetSpan(0, 1)
	sl.header.SetForward(1, node2)
	sl.header.SetSpan(1, 2)
	sl.header.SetForward(2, node6)
	sl.header.SetSpan(2, 4)
	sl.header.SetForward(3, node10)
	sl.header.SetSpan(3, 8)
	sl.level = 4
	sl.tail = node10

	debugPrintSkipList(t, sl)

	assert.Equal(t, node1.key, sl.GetNodeByKey(&testValue{value: 1}).key)
	assert.Equal(t, node2.key, sl.GetNodeByKey(&testValue{value: 2}).key)
	assert.Equal(t, node5.key, sl.GetNodeByKey(&testValue{value: 5}).key)
	assert.Equal(t, node6.key, sl.GetNodeByKey(&testValue{value: 6}).key)
	assert.Equal(t, node7.key, sl.GetNodeByKey(&testValue{value: 7}).key)
	assert.Equal(t, node8.key, sl.GetNodeByKey(&testValue{value: 8}).key)
	assert.Equal(t, node9.key, sl.GetNodeByKey(&testValue{value: 9}).key)
	assert.Equal(t, node10.key, sl.GetNodeByKey(&testValue{value: 10}).key)

	assert.Equal(t, float64(1), sl.GetNodeByKey(&testValue{value: 1}).value)
	assert.Equal(t, float64(2), sl.GetNodeByKey(&testValue{value: 2}).value)
	assert.Equal(t, float64(5), sl.GetNodeByKey(&testValue{value: 5}).value)
	assert.Equal(t, float64(6), sl.GetNodeByKey(&testValue{value: 6}).value)
	assert.Equal(t, float64(7), sl.GetNodeByKey(&testValue{value: 7}).value)
	assert.Equal(t, float64(8), sl.GetNodeByKey(&testValue{value: 8}).value)
	assert.Equal(t, float64(9), sl.GetNodeByKey(&testValue{value: 9}).value)
	assert.Equal(t, float64(10), sl.GetNodeByKey(&testValue{value: 10}).value)

	assert.Equal(t, 0, sl.GetKeyRank(&testValue{value: 1}))
	assert.Equal(t, 1, sl.GetKeyRank(&testValue{value: 2}))
	assert.Equal(t, 2, sl.GetKeyRank(&testValue{value: 5}))
	assert.Equal(t, 3, sl.GetKeyRank(&testValue{value: 6}))
	assert.Equal(t, 4, sl.GetKeyRank(&testValue{value: 7}))
	assert.Equal(t, 5, sl.GetKeyRank(&testValue{value: 8}))
	assert.Equal(t, 6, sl.GetKeyRank(&testValue{value: 9}))
	assert.Equal(t, 7, sl.GetKeyRank(&testValue{value: 10}))

}

func TestSkipList_Insert(t *testing.T) {
	comparator := &testComparator[*testValue]{}

	sl := NewSkipList[*testValue, float64](comparator)

	testCases := make([]*testCase, n)

	for i := 0; i < n; i++ {
		randValue := r.Intn(n) + i*n
		testCases[i] = &testCase{
			key: &testValue{
				value: randValue,
			},
			value: float64(i),
		}
	}

	sort.Slice(testCases, func(i, j int) bool {
		return testCases[i].key.LT(testCases[j].key)
	})

	for _, v := range testCases {
		sl.Insert(&testValue{
			value: v.key.value,
		}, v.value)
	}

	debugPrintTestCases(t, testCases)
	debugPrintSkipList(t, sl)

	for idx, v := range testCases {
		node := sl.GetNodeByKey(v.key)
		assert.NotNil(t, node)
		assert.Equal(t, v.key, node.key)
		assert.Equal(t, v.value, node.value)
		rank := sl.GetKeyRank(v.key)
		assert.Equal(t, idx, rank)
	}
}

func TestSkipList_Update(t *testing.T) {
	comparator := &testComparator[*testValue]{}

	sl := NewSkipList[*testValue, float64](comparator)

	testCases := make([]*testCase, n)

	for i := 0; i < n; i++ {
		randValue := r.Intn(n) + i*n
		testCases[i] = &testCase{
			key: &testValue{
				value: randValue,
			},
			value: float64(i),
		}
	}

	for _, v := range testCases {
		sl.Insert(
			&testValue{
				value: v.key.value,
			}, v.value,
		)
	}

	debugPrintTestCases(t, testCases)
	debugPrintSkipList(t, sl)

	for _, v := range testCases {
		v.updatedKey = &testValue{
			value: v.key.value + r.Intn(n)*n,
		}

		sl.Update(&testValue{value: v.key.value}, &testValue{value: v.updatedKey.value}, v.value)
	}

	sort.Slice(testCases, func(i, j int) bool {
		return testCases[i].updatedKey.LT(testCases[j].updatedKey)
	})

	for idx, v := range testCases {
		node := sl.GetNodeByKey(&testValue{
			value: v.updatedKey.value,
		})

		assert.NotNil(t, node)

		rank := sl.GetKeyRank(&testValue{value: v.updatedKey.value})
		assert.Equal(t, idx, rank)
	}

	debugPrintTestCases(t, testCases)
	debugPrintSkipList(t, sl)
}

func TestSkipList_Delete(t *testing.T) {
	comparator := &testComparator[*testValue]{}

	sl := NewSkipList[*testValue, float64](comparator)

	testCases := make([]*testCase, n)

	for i := 0; i < n; i++ {
		randValue := r.Intn(n) + i*n
		testCases[i] = &testCase{
			key: &testValue{
				value: randValue,
			},
			value: float64(i),
		}
	}

	for _, v := range testCases {
		sl.Insert(
			&testValue{
				value: v.key.value,
			}, v.value,
		)
	}

	debugPrintTestCases(t, testCases)
	debugPrintSkipList(t, sl)

	deletedCases := make([]*testCase, 0, n)
	reservedCases := make([]*testCase, 0, n)
	mod := r.Intn(n)

	for idx, v := range testCases {
		if idx%mod == 0 {
			deletedCases = append(deletedCases, &testCase{
				key: &testValue{
					value: v.key.value,
				},
				value: v.value,
			})
		} else {
			reservedCases = append(reservedCases, &testCase{
				key: &testValue{
					value: v.key.value,
				},
				value: v.value,
			})
		}
	}

	for _, v := range deletedCases {
		sl.Delete(&testValue{
			value: v.key.value,
		})
	}

	debugPrintTestCases(t, deletedCases)
	debugPrintSkipList(t, sl)

	sort.Slice(reservedCases, func(i, j int) bool {
		return reservedCases[i].key.LT(reservedCases[j].key)
	})

	for _, v := range deletedCases {
		node := sl.GetNodeByKey(&testValue{value: v.key.value})
		rank := sl.GetKeyRank(&testValue{value: v.key.value})

		assert.Nil(t, node)
		assert.Equal(t, -1, rank)
	}

	for idx, v := range reservedCases {
		node := sl.GetNodeByKey(&testValue{value: v.key.value})
		rank := sl.GetKeyRank(&testValue{value: v.key.value})

		assert.NotNil(t, node)
		assert.Equal(t, idx, rank)
	}

	assert.Equal(t, sl.length, len(reservedCases))
}

func TestSkipList_GetRange(t *testing.T) {
	comparator := &testComparator[*testValue]{}

	sl := NewSkipList[*testValue, float64](comparator)

	testCases := make([]*testCase, n)

	for i := 0; i < n; i++ {
		randValue := r.Intn(n) + i*n
		testCases[i] = &testCase{
			key: &testValue{
				value: randValue,
			},
			value: float64(i),
		}
	}

	for _, v := range testCases {
		sl.Insert(
			&testValue{
				value: v.key.value,
			}, v.value,
		)
	}

	sort.Slice(testCases, func(i, j int) bool {
		return testCases[i].key.LT(testCases[j].key)
	})

	debugPrintTestCases(t, testCases)
	debugPrintSkipList(t, sl)

	var randRange func(int, int) (int, int)
	randRange = func(min, max int) (int, int) {
		leftBound := r.Intn(max-min) + min
		rightBound := r.Intn(max-min) + min
		if leftBound < rightBound {
			return leftBound, rightBound
		}

		if leftBound > rightBound {
			return rightBound, leftBound
		}

		return randRange(min, max)
	}

	t.Run("Test range 1", func(t *testing.T) {
		nodes := sl.GetNodeByRange(0, sl.length-1)
		if !assert.Equal(t, sl.length, len(nodes)) {
			return
		}

		for idx, node := range nodes {
			assert.Equal(t, testCases[idx].key, node.key)
			assert.Equal(t, testCases[idx].value, node.value)
			assert.Equal(t, idx, sl.GetKeyRank(node.key))
		}
	})

	t.Run("Test range 2", func(t *testing.T) {
		rank := sl.GetKeyRank(&testValue{value: testCases[0].key.value})
		t.Log(rank)
	})
}

func TestSkipList_GetNodeByRank(t *testing.T) {
	comparator := &testComparator[*testValue]{}

	sl := NewSkipList[*testValue, float64](comparator)

	testCases := make([]*testCase, n)

	for i := 0; i < n; i++ {
		randValue := r.Intn(n) + i*n
		testCases[i] = &testCase{
			key: &testValue{
				value: randValue,
			},
			value: float64(i),
		}
	}

	for _, v := range testCases {
		sl.Insert(
			&testValue{
				value: v.key.value,
			}, v.value,
		)
	}

	sort.Slice(testCases, func(i, j int) bool {
		return testCases[i].key.LT(testCases[j].key)
	})

	debugPrintSkipList(t, sl)

	t.Run("Test get node by rank 0", func(t *testing.T) {
		node := sl.GetNodeByRank(0)
		assert.NotNil(t, node)
		assert.Equal(t, testCases[0].key, node.key)
	})

	t.Run("Test get node by rank: length-1", func(t *testing.T) {
		node := sl.GetNodeByRank(sl.length - 1)
		assert.NotNil(t, node)
		assert.Equal(t, testCases[len(testCases)-1].key, node.key)
	})

	t.Run("Test get node by rank: length", func(t *testing.T) {
		node := sl.GetNodeByRank(sl.length)
		assert.Nil(t, node)
	})
}
