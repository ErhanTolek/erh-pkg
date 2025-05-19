package main

import (
	"erh-pkg/sort"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	// Example usage of the Sort struct
	list := []map[string]string{
		{"a": "d"},
		{"a": "b"},
		{"a": "e"},
	}
	lessFunc := func(a, b string) bool {
		if a < b {
			return true
		}
		return false
	}
	sortClient := sort.NewSort[string, string](lessFunc)

	sortedList := sortClient.BubbleSort(list, "a", lessFunc)

	fmt.Printf("Sorted list: %v", sortedList[0]["a"])
}
