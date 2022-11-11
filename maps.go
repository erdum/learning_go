package main

import "fmt"

func main() {
	list := [4]string{"a", "b", "c", "d"}
	payload := make(map[string]map[string]string)
	payload["headers"] = map[string]string{
		"Content-Type": "application/json",
	}

	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}
