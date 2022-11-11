package main

import "fmt"

func main() {
	var data []string
	data = append(data, "test")
	for _, value := range data {
		fmt.Println(value)
	}

	json := make(map[string]string)
	json["Hello"] = "World"
	fmt.Println(len(json))
}

