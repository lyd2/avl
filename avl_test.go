package avl

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAvl(t *testing.T) {
	const maxCount = 100 * 10000

	avlTree := New()

	for i := 0; i < maxCount; i++ {
		avlTree.Insert(strconv.FormatInt(int64(i), 10), i)
	}

	for i := 0; i < maxCount; i++ {
		val, err := avlTree.Search(strconv.FormatInt(int64(i), 10))

		if err != nil {
			t.Error("key=" + strconv.FormatInt(int64(i), 10) + ", " + err.(string))
		}

		if val.(int) != i {
			t.Error("key=" + strconv.FormatInt(int64(i), 10))
		}
	}
}

func TestAvl_Insert(t *testing.T) {
	tests := []struct {
		input      []string // 输入节点序列
		heightList []int    // 中序遍历时节点高度的输出序列
	}{
		{[]string{"1", "2", "3", "4", "5", "6", "7", "8"}, []int{1, 2, 1, 4, 1, 3, 2, 1}},
		{[]string{"8", "5", "3", "4", "2", "1"}, []int{1, 2, 3, 1, 2, 1}},
	}

	for _, v := range tests {
		avlTree := New()
		for _, val := range v.input {
			avlTree.Insert(val, val)
		}

		it := avlTree.InOrder()
		if len(it.List) != len(v.heightList) {
			t.Error("Inconsistent number of nodes")
		}
		for i := 0; i < len(it.List); i++ {
			if v.heightList[i] != it.List[i].GetHeight() {
				t.Error(fmt.Sprintf("Node height error, key=%s, i=%d, nodeH=%d, iH=%d", it.List[i].GetKey(), i, it.List[i].GetHeight(), v.heightList[i]))
			}
		}
	}
}
