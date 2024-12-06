package main

import (
	"os"
	"io"
	"log"
	"strconv"
	"bufio"
	"strings"
	"fmt"
)

func CollatePages(fp *os.File, rules *[][]int, ordering *[][]int) {
	reader := bufio.NewReader(fp)
	var pageOrdering = 0
	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Fatal("Error: ",err)
		} else if err == io.EOF {
			break
		}
		if len(bytes) < 2{
			pageOrdering = 1
		} else if pageOrdering == 0 {
			rule := string(bytes[:len(bytes)-1])
			
			s := strings.Split(rule,"|")
			prior, err := strconv.Atoi(s[0])
			if err != nil {
				log.Fatal("Rule Error:", err)
			}
			post, err := strconv.Atoi(s[1])
			if err != nil {
				log.Fatal("Rule Error:", err)
			}
			
			nums := []int{prior,post}
			*rules = append(*rules,nums)
		} else {
			bytes = bytes[:len(bytes)-1]
			order := string(bytes)
			s := strings.Split(order, ",")
			var vals []int
			for _, val := range s {
				num, err := strconv.Atoi(val)
				if err != nil {
					log.Fatal("Error Ordering:", err)
				}
				vals = append(vals,num)
			} 
			*ordering = append(*ordering, vals)
		}
	}
}

type rule struct {
	value int
	after []int
}

func CreateRule(before int, beforepage int) rule {
	 ruleSet := rule {
		value: before,
		after: []int{beforepage},
	}
	return ruleSet
}

func CollateRules(rules **[][]int, ruleSets *[]rule) {
	for _,set := range **rules {
		if len(*ruleSets) < 1{
			ruleSet := CreateRule(set[0],set[1]) 
			*ruleSets = append(*ruleSets, ruleSet)
		} else {
			for index,rule := range *ruleSets {
				if rule.value == set[0] {
					rule.after = append(rule.after,set[1])
					(*ruleSets)[index] = rule
				} else if index == len(*ruleSets) - 1 {
					ruleSet := CreateRule(set[0],set[1]) 
					*ruleSets = append(*ruleSets, ruleSet)
				}
			}
		}
	}
}

func PartOne(rules *[][]int, ordering *[][]int) {
	var ruleSets []rule
	CollateRules(&rules, &ruleSets)
	fmt.Println(ruleSets)
}

func main() {
	fp, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal("Error: ",err)
	}
	var rules [][]int
	var ordering [][]int
	CollatePages(fp, &rules, &ordering)
	fp.Close()
	PartOne(&rules,&ordering)
}
