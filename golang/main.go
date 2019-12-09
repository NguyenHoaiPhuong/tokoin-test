package main

import (
	"fmt"

	"github.com/NguyenHoaiPhuong/tokoin-test/golang/config"
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/utils"
	"github.com/logrusorgru/aurora"
)

var cfg *config.Configurations

func init() {
	cfg = config.SetupConfig("./resources", "default.toml")
}

func introduction() {
	fmt.Println("Type", aurora.Magenta("quit"), "to exit at any time, Press", aurora.Magenta("Enter"), "to continue")
	fmt.Println(aurora.BrightYellow("\tSelect search options:"))
	fmt.Println(aurora.BrightYellow("\t* Press 1 to search"))
	fmt.Println(aurora.BrightYellow("\t* Press 2 to view a list of searchable fields"))
	fmt.Println(aurora.BrightYellow("\t* Type quit to exit"))
}

func main() {
	LoadData()

	cont := true
	for cont {
		introduction()
		sel := utils.ReadInput()
		switch sel {
		case "quit":
			cont = false
		case "1":
			Search()
		case "2":
			SearchableFields()
		default:
			fmt.Println(aurora.Red("Selected option is invalid. Please try again."))
		}
	}
}
