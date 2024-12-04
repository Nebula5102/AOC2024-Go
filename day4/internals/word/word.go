package word

import (
	"github.com/Nebula5102/AOC2024-Go/day4/internals/directions"
)

func FindNorthWords(wordSearch *[][]byte, xs *[][]int) int {
	var sum int
	for _, coords := range *xs {
		i,j := directions.CheckN(&wordSearch,&coords,"M")
		if i < 0 || j < 0 {continue}
		m := []int{i,j}
		i,j = directions.CheckN(&wordSearch,&m,"A")
		if i < 0 || j < 0 {continue}
		a := []int{i,j}
		i,j = directions.CheckN(&wordSearch,&a,"S")
		if i >= 0 && j >= 0 {sum++}
	}
	return sum
}

func FindNorthEastWords(wordSearch *[][]byte, xs *[][]int) int {
	var sum int
	for _, coords := range *xs {
		i,j := directions.CheckNE(&wordSearch,&coords,"M")
		if i < 0 || j < 0 {continue}
		m := []int{i,j}
		i,j = directions.CheckNE(&wordSearch,&m,"A")
		if i < 0 || j < 0 {continue}
		a := []int{i,j}
		i,j = directions.CheckNE(&wordSearch,&a,"S")
		if i >= 0 && j >= 0 {sum++}
	}
	return sum
}

func FindEastWords(wordSearch *[][]byte, xs *[][]int) int {
	var sum int
	for _, coords := range *xs {
		i,j := directions.CheckE(&wordSearch,&coords,"M")
		if i < 0 || j < 0 {continue}
		m := []int{i,j}
		i,j = directions.CheckE(&wordSearch,&m,"A")
		if i < 0 || j < 0 {continue}
		a := []int{i,j}
		i,j = directions.CheckE(&wordSearch,&a,"S")
		if i >= 0 && j >= 0 {sum++}
	}
	return sum
}

func FindSouthEastWords(wordSearch *[][]byte, xs *[][]int) int {
	var sum int
	for _, coords := range *xs {
		i,j := directions.CheckSE(&wordSearch,&coords,"M")
		if i < 0 || j < 0 {continue}
		m := []int{i,j}
		i,j = directions.CheckSE(&wordSearch,&m,"A")
		if i < 0 || j < 0 {continue}
		a := []int{i,j}
		i,j = directions.CheckSE(&wordSearch,&a,"S")
		if i >= 0 && j >= 0 {sum++}
	}
	return sum
}

func FindSouthWords(wordSearch *[][]byte, xs *[][]int) int {
	var sum int
	for _, coords := range *xs {
		i,j := directions.CheckS(&wordSearch,&coords,"M")
		if i < 0 || j < 0 {continue}
		m := []int{i,j}
		i,j = directions.CheckS(&wordSearch,&m,"A")
		if i < 0 || j < 0 {continue}
		a := []int{i,j}
		i,j = directions.CheckS(&wordSearch,&a,"S")
		if i >= 0 && j >= 0 {sum++}
	}
	return sum
}

func FindSouthWestWords(wordSearch *[][]byte, xs *[][]int) int {
	var sum int
	for _, coords := range *xs {
		i,j := directions.CheckSW(&wordSearch,&coords,"M")
		if i < 0 || j < 0 {continue}
		m := []int{i,j}
		i,j = directions.CheckSW(&wordSearch,&m,"A")
		if i < 0 || j < 0 {continue}
		a := []int{i,j}
		i,j = directions.CheckSW(&wordSearch,&a,"S")
		if i >= 0 && j >= 0 {sum++}
	}
	return sum
}

func FindWestWords(wordSearch *[][]byte, xs *[][]int) int {
	var sum int
	for _, coords := range *xs {
		i,j := directions.CheckW(&wordSearch,&coords,"M")
		if i < 0 || j < 0 {continue}
		m := []int{i,j}
		i,j = directions.CheckW(&wordSearch,&m,"A")
		if i < 0 || j < 0 {continue}
		a := []int{i,j}
		i,j = directions.CheckW(&wordSearch,&a,"S")
		if i >= 0 && j >= 0 {sum++}
	}
	return sum
}

func FindNorthWestWords(wordSearch *[][]byte, xs *[][]int) int {
	var sum int
	for _, coords := range *xs {
		i,j := directions.CheckNW(&wordSearch,&coords,"M")
		if i < 0 || j < 0 {continue}
		m := []int{i,j}
		i,j = directions.CheckNW(&wordSearch,&m,"A")
		if i < 0 || j < 0 {continue}
		a := []int{i,j}
		i,j = directions.CheckNW(&wordSearch,&a,"S")
		if i >= 0 && j >= 0 {sum++}
	}
	return sum
}
