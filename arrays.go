package main
import "fmt"

func main() {
	list := [2]string{ "one", "two" }
	mySlice := list[0:1]
	fmt.Println(len(mySlice))
	for _, val := range list {
		fmt.Println(val)
	}
}