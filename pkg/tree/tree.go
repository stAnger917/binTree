package tree

import "fmt"

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
	Index int
}

func (t *Node[T]) Insert(index int, value T) error {
	if t == nil {
		return errTreeIsNil
	}
	if t.Index == 0 {
		t.Index = index
		t.Value = value
		return nil
	}
	if t.Index == index {
		return errNodeAlreadyExist
	}
	if t.Index > index {
		if t.Left == nil {
			t.Left = &Node[T]{Value: value, Index: index}
			return nil
		}
		return t.Left.Insert(index, value)
	}
	if t.Index < index {
		if t.Right == nil {
			t.Right = &Node[T]{Value: value, Index: index}
			return nil
		}
		return t.Right.Insert(index, value)
	}
	return nil
}

func (t *Node[T]) FindMin() int {
	if t.Left == nil {
		return t.Index
	}
	return t.Left.FindMin()
}

func (t *Node[T]) FindMax() int {
	if t.Right == nil {
		return t.Index
	}
	return t.Right.FindMax()
}

func (t *Node[T]) PrintInorder() {
	if t == nil {
		return
	}
	t.Left.PrintInorder()
	fmt.Print(t.Index, t.Value)
	t.Right.PrintInorder()
}

func (t *Node[T]) GetValueByIndex(index int) (T, error) {
	if t == nil {
		return t.Value, errTreeIsNil
	}
	if t.Index == index {
		return t.Value, nil
	}
	if index >= t.Index {
		if t.Right != nil {
			if t.Right.Index == index {
				return t.Right.Value, nil
			} else {
				return t.Right.GetValueByIndex(index)
			}
		}
	}
	if index <= t.Index {
		if t.Left != nil {
			if t.Left.Index == index {
				return t.Left.Value, nil
			} else {
				return t.Left.GetValueByIndex(index)
			}
		}
	}
	return t.Value, errValueNotFound
}
