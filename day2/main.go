package main

import (
	"os"
	"log"
	"bufio"
	"io"
	"regexp"
	"strconv"
)

type level struct {
	safe bool
	dampener bool
	bad int
}

func main () {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	var levels []level
	
	for {
		buf, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error: ", err)
		}
		s := regexp.MustCompile(" ").Split(string(buf[:len(buf)-1]), 10)
		log.Println(s)
		var increasing bool
		var prev int 
		var safe = true 
		var dampener = false
		var hiccups int
		for index, value := range s {
			if index == 0 {
				prev, err = strconv.Atoi(value)	
				if err != nil {
					log.Fatal("Error", err)
				}
				continue 
			}
			cur, err := strconv.Atoi(value)
			if index == 1 {
				if prev < cur {
					increasing = true
				} else {
					increasing = false
				}
			}
			if err != nil {
				log.Fatal("Error", err)
			}
			if cur > prev && !increasing {
				println("in decreasing")
				safe = false
				if index < len(s)-1 {
					future, err := strconv.Atoi(s[index+1])
					if err != nil {
						log.Fatal("Error: ", err)
					}
					if prev > future{
						if dampener {
							dampener = false
							break
						}
						dampener = true 
					}
				}
			} else if cur < prev && increasing {
				println("in increasing")
				safe = false
				if index < len(s)-1 {
					future, err := strconv.Atoi(s[index+1])
					if err != nil {
						log.Fatal("Error: ", err)
					}
					if prev < future{
						if dampener {
							dampener = false
							break
						}
						dampener = true 
					}
				}
			} else if (cur - prev) > 3 || (prev - cur) > 3 {
				println("in 3")
				safe = false
				if index < len(s)-1 {
					future, err := strconv.Atoi(s[index+1])
					if err != nil {
						log.Fatal("Error: ", err)
					}
					if (future - prev) < 3 && (future - prev) > -3{
						if dampener {
							dampener = false
							break
						}
						dampener = true 
					} 
				}
			} else if (cur - prev) < -3 || (prev - cur) < -3 {
				println("in -3")
				if index < len(s)-1 {
					future, err := strconv.Atoi(s[index+1])
					if err != nil {
						log.Fatal("Error: ", err)
					}
					if (future - prev) < 3 && (future - prev) > -3{
						if dampener {
							dampener = false
							break
						}
						dampener = true 
					} 
				}
			} else if cur == prev {
				safe = false
				if index < len(s)-1 {
					future, err := strconv.Atoi(s[index+1])
					if err != nil {
						log.Fatal("Error: ", err)
					}
					if prev < future && increasing{
						if dampener {
							dampener = false
							break
						}
						dampener = true 
					} else if prev > future && !increasing{
						if dampener {
							dampener = false
							break
						}
						dampener = true 
					}
				}
			}
			println(dampener, value)
			prev = cur
		}
		log.Println(s, safe, dampener)
		cur_level := level{
			safe: safe,
			dampener: dampener,
			bad: hiccups,
		}
		log.Println(hiccups, safe)
		levels = append(levels,cur_level)
	}
	var sum int 
	var sum2 int 
	for _, level := range levels {
		if level.safe {
			sum++
			sum2++
		} 
		if level.dampener{
			sum2++
		}
	}
	println(sum)
	println(sum2)
}
