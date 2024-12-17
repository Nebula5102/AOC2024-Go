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
	multiply := bytes.Repeat([]byte{'*'},len(expression)-2)
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

func Operation(operator byte, res int, v int) int {
	var result int
	if operator == '*' {
		result = res*v
	} else if operator == '+' {
		result = res+v
	}
	return result
}

func PartOne(results []int, expressions [][]int) {
	var ec [][]string
	var correct []int
	for _,expression := range expressions {
		var permutations []string
		combos := GetCombos(expression)
		permutations = append(permutations,string(combos[0]))
		for _,combo := range combos[1:len(combos)-1] {
			permutations = append(permutations,Permutations(combo)...)
		}
		permutations = append(permutations,string(combos[len(combos)-1]))
		ec = append(ec,permutations)
	}
	for _,expression := range expressions {
		combos := ec[expression[0]]
		for _,combo :=range combos {
			res := expression[1]
			for i := 0; i < len(combo); i++ {
				res = Operation(combo[i],res,expression[i+2])
			}
			if res == results[expression[0]] {
				if 1 > len(correct) {
					correct = append(correct,expression[0])
				}
				for ind,i := range correct {
					if i == expression[0] {
						break
					} else if ind == len(correct)-1 {
						correct = append(correct,expression[0])
					} 
				}
			}
		}
	}
	var sum int
	for _,i := range correct {
		sum += results[i]
	}
	fmt.Println(sum)
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
