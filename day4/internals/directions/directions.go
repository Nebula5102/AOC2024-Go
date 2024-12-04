package directions 


func CheckN(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[0] - 1) < 0 {return -1, -1}
	if string((**wordSearch)[(*coords)[0]-1][(*coords)[1]]) == find {return (*coords)[0] -1, (*coords)[1]}	
	return -1, -1
}

func CheckNE(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[0] - 1) < 0 {return -1, -1}
	if string((**wordSearch)[(*coords)[0]-1][(*coords)[1]+1]) == find {return (*coords)[0]-1 , (*coords)[1]+1}	
	return -1, -1
}

func CheckE(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if string((**wordSearch)[(*coords)[0]][(*coords)[1]+1]) == find {return (*coords)[0] , (*coords)[1]+1}	
	return -1, -1
}

func CheckSE(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[0] + 1) > (len(**wordSearch) - 1) {return -1, -1}
	if string((**wordSearch)[(*coords)[0]+1][(*coords)[1]+1]) == find {return (*coords)[0]+1 , (*coords)[1]+1}	
	return -1, -1
}

func CheckS(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[0] + 1) > (len(**wordSearch) - 1) {return -1, -1}
	if string((**wordSearch)[(*coords)[0]+1][(*coords)[1]]) == find {return (*coords)[0] +1, (*coords)[1]}	
	return -1, -1
}

func CheckSW(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[1] - 1) < 0 || ((*coords)[0] + 1) > (len(**wordSearch) - 1) {return -1, -1}
	if string((**wordSearch)[(*coords)[0]+1][(*coords)[1]-1]) == find {return (*coords)[0]+1 , (*coords)[1]-1}	
	return -1, -1
}

func CheckW(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[1] - 1) < 0 {return -1, -1}
	if string((**wordSearch)[(*coords)[0]][(*coords)[1]-1]) == find {return (*coords)[0] , (*coords)[1]-1}	
	return -1, -1
}

func CheckNW(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[0] - 1) < 0 || ((*coords)[1] - 1) < 0 {return -1, -1}
	if string((**wordSearch)[(*coords)[0]-1][(*coords)[1]-1]) == find {return (*coords)[0] -1, (*coords)[1]-1}	
	return -1, -1
}
