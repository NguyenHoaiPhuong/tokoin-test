package main

import (
	"fmt"

	"github.com/NguyenHoaiPhuong/tokoin-test/golang/config"
	"github.com/logrusorgru/aurora"
)

var cfg *config.Configurations

func init() {
	cfg = config.SetupConfig("./resources", "default.toml")
}

func main() {
	// For testing configurations
	// cfg.Print()

	LoadData()

	cont := true
	for cont {
		sel := SelectOption()
		switch sel {
		case 0:
			cont = false
		case 1:
			Search()
		case 2:
			SearchableFields()
		case 3:
			fmt.Println(aurora.Red("Selected option is invalid. Please type again."))
		}
	}
}
