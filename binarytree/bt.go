package binarytree

type Comparator[T any] interface {
	Compare(other T) int
	Value() T
}

type Node[T any] struct {
	Data  Comparator[T]
	Left  *Node[T]
	Right *Node[T]
}

type BTree[T Comparator[T]] interface {
	Add(item T)
	Remove(item T)
	PreOrder() []T
	InOrder() []T
	PostOrder() []T
	Max() T
	Min() T
}

func NewBTree[T Comparator[T]](data []T) *bTree[T] {
	res := new(bTree[T])
	for i := range data {
		res.Add(data[i])
	}
	return res
}

type bTree[T Comparator[T]] struct {
	root   *Node[T]
	length int
}

func (bt *bTree[T]) Add(item T) {
	if bt.root == nil {
		bt.root = &Node[T]{Data: item}
	} else {
		bt.addNode(bt.root, item)
	}
	bt.length++
}

func (bt *bTree[T]) Remove(item T) bool {
	if bt.root == nil {
		return false
	}

	if bt.removeNode(bt.root, item) {
		bt.length--
		return true
	}

	return false
}

func (bt *bTree[T]) Max() (val T, ok bool) {
	var zero T
	if bt.root == nil {
		return zero, false
	}

	node := bt.root
	for node.Right != nil {
		node = node.Right
	}

	return node.Data.Value(), true
}

func (bt *bTree[T]) Min() (val T, ok bool) {
	var zero T
	if bt.root == nil {
		return zero, false
	}

	node := bt.root
	for node.Left != nil {
		node = node.Left
	}

	return node.Data.Value(), true
}

func (bt *bTree[T]) PreOrder() []T {
	res := make([]T, 0, bt.length)
	bt.preOrder(bt.root, &res)
	return res
}

func (bt *bTree[T]) InOrder() []T {
	res := make([]T, 0, bt.length)
	bt.inOrder(bt.root, &res)
	return res
}

func (bt *bTree[T]) PostOrder() []T {
	res := make([]T, 0, bt.length)
	bt.postOrder(bt.root, &res)
	return res
}

func (bt *bTree[T]) addNode(target *Node[T], newData T) {
	result := target.Data.Compare(newData)

	// targetNode more than newData
	if result >= 0 {
		if target.Left != nil {
			bt.addNode(target.Left, newData)
		} else {
			target.Left = &Node[T]{Data: newData}
		}
	}

	// targetNode less than newData
	if result < 0 {
		if target.Right != nil {
			bt.addNode(target.Right, newData)
		} else {
			target.Right = &Node[T]{Data: newData}
		}
	}
}

func (bt *bTree[T]) removeNode(target *Node[T], data T) bool {
	return true
}

func (bt *bTree[T]) preOrder(target *Node[T], result *[]T) {
	*result = append(*result, target.Data.Value())

	if target.Left != nil {
		bt.preOrder(target.Left, result)
	}

	if target.Right != nil {
		bt.preOrder(target.Right, result)
	}
}

func (bt *bTree[T]) inOrder(target *Node[T], result *[]T) {

	if target.Left != nil {
		bt.inOrder(target.Left, result)
	}
	*result = append(*result, target.Data.Value())

	if target.Right != nil {
		bt.inOrder(target.Right, result)
	}
}

func (bt *bTree[T]) postOrder(target *Node[T], result *[]T) {

	if target.Left != nil {
		bt.postOrder(target.Left, result)
	}

	if target.Right != nil {
		bt.postOrder(target.Right, result)
	}

	*result = append(*result, target.Data.Value())
}
