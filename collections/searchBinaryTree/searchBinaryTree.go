package searchbinarytree

import . "golang.org/x/exp/constraints"

func NewSBTree[T Ordered]() SearchBinaryTree[T] {
    return SearchBinaryTree[T]{}
}

type SearchBinaryTree[T Ordered] struct {
    root *BinaryNode[T]
}

func (t *SearchBinaryTree[T]) GetHeight(value T) int {
    return getHeightRecursive(t.root, value, 0)
}

func getHeightRecursive[T Ordered](root *BinaryNode[T], value T, height int) int {
    if root == nil {
        return height
    }

    if root.value == value {
        return height
    }

    if value < root.value {
        return getHeightRecursive(root.left, value, height + 1)
    } else {
        return getHeightRecursive(root.right, value, height + 1)
    }
}

func (t *SearchBinaryTree[T]) AddValue(value T) {
    t.root = addValueRecursive(t.root, value) 
}

func addValueRecursive[T Ordered](root *BinaryNode[T], value T) *BinaryNode[T] {
    if root == nil {
        return &BinaryNode[T] {
            value:value,
        }
    }

    if value < root.value {
        root.left = addValueRecursive(root.left, value)
    } else if value > root.value {
        root.right = addValueRecursive(root.right, value)
    }

    return root
}

func (t *SearchBinaryTree[T]) HasElement(value T) bool {
    return hasElementRecursive(t.root, value)
}

func hasElementRecursive[T Ordered](root *BinaryNode[T], value T) bool {
    if root == nil {
        return false
    }

    if root.value == value {
        return true
    }

    if value < root.value {
        return hasElementRecursive(root.left, value)
    } else {
        return hasElementRecursive(root.right,  value)
    }
}

type BinaryNode[T Ordered] struct {
    left *BinaryNode[T]
    right *BinaryNode[T]
    value T
}
