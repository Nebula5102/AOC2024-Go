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
		if string(bytes) == "" {
			pageOrdering = 1
			continue
		}
		if pageOrdering == 0 {
			rule := string(bytes)
			s := strings.Split(rule,"|")
			post := s[len(s)-1]
			post = post[:len(post)-1]
			prior, err := strconv.Atoi(s[0])
			if err != nil {
				log.Fatal("Error: ",err)
			}
			post, err = strconv.Atoi(post)
			if err != nil {
				log.Fatal("Error: ",err)
			}
			nums := []int{prior,post}
			*rules = append(*rules,nums)
		} else {
			ordering := string(bytes)
			s := strings.Split(ordering, ",")
			var vals []int
			for _, val := range s {
				num, err := strconv.Atoi(val)
				if err != nil {
					log.Fatal("Error: ", err)
				}
				vals = append(vals,num)
			} 
		}
	}
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error: ",err)
	}
	var rules [][]int
	var ordering [][]int
	CollatePages(fp, &rules, &ordering)
	fmt.Println(rules)
	fmt.Println(ordering)
	fp.Close()
}
