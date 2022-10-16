package main

import "fmt"

func main() {
	config := map[string]string{
		"debug":    "1",
		"logLevel": "info",
		"version":  "1.0",
	}

	for key, value := range config {
		fmt.Println(key, "=", value)
	}

}
