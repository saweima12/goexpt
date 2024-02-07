package binarytree_test

import (
	"fmt"
	"goexpt/binarytree"
	"strconv"
	"testing"
)

type Person int

func (pe Person) String() string {
	return strconv.Itoa(int(pe))
}

func (p Person) Compare(other Person) int {
	if p > other {
		return 1
	} else if p == other {
		return 0
	}
	return -1
}

func (p Person) Value() Person {
	return p
}

func TestBt(t *testing.T) {
	data := []Person{
		10, 7, 8, 9, 11, 22, 5, 4, 17, 1, 0, -1,
	}
	nbt := binarytree.NewBTree[Person](data)
	fmt.Println(nbt.InOrder())
	fmt.Println(nbt.Max())
	fmt.Println(nbt.Min())
}
