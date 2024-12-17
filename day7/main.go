package main

import (
	"os"
	"bufio"
	"io"
	"fmt"
	"log"
	"strings"
	"bytes"
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

func Permutations(combination []byte) []string {
	var ps []string
	var order = make([]byte,len(combination))
	copy(order,combination)
	if len(combination) < 3 {
		ps = append(ps,string(combination))
		ps = append(ps,string(combination[1])+string(combination[0]))
		return ps
	} else {
		for i := 0; i<len(combination); i++ {
			copy(combination,order)
			combination[0] = order[i]
			combination[i] = order[0]
			list := Permutations(combination[1:])
			for _, combo := range list {
				combo = string(combination[0])+combo
				ps = append(ps, combo)
			}
		}
	}
	return ps
}

func GetCombos(expression []int) [][]byte {
	var combos [][]byte
	multiply := bytes.Repeat([]byte{'*'},len(expression)-1)
	for i := 0; i < len(multiply)+1; i++ {
		var temp = make([]byte,len(multiply))
		copy(temp,multiply)
		for j:=0; j<i; j++{
			temp[j] = '+'
		}
		combos = append(combos, temp)
	}
	return combos
}

func PartOne(results []int, expressions [][]int) {
	for _,expression := range expressions {
		combos := GetCombos(expression)
		var permutations []string
		permutations = append(permutations,string(combos[0]))
		for _,combo := range combos[1:len(combos)-1] {
			permutations = append(permutations,Permutations(combo)...)
		}
		permutations = append(permutations,string(combos[len(combos)-1]))
	}
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error:",err)
	}
	var expressions [][]int
	results := GetExpressionVariables(fp, &expressions)
	fp.Close()
	PartOne(results,expressions)
}
