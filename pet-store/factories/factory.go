package factories

import (
	"fmt"
	"pet-store/factories/seeders"
	"strconv"
	"strings"
)

func Factory() {
	var purpose string
	var factMod int16
	var total int
	var isPrint bool
	msgFailed := "Failed to generate"

	fmt.Println("\nWelcome to Pet-store Data Factories\n ")

	fmt.Print("Factories Name : ")
	fmt.Scanln(&purpose)

	trimmedPurpose := strings.TrimSpace(purpose)
	if len(trimmedPurpose) > 0 {
		fmt.Print("Factories Module : \n[1] Tag\nChoose a module to auto generate dummy : ")
		fmt.Scanln(&factMod)

		fmt.Print("How many data (Max : 100) : ")
		fmt.Scanln(&total)

		var toogle string
		fmt.Print("\nShow result in command [Y/n] : ")
		fmt.Scanln(&toogle)

		if toogle == "Y" || toogle == "n" {
			if toogle == "Y" {
				isPrint = true
			} else if toogle == "n" {
				isPrint = false
			}

			switch factMod {
			case 1:
				seeders.SeedTags(total, isPrint)
				// ID is empty after run
			default:
				fmt.Println("\n" + msgFailed + " : Invalid module")
			}

			if total > 0 && total <= 100 {
				fmt.Println("\nSuccess run : ", purpose)
				num := strconv.Itoa(total)
				fmt.Println("With " + num + " data created")
			} else {
				if total <= 0 {
					fmt.Println("\n" + msgFailed + " : Total is invalid")
				} else {
					fmt.Println("\n" + msgFailed + " : Data too many")
				}
			}
		} else {
			fmt.Println("\n" + msgFailed + " : Command not valid")
		}
	} else {
		fmt.Println("\n" + msgFailed + " : Factories name cant be empty")
	}
}
