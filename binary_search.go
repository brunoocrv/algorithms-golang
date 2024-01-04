package main

import (
	"fmt"
	"sort"
	"strings"
)

func generateRandomNames() []string {
	names := "Alice, Bob, Charlie, David, Eva, Frank, Grace, Henry, Ivy, Jack"

	splited := strings.Split(names, ",")
	allNames := []string{}

	for _, v := range splited {
		allNames = append(allNames, strings.TrimLeft(v, " "))
	}
	sort.Strings(allNames)

	return allNames
}

func main() {
	listNames := generateRandomNames()
	fmt.Println("Names: ", len(listNames))

	nameSearched := "Eva"

	start := 0
	end := len(listNames) - 1
	
	for start <= end {
		middle := (end + start) / 2
		current := listNames[middle]

		if current == nameSearched {
			fmt.Println("Name finded in list: ", current)
			break
		}

		if current > nameSearched {
			fmt.Println("Current name less than searched")
			end = middle - 1
		}

		if current < nameSearched {
			fmt.Println("Current name bigger than searched")
			start = middle + 1
		}
	}
}