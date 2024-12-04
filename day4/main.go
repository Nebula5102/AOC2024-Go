package main

import (
	"os"
	"bufio"
	"log"
	"io"
	"fmt"
	"github.com/Nebula5102/AOC2024-Go/day4/isDirection"
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

func FindMs(wordSearch *[][]byte, xs *[][]int, ms *[][]int) {
	for _, coords := range *xs {
		i,j := isDirection.checkNW(&wordSearch,&coords,"M")
		if i >= 0 && j >= 0 {fmt.Println(i,j)}
	}
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error: ",err)
	}
	defer fp.Close()
	var wordSearch = make([][]byte,10)
	CollateWordsearch(fp,&wordSearch)
	var xs [][]int
	FindXs(&wordSearch,&xs)
	var ms [][]int
	FindMs(&wordSearch,&xs,&ms)
}
