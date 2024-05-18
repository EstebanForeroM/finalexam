package platefinder

import (
	hashmap "finalexam/collections/hashMap"
	"fmt"
	"strings"
)

type PlateFinder struct {
    innerHashMap hashmap.HashTable
}

func NewPlateFinder(size uint) PlateFinder {
    return PlateFinder {
        innerHashMap: hashmap.NewHashMap(size),
    }
}

func (p *PlateFinder) AddPlate(plate string) error {

    plate = strings.ToUpper(plate)

    err := comprovePlate(plate)

    if err != nil { return err }

    err = p.innerHashMap.AddElement(plate)

    if err != nil {
        return fmt.Errorf("Plate already exists")
    }

    return nil
}

func comprovePlate(plate string) error {
    if len(plate) != 4 {
        return fmt.Errorf("Plate must have 4 characters")
    }

    for i := 2; i < 4; i++ {
        err := comproveCharacterIsDigit(plate[i])

        if err != nil {
            return err
        }
    }

    return nil
}

func comproveCharacterIsDigit(character byte) error {
    if character < 48 || character > 57 {
        return fmt.Errorf("Character %c is not a digit", character)
    }

    return nil
}

func (p *PlateFinder) FindPlate(plate string) string {
    plate = strings.ToUpper(plate)
    if p.innerHashMap.HasElement(plate) {
        return "Plate exists"
    }

    return "Plate does not exist"
}

func (p *PlateFinder) SearchPlus(plate string) string {
    plate = strings.ToUpper(plate)
    plateData := p.innerHashMap.SearchPlus(plate)

    if plateData.Exists {
        return fmt.Sprintf("Plate exists in index %d and in tree with height %v", plateData.HashTableIndex, plateData.BinaryTreeHeight)
    }

    return "Plate does not exist"
}


