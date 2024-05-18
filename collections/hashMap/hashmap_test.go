package hashmap

import "testing"

func TestAddElement(t *testing.T) {
    hmap := NewHashMap(10)

    hmap.AddElement("hello")

    res := hmap.HasElement("hello")

    if res != true {
        t.Error("Value should exist")
    }
}

func TestHasElement(t *testing.T) {
    hmap := NewHashMap(10)

    hmap.AddElement("hello")

    res := hmap.HasElement("hello")

    if res != true {
        t.Error("Value should exist")
    }

    res = hmap.HasElement("world")

    if res != false {
        t.Error("Value should not exist")
    }

    hmap.AddElement("world")

    res = hmap.HasElement("world")

    if res != true {
        t.Error("Value should exist")
    }

}

func TestHasElementWithCollitions(t *testing.T) {
    hmap := NewHashMap(2)

    hmap.AddElement("hello")
    hmap.AddElement("world")
    err := hmap.AddElement("world")


    if err == nil {
        t.Error("Value should not exist: ", err)
    }

    hmap.AddElement("esteban")

    res := hmap.HasElement("world")

    if res != true {
        t.Error("Value should exist")
    }

    res = hmap.HasElement("esteban")

    if res != true {
        t.Error("Value should exist")
    }
}

func TestSearchPlus(t *testing.T) {
    hmap := NewHashMap(1)

    hmap.AddElement("hello")
    hmap.AddElement("world")
    hmap.AddElement("esteban")

    res := hmap.SearchPlus("world")

    if res.Exists != true {
        t.Error("Value should exist: ", res)
    }

    if res.HashTableIndex != 0 {
        t.Error("Index should be 0: ", res)
    }

    if res.BinaryTreeHeight != 1 {
        t.Error("Height should be 1: ", res)
    }
}
