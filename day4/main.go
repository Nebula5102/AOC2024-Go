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

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error: ",err)
	}
	defer fp.Close()
	var wordSearch = make([][]byte,10)
	CollateWordsearch(fp,&wordSearch)
	PartOne(&wordSearch)
}
