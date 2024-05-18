package ui

import (
	. "finalexam/plateFinder"
	"fmt"
	"strings"
)

func Init() {
    fmt.Println("Plate adder initialized")
    
    plateFinder := NewPlateFinder(1)

    for {
        if MenuSelector(plateFinder) { break } 
    }
}

func MenuSelector(plateFinder PlateFinder) (exit bool) {
    fmt.Println("Select an option")
    fmt.Println("1. Search plate")
    fmt.Println("2. Search plate plus")
    fmt.Println("3. Add plate")
    fmt.Println("0. Exit")

    var input string
    fmt.Scanln(&input)

    fmt.Print("\033[H\033[2J")

    switch input {
    case "0":
        exit = true
    case "1":
        SearchPlateMenu(plateFinder)
    case "2":
        SearchPlusPlateMenu(plateFinder)
    case "3":
        AddPlateMenu(plateFinder)
    default:
        fmt.Println("Invalid option")
    }

    fmt.Println("Press enter to continue")
    fmt.Scanln(&input)

    fmt.Print("\033[H\033[2J")

    return exit
}

func SearchPlateMenu(plateFinder PlateFinder) {
    plate := GetUserInput("Enter a plate number") 

    fmt.Println("Searching for plate:", plate)

    fmt.Println(plateFinder.FindPlate(plate))
}

func SearchPlusPlateMenu(plateFinder PlateFinder) {
    plate := GetUserInput("Enter a plate number") 

    fmt.Println("Searching for plate:", plate)

    fmt.Println(plateFinder.SearchPlus(plate))
}

func AddPlateMenu(plateFinder PlateFinder) {
    plate := GetUserInput("Enter a plate number")

    err := plateFinder.AddPlate(plate)

    if err != nil {
        fmt.Println("Error adding plate:", err)
    }

    fmt.Println("Plate added successfully")

}

func GetUserInput(message string) string {
    fmt.Println(message)

    var input string
    fmt.Scanln(&input)

    return strings.TrimSpace(input)
}
