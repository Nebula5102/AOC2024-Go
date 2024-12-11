package main

import (
	"os"
	"bufio"
	"io"
	"fmt"
	"log"
	"strings"
	"strconv"
)

func GetExpressionVariables(fp *os.File,expressions *[][]int) ([]int) {
	reader := bufio.NewReader(fp)
	expressionCount := 0
	var results []int
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Fatal("Error: ",err)
		} else if err == io.EOF {
			break
		}
		line = line[:len(line)-1]
		expression := strings.Split(string(line),": ")
		result, err := strconv.Atoi(expression[0])
		if err != nil {
			log.Fatal("Error:",err)
		}
		vars := strings.Split(expression[1]," ")
		var elements []int
		elements = append(elements, expressionCount)
		expressionCount++
		for _,val := range vars {
			variable, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("Error:",err)
			}
			elements = append(elements,variable)
		}
		*expressions = append(*expressions, elements)
		results = append(results,result)
	}
	return results
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error:",err)
	}
	var expressions [][]int
	results := GetExpressionVariables(fp, &expressions)
	fp.Close()
	fmt.Println(results,expressions)
}
