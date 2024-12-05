package main

import (
	"os"
	"bufio"
	"log"
	"io"
	"fmt"
	"github.com/Nebula5102/AOC2024-Go/day4/internals/word"
)

func CollateWordsearch(file *os.File, wordSearch *[][]byte) {
	reader := bufio.NewReader(file)
	for index,_ := range *wordSearch {
		bytes, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Fatal("Error: ",err)
		} else if err == io.EOF {
			break
		}
		(*wordSearch)[index] = bytes
	}
}

func FindXs(wordSearch *[][]byte, xs *[][]int) {
	for i, row := range *wordSearch {
		for j, letter := range row {
			if string(letter) == "X" {
				coords := []int{i,j}
				*xs = append(*xs,coords)
			}
		} 
	}
}

func FindAs(wordSearch *[][]byte, as *[][]int) {
	for i, row := range *wordSearch {
		for j, letter := range row {
			if string(letter) == "A" {
				coords := []int{i,j}
				*as = append(*as,coords)
			}
		} 
	}
}

func PartOne(wordSearch *[][]byte) {
	var sum int
	var xs [][]int
	partOneWordSearch := *wordSearch
	FindXs(&partOneWordSearch,&xs)

	sum += word.FindWestWords(&partOneWordSearch,&xs)
	sum += word.FindNorthWestWords(&partOneWordSearch,&xs)
	sum += word.FindSouthWestWords(&partOneWordSearch,&xs)
	sum += word.FindEastWords(&partOneWordSearch,&xs)
	sum += word.FindNorthEastWords(&partOneWordSearch,&xs)
	sum += word.FindSouthEastWords(&partOneWordSearch,&xs)
	sum += word.FindNorthWords(&partOneWordSearch,&xs)
	sum += word.FindSouthWords(&partOneWordSearch,&xs)

	fmt.Println(sum)
}

func PartTwo(wordSearch *[][]byte) {
	var sum int
	var as [][]int
	partTwoWordSearch := *wordSearch
	FindAs(&partTwoWordSearch,&as)

	sum += word.FindWestXMAS(&partTwoWordSearch,&as)
	sum += word.FindEastXMAS(&partTwoWordSearch,&as)
	sum += word.FindNorthXMAS(&partTwoWordSearch,&as)
	sum += word.FindSouthXMAS(&partTwoWordSearch,&as)

	fmt.Println(sum)
}

func main() {
	fp, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal("Error: ",err)
	}
	defer fp.Close()
	var wordSearch = make([][]byte,140)
	//var wordSearch = make([][]byte,10)
	CollateWordsearch(fp,&wordSearch)
	PartOne(&wordSearch)
	PartTwo(&wordSearch)
}
