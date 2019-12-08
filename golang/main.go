package main

import (
	"github.com/NguyenHoaiPhuong/tokoin-test/golang/config"
)

var cfg *config.Configurations

func init() {
	cfg = config.SetupConfig("./resources", "default.toml")
}

func main() {
	cfg.Print()
}
