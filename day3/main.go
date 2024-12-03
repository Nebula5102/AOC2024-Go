package main

import (
	"os"
	"io"
	"log"
	"fmt"
	"strings"
	"regexp"
	"strconv"
	"sort"
)

func Multiply(expression []byte) int {
	re := regexp.MustCompile(`[0-9]+,[0-9]+`)

	nums := re.FindAll(expression,-1)
	str := nums[0]	
	vals := strings.Split(string(str),",")

	i,err := strconv.Atoi(vals[0])
	if err != nil {
		log.Fatal("Error:",err)
	}
	j,err := strconv.Atoi(vals[1])
	if err != nil {
		log.Fatal("Error:",err)
	}

	return i*j
}

func InRange(dos [][]int, lowerBound []int, upperBound []int) int {
	for _, do := range dos {
		if lowerBound[0] < do[0] && do[0] < upperBound[1] {
			return do[0]
		}
	}
	return -1 
}

func PartOne(all []byte, re *regexp.Regexp) {

	expressions := re.FindAll(all,-1)

	var sum int
	for _,expression := range expressions {
		sum += Multiply(expression)
	}
	fmt.Printf("%d\n",sum)

}

func PartTwo(all []byte, re *regexp.Regexp) {
	do := regexp.MustCompile(`do(\()(\))`)
	dos := do.FindAllIndex(all,-1)
	dont := regexp.MustCompile(`don't(\()(\))`)
	donts := dont.FindAllIndex(all,-1)

	sort.Slice(donts, func(i, j int) bool {
		return donts[i][0] < donts[j][0]
	})
	sort.Slice(dos, func(i, j int) bool {
		return donts[i][0] < donts[j][0]
	})

	var doFrom [][]int
	for i, dont := range donts {
		if i < len(donts)-1 {
			val := InRange(dos, dont, donts[i+1])
			if val > 0 {
				ranges := []int{dont[0], val, donts[i+1][0]}
				doFrom = append(doFrom, ranges)
			}
		}
	}

	var	newAll []byte
	for i, do := range doFrom {
		if i == 0 {
			newAll = append(newAll, all[0:do[0]]...)
			newAll = append(newAll, all[do[1]:do[2]]...)
		} else {
			newAll = append(newAll, all[do[1]:do[2]]...)
		}
	}

	expressions := re.FindAll(newAll,-1)
	var sum int
	for _,expression := range expressions {
		sum += Multiply(expression)
	}
	fmt.Printf("%d\n",sum)
}

func main() {
	fp, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal("Error:",err)	
	}
	all, err := io.ReadAll(fp)
	if err != nil && err != io.EOF {
		log.Fatal("Error:",err)	
	}
	fp.Close()

	re := regexp.MustCompile(`mul(\()[0-9]+,[0-9]+(\))`)
	PartOne(all,re) 
	PartTwo(all,re)
}
