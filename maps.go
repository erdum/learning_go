package main
import "fmt"

func main() {
	list := [4]string{'a', 'b', 'c', 'd'}
	payload := make(map[string]map[string]string)
	payload["headers"] = map[string]string{
		"Content-Type": "application/json",
	}

	for val, _ := range list {
		fmt.Print(val)
	}
}

