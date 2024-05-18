package searchbinarytree

import "testing"

func TestAddValue(t *testing.T) {
    sbtree := NewSBTree[int]()

    sbtree.AddValue(10)

    res := sbtree.HasElement(10)

    if res != true {
        t.Error("Value should exist")
    }
}

func TestHasElement(t *testing.T) {
    sbtree := NewSBTree[int]()

    sbtree.AddValue(10)

    res := sbtree.HasElement(10)

    if res != true {
        t.Error("Value should exist")
    }

    res = sbtree.HasElement(15)

    if res != false {
        t.Error("Value should not exist")
    }

    sbtree.AddValue(15)

    res = sbtree.HasElement(15)

    if res != true {
        t.Error("Value should exist")
    }

}

func TestSearchValue(t *testing.T) {
    sbtree := NewSBTree[int]()

    sbtree.AddValue(10)
    sbtree.AddValue(8)
    sbtree.AddValue(20)
    sbtree.AddValue(3)
    sbtree.AddValue(9)
    sbtree.AddValue(15)
    sbtree.AddValue(25)

    res := sbtree.HasElement(15)

    if res != true {
        t.Error("Value should exist")
    }
}

func TestAddDuplicateValue(t *testing.T) {
    sbtree := NewSBTree[int]()

    sbtree.AddValue(10)
    sbtree.AddValue(10)

    res := sbtree.HasElement(10)
    if res != true {
        t.Error("Value should exist")
    }
}

func TestSearchNonExistentValue(t *testing.T) {
    sbtree := NewSBTree[int]()

    sbtree.AddValue(10)
    sbtree.AddValue(8)
    sbtree.AddValue(20)

    res := sbtree.HasElement(15)
    if res != false {
        t.Error("Value should not exist")
    }
}

func TestAddValuesDescending(t *testing.T) {
    sbtree := NewSBTree[int]()

    values := []int{20, 15, 10, 5, 0}
    for _, value := range values {
        sbtree.AddValue(value)
    }

    for _, value := range values {
        if sbtree.HasElement(value) != true {
            t.Errorf("Value %d should exist", value)
        }
    }
}

func TestAddValuesAscending(t *testing.T) {
    sbtree := NewSBTree[int]()

    values := []int{0, 5, 10, 15, 20}
    for _, value := range values {
        sbtree.AddValue(value)
    }

    for _, value := range values {
        if sbtree.HasElement(value) != true {
            t.Errorf("Value %d should exist", value)
        }
    }
}

func TestAddAndSearchMultipleValues(t *testing.T) {
    sbtree := NewSBTree[int]()

    valuesToAdd := []int{10, 20, 5, 15, 25, 0}
    valuesToSearch := []int{10, 15, 0, 25}

    for _, value := range valuesToAdd {
        sbtree.AddValue(value)
    }

    for _, value := range valuesToSearch {
        if sbtree.HasElement(value) != true {
            t.Errorf("Value %d should exist", value)
        }
    }
}

func TestTreeIsEmptyInitially(t *testing.T) {
    sbtree := NewSBTree[int]()

    if sbtree.HasElement(10) != false {
        t.Error("Value should not exist in an empty tree")
    }
}

func TestGetHeight(t *testing.T) {
    sbtree := NewSBTree[int]()

    values := []int{20, 15, 10, 5, 0}
    for _, value := range values {
        sbtree.AddValue(value)
    }

    
    for i := 0; i < len(values); i++ {
        if sbtree.GetHeight(values[i]) != i {
            t.Errorf("Height should be %d", i)
        }
    }
}
