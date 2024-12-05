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

func SortRules(rules *[][]int) {
	rulesHash := make([][]rule,23,100)
	for _,set := range *rules {
		pos := set[0]%23
		for _,hash := range rulesHash[pos]{
			if len(hash.after) < 1 {

				ruleSet := rule{
					value: set[0],
					after: []int{set[1]}, 
				}
				hash.after = append(hash.after,ruleSet)
			} else {
					
				for i := 0; i < len(hash)-1; i++ {
					if 0 < 1 {
						hash[i].after = append(hash[i].after,set[1])
					} else {
						
						ruleSet := rule{
							value: set[0],
							after: []int{set[1]}, 
						}
						hash = append(hash,ruleSet)
					}
				}
			}
		}
	}
	fmt.Println(ruleHash)
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
	SortRules(&rules)
}
