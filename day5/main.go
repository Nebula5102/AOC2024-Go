package main

import (
	"os"
	"io"
	"log"
	"strconv"
	"bufio"
	"strings"
	"slices"
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
	var ruleSet rule
	if beforepage < 0 {
		ruleSet = rule {
			value: before,
			after: []int{},
		}
	}else {
		ruleSet = rule {
			value: before,
			after: []int{beforepage},
		}
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
					break
				} else if index == len(*ruleSets) -1 {
					ruleSet := CreateRule(set[0],set[1]) 
					*ruleSets = append(*ruleSets, ruleSet)
				}
			}
		}
	}
}

func IsIn(is int,in *[]int) bool {
	for _, val := range *in {
		if val == is {
			return true
		}
	}
	return false
}

func CreateValidRules(rules *[]rule, order *[]int, validRs *[]rule) {
	for _, val := range *order {
		for index, rule := range *rules {
			if val == rule.value {
				*validRs = append(*validRs,rule)
				break
			} else if index == len(*rules) -1 {
				ruleSet := CreateRule(val,-1) 
				*validRs = append(*validRs,ruleSet)
			}
		}
	}
}

func OrderCorrect(rules *[]rule, order *[]int) bool {
	var validRs []rule
	CreateValidRules(rules,order,&validRs)
	for i := 0; i<len(validRs); i++ {
		for _,check := range validRs[i+1:] {
			val := false
			if IsIn(check.value,&validRs[i].after){
				val = true
			}
			if !val {
				return false
			}
		}
	}
	return true
}

func PartOne(rules *[][]int, ordering *[][]int) {
	var ruleSets []rule
	CollateRules(&rules, &ruleSets)
	var sum int
	for _, order := range *ordering{
		if OrderCorrect(&ruleSets,&order) {
			sum += (order[len(order)/2])
		}
	}
	fmt.Println(sum)
}

func GetIncorrect(rules *[]rule,ordering **[][]int, incorrectOrders *[][]int) {
	localRules := *rules
	for _, order := range **ordering {
		if !OrderCorrect(&localRules,&order) {
			*incorrectOrders = append(*incorrectOrders,order)
		}
	}
}

func SortRules(validRs *[]rule) {
	var order []rule
	for _, rule := range *validRs {
		if len(order) < 1 {
			order = append(order, rule)
		} else {
			for index, orderedRule := range order {
				if IsIn(orderedRule.value, &rule.after) {
					order = slices.Insert(order,index,rule)
					break
				} else if index == len(order) - 1{
					order = append(order, rule)
				} 
			}
		}
	}
	*validRs = order
}

func Correct(rules *[]rule ,iOrd *[][]int, correctOrder *[][]int) {	
	for _,order := range *iOrd {
		var validRs []rule
		CreateValidRules(rules,&order,&validRs)
		SortRules(&validRs)
		var ordering []int
		for _,val := range validRs {
			ordering = append(ordering, val.value)
		}
		*correctOrder = append(*correctOrder,ordering)
	}
}

func PartTwo(rules *[][]int, ordering *[][]int) {
	var ruleSets []rule
	CollateRules(&rules, &ruleSets)
	var iOrd [][]int
	GetIncorrect(&ruleSets,&ordering,&iOrd)
	var correctOrder [][]int
	Correct(&ruleSets,&iOrd,&correctOrder)
	var sum int
	for _, order := range correctOrder{
		sum += (order[len(order)/2])
	}
	fmt.Println(sum)
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error: ",err)
	}
	var rules [][]int
	var ordering [][]int
	CollatePages(fp, &rules, &ordering)
	fp.Close()
	PartOne(&rules,&ordering)
	PartTwo(&rules,&ordering)
}
