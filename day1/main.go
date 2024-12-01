package main 

import (
	"os"
	"log"
	"io"
	"sort"
	"strconv"
)

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error reading file:",err)
	}

	buf := make([]byte,14)
	var list1 []int
	var list2 []int
	for {
		n, err := fp.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal("Error:",err)
		}
		if n == 0 {
			break
		}

		int1, err := strconv.Atoi(string(buf[0:5]))
		int2, err := strconv.Atoi(string(buf[8:13]))
		list1 = append(list1, int1)
		list2 = append(list2, int2)
	}

	sort.Slice(list1, func(i, j int) bool { return list1[i] < list1[j] })	
	sort.Slice(list2, func(i, j int) bool { return list2[i] < list2[j] })	
	var sum int
	var similarity int
	var score int
	for i := 0; i < len(list1); i++ {
		sum += Abs(list1[i] - list2[i])
		find := list1[i]
		for j:= 0; j < len(list2); j++ {
			if find == list2[j] {
				similarity++
			}
		}
		score += find * similarity
		similarity = 0
	}
	println(sum)
	println(score)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
