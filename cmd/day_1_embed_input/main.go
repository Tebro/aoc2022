package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var data string

func main() {

	start := time.Now()

	inventories := []int{0}

	lines := strings.Split(data, "\n")

	for _, line := range lines {
		if line == "" {
			inventories = append(inventories, 0)
		} else {
			asNum, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("could not parse line: '%s': %v\n", line, err)
				return
			}
			inventories[len(inventories)-1] += asNum
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
