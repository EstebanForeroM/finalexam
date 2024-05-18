package hashmap

import (
	"errors"
	. "finalexam/collections/searchBinaryTree"
)

type HashTable struct {
    table []SearchBinaryTree[string]
    len uint
}

func NewHashMap(size uint) HashTable {
    if size == 0 { panic("Try to initialize a hashMap with a initial size of 1")}
    return HashTable {
        table: make([]SearchBinaryTree[string], size),
        len: size,
    }
}

func (h *HashTable) AddElement(element string) error {
    index := h.getIndex(element)
    itemTree := &h.table[index] 

    if !(*itemTree).HasElement(element) {
        (*itemTree).AddValue(element)
        return nil
    } else {
        return errors.New("Element already exists")
    }
}

func (h *HashTable) HasElement(element string) bool {
    itemList := h.table[h.getIndex(element)]

    return itemList.HasElement(element)
}

func (h *HashTable) getIndex(str string) int {
    return int(hashString(str) % int64(h.len))
}

func hashString(str string) int64 {
    val := int64(0)

    for _, char := range str {
        val += int64(char)   
    }

    return val
}


func (h *HashTable) SearchPlus(element string) SearchData {
    index := h.getIndex(element)
    itemTree := h.table[index]

    return SearchData {
        Exists: itemTree.HasElement(element),
        HashTableIndex: index,
        BinaryTreeHeight: itemTree.GetHeight(element),
    }
}

type SearchData struct {
    Exists bool
    HashTableIndex int
    BinaryTreeHeight int
}
