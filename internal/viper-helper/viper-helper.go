package viperhelper

import (
	"fmt"

	"github.com/spf13/viper"
)

// Helper Function to view Viper Settings
func ShowAllViperSettings() {
	settings := viper.AllSettings()
	fmt.Println("=============VIPER START==================")

	for key, value := range settings {
		fmt.Printf("%s: %#v\n", key, value)
	}
	fmt.Println("=============VIPER END====================")
}
