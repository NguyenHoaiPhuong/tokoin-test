package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Configurations includes all configurations for the App
type Configurations struct {
	filePath string
	fileName string

	URLs URLConfig
}

// URLConfig : urls configuration
type URLConfig struct {
	UserURL         string
	TicketURL       string
	OrganizationURL string
}

var globalConfig *Configurations

func init() {
	globalConfig = new(Configurations)
}

// SetupConfig : init configurations
func SetupConfig(filePath, fileName string) *Configurations {
	globalConfig.filePath = filePath
	globalConfig.fileName = fileName
	globalConfig.init()

	return globalConfig
}

func (c *Configurations) init() {
	strSlice := strings.Split(c.fileName, ".")
	fileName := strSlice[0]
	ext := strSlice[1]
	// Set the file name of the configurations file
	viper.SetConfigName(fileName)
	viper.SetConfigType(ext)

	// Set the path to look for the configurations file
	viper.AddConfigPath(c.filePath)

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}

	// Set undefined variables
	viper.SetDefault("URLs.UserURL", "")
	viper.SetDefault("URLs.TicketURL", "")
	viper.SetDefault("URLs.OrganizationURL", "")

	err := viper.Unmarshal(c)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
	}

	// c.Print()
}

// Print : prints all configurations on the terminal
func (c *Configurations) Print() {
	// Reading variables using the model
	fmt.Println("Reading variables using the model.")
	fmt.Println("UserURL is\t\t", c.URLs.UserURL)
	fmt.Println("TicketURL is\t\t", c.URLs.TicketURL)
	fmt.Println("OrganizationURL is\t", c.URLs.OrganizationURL)

	// Reading variables without using the model
	fmt.Println("\nReading variables without using the model.")
	fmt.Println("UserURL is\t\t", viper.GetString("URLs.UserURL"))
	fmt.Println("TicketURL is\t\t", viper.GetString("URLs.TicketURL"))
	fmt.Println("OrganizationURL is\t", viper.GetString("URLs.OrganizationURL"))
}

// GlobalConfig init new config and return it
func GlobalConfig() *Configurations {
	return globalConfig
}
