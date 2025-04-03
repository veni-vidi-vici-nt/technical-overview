package configreader

import "fmt"

// Reader is an interface for processing configuration data.
type Reader interface {
	ReadConfig() (map[string]interface{}, error)
}

// Functions that utilize the interface
func PrintConfig(r Reader) {
	config, err := r.ReadConfig()
	if err != nil {
		fmt.Printf("Error reading config: \n", err)
	}

	fmt.Printf("Configuration count: %d\n", len(config))
	for key, value := range config {
		fmt.Printf("%s: %v \n", key, value)
	}
}
