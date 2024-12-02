package main

import (
	"os"
	"log"
	"bufio"
	"io"
	"regexp"
	"strconv"
)

type level struct {
	safe bool
	increasing bool
	dampener int
	row []int
}

func CollectLevels(fp *os.File, lvls *[]level) {
	reader := bufio.NewReader(fp)
	for {
		buf, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error: ", err)
		}
		s := regexp.MustCompile(" ").Split(string(buf[:len(buf)-1]), 10)
		var nums []int
		for _, val := range s {
			i, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("Error:",err)
			}
			nums = append(nums, i)
		}
		lvl := level{
			safe: true,
			increasing: nums[0] < nums[1] && nums[1] < nums[2],
			dampener: 0,
			row: nums,
		}
		*lvls = append(*lvls, lvl)
	}
}

func SafeReports(lvls *[]level) {
	for i, lvl := range *lvls {
		prev := lvl.row[0]
		for j, cur := range lvl.row {
			// 0 1 5 3 4 5
			if prev > cur && lvl.increasing {
				lvl.safe = false
			// 5 4 3 7 1 0
			} else if prev < cur && !lvl.increasing {
				lvl.safe = false
			// 5 4 3 3 1 0
			} else if prev == cur && j > 0 {
				lvl.safe = false
			// 0 1 2 3 7 8
			} else if (prev - cur) > 3 {
				lvl.safe = false
			// 8 7 3 2 1 0
			} else if (cur - prev) > 3 {
				lvl.safe = false
			}
			prev = cur
		}
		(*lvls)[i] = lvl
	}	
}

func PartOne(lvls *[]level) {
	var sum int
	for _, lvl := range *lvls {
		if lvl.safe {
			sum++
		}
	}
	println(sum)
}

func PartTwo(lvls *[]level) {
	var sum int
	for _, lvl := range *lvls {
		if lvl.safe {
			sum++
		}
		if !lvl.safe && lvl.dampener >= 1 && lvl.dampener < 2 {
			sum++
		}
	}
	println(sum)
}

func main () {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	
	var levels []level
	CollectLevels(fp, &levels) 
	SafeReports(&levels)
	PartOne(&levels)
	PartTwo(&levels)
	fp.Close()
}
