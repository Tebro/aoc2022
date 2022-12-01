package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

const inputFileName = "input/day_1"

func main() {

	start := time.Now()

	inventories := []int{0}

	f, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
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

	end := time.Now()

	fmt.Printf("Biggest inventory is: %d\n", sorted[len(sorted)-1])
	fmt.Printf("Sum of the three biggest inventories is: %d\n", sum)

	duration := end.Sub(start) / time.Microsecond

	fmt.Printf("Run took %d microseconds\n", duration)

}
