package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"time"
)

//go:embed input.txt
var data []byte

func main() {

	start := time.Now()

	inventories := []int{0}

	currentNum := ""

	for i, b := range data {
		// End of line
		if rune(b) == '\n' {
			// On empty line now
			if rune(data[i-1]) == '\n' {
				continue
			}
			asNum, err := strconv.Atoi(currentNum)
			if err != nil {
				fmt.Printf("could not parse line: '%s': %v\n", currentNum, err)
				return
			}
			inventories[len(inventories)-1] += asNum
			currentNum = ""
			// Next line is empty
			if i < len(data)-1 && rune(data[i+1]) == '\n' {
				inventories = append(inventories, 0)
				continue
			}
		} else {
			currentNum += string(b)
		}
	}

	sorted := sort.IntSlice(inventories)
	sorted.Sort()
	top3 := sorted[len(sorted)-3:]
	sum := 0
	for _, v := range top3 {
		sum += v
	}

	fmt.Printf("Biggest inventory is: %d\n", sorted[len(sorted)-1])
	fmt.Printf("Sum of the three biggest inventories is: %d\n", sum)

	end := time.Now()
	duration := end.Sub(start) / time.Microsecond

	fmt.Printf("Run took %d microseconds\n", duration)

}
