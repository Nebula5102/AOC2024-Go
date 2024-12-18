package main

import (
	"os"
	"bufio"
	"log"
	"fmt"
	"io"
)

func CheckNorth(wordSearch **[][]byte, coords *[]int, find string) bool {
	if string((**wordSearch)[(*coords)[0]-1][(*coords)[1]]) == find {
		return true
	}	
	return false 
}

func CheckEast(wordSearch **[][]byte, coords *[]int, find string) bool {
	if string((**wordSearch)[(*coords)[0]][(*coords)[1]+1]) == find {
		return true
	}	
	return false 
}

func CheckSouth(wordSearch **[][]byte, coords *[]int, find string) bool {
	if string((**wordSearch)[(*coords)[0]+1][(*coords)[1]]) == find {
		return true
	}	
	return false 
}

func CheckWest(wordSearch **[][]byte, coords *[]int, find string) bool {
	if string((**wordSearch)[(*coords)[0]][(*coords)[1]-1]) == find {
		return true
	}	
	return false 
}

func CreateMap(fp *os.File, lab *[][]byte) {
	reader := bufio.NewReader(fp)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Fatal("Error:", err)
		} else if err == io.EOF {
			break
		}

		*lab = append(*lab, line[:len(line)-1])
	}
}

func FindStart(lab *[][]byte) []int {
	for i, row := range *lab {
		for j, position := range row {
			if string(position) == "^" {
				return []int{i,j}
			}
		}
	}
	return []int{-1,-1}
}

func GoNorth(lab *[][]byte, curPos *[]int, marker byte) int {
	if (*curPos)[0] == 0 {
		(*lab)[(*curPos)[0]][(*curPos)[1]] = marker
		return -1
	}
	if CheckNorth(&lab,curPos,"#") {
		(*lab)[(*curPos)[0]][(*curPos)[1]] = '>'
	} else {
		(*lab)[(*curPos)[0]][(*curPos)[1]] = marker
		(*lab)[(*curPos)[0]-1][(*curPos)[1]] = '^'
		(*curPos)[0] -= 1
	}
	return 0
}

func GoEast(lab *[][]byte, curPos *[]int, marker byte) int {
	if (*curPos)[1] == len((*lab)[0])-1 {
		(*lab)[(*curPos)[0]][(*curPos)[1]] = marker
		return -1
	}
	if CheckEast(&lab,curPos,"#") {
		(*lab)[(*curPos)[0]][(*curPos)[1]] = 'v'
	} else {
		(*lab)[(*curPos)[0]][(*curPos)[1]] = marker
		(*lab)[(*curPos)[0]][(*curPos)[1]+1] = '>'
		(*curPos)[1] += 1
	}
	return 0
}

func GoSouth(lab *[][]byte, curPos *[]int, marker byte) int {
	if (*curPos)[0] == len(*lab)-2 {
		(*lab)[(*curPos)[0]][(*curPos)[1]] = marker
		return -1
	}
	if CheckSouth(&lab,curPos,"#") {
		(*lab)[(*curPos)[0]][(*curPos)[1]] = '<'
	} else {
		(*lab)[(*curPos)[0]][(*curPos)[1]] = marker
		(*lab)[(*curPos)[0]+1][(*curPos)[1]] = 'v'
		(*curPos)[0] += 1
	}
	return 0
}

func GoWest(lab *[][]byte, curPos *[]int, marker byte) int {
	if (*curPos)[1] == 0 {
		(*lab)[(*curPos)[0]][(*curPos)[1]] = marker
		return -1
	}
	if CheckWest(&lab,curPos,"#") {
		(*lab)[(*curPos)[0]][(*curPos)[1]] = '^'
	} else {
		(*lab)[(*curPos)[0]][(*curPos)[1]] = marker
		(*lab)[(*curPos)[0]][(*curPos)[1]-1] = '<'
		(*curPos)[1] -= 1
	}
	return 0
}

func PartOne(lab *[][]byte) {
	var sum int
	for _, row := range *lab {
		for _, val := range row {
			if string(val) == "X" {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

func CheckLoop(lab *[][]byte, coords *[]int, start *[]int) bool{
	begin := make([]int,2)
	begin[0] = (*start)[0]
	begin[1] = (*start)[1]

	temp := *lab 
	temp[(*coords)[0]][(*coords)[1]] = '#'
	val := Traverse(&temp,begin,'X')
	(*lab)[(*coords)[0]][(*coords)[1]] = '.'
	return val 
}

func ResetLab(lab *[][]byte, start *[]int) {
	for i , row := range *lab {
		for j , val := range row {
			if (i == (*start)[0] && j == (*start)[1]) {
				(*lab)[i][j] = '^'
			} else if string(val) == "X" {
				(*lab)[i][j] = '.'
			}
		}
	}
}

func PartTwo(lab *[][]byte, start *[]int) {
	var positions [][]int 
	for i , row := range *lab {
		for j , val := range row {
			if (i == (*start)[0] && j == (*start)[1]) {
				(*lab)[i][j] = '^'
			} else if string(val) == "X" {
				(*lab)[i][j] = '.'
				positions = append(positions, []int{i,j})
			}
		}
	}
	var sum int
	for _, coords := range positions {
		if CheckLoop(lab,&coords,start) {
			sum++
		}
		ResetLab(lab,start)
	}
	fmt.Println(sum)
}

func Traverse(lab *[][]byte, position []int, marker byte) bool{
	leaving := 0
	var count int
	for leaving >= 0 {
		facing := string((*lab)[position[0]][position[1]])
		if facing == "^" {
			leaving = GoNorth(lab,&position,marker)
		} else if facing == ">" {
			leaving = GoEast(lab,&position,marker)
		} else if facing == "v" {
			leaving = GoSouth(lab,&position,marker)
		} else if facing == "<" {
			leaving = GoWest(lab,&position,marker)
		}
		count++
		if count > (10*len(*lab)*len((*lab)[0])) {
			return true
		}
	}
	return false
}

func main() {
	var lab [][]byte
	fp, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal("Error:", err)
	}
	CreateMap(fp,&lab)
	fp.Close()
	position := FindStart(&lab)
	startPrtTwo := FindStart(&lab)
	Traverse(&lab,position,'X')
	PartOne(&lab)
	PartTwo(&lab, &startPrtTwo)
}
