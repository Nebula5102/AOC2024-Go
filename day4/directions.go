package isDirection


func checkN(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[0] - 1) < 0 {return -1, -1}
	if string((**wordSearch)[(*coords)[0]-1][(*coords)[1]]) == find {return (*coords)[0] -1, (*coords)[1]}	
	return -1, -1
}

func checkNE(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[0] - 1) < 0 {return -1, -1}
	if string((**wordSearch)[(*coords)[0]-1][(*coords)[1]+1]) == find {return (*coords)[0]-1 , (*coords)[1]+1}	
	return -1, -1
}

func checkE(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if string((**wordSearch)[(*coords)[0]][(*coords)[1]+1]) == find {return (*coords)[0] , (*coords)[1]+1}	
	return -1, -1
}

func checkSE(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[0] + 1) > (len(**wordSearch) - 1) {return -1, -1}
	if string((**wordSearch)[(*coords)[0]+1][(*coords)[1]+1]) == find {return (*coords)[0]+1 , (*coords)[1]+1}	
	return -1, -1
}

func checkS(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[0] + 1) > (len(**wordSearch) - 1) {return -1, -1}
	if string((**wordSearch)[(*coords)[0]+1][(*coords)[1]]) == find {return (*coords)[0] +1, (*coords)[1]}	
	return -1, -1
}

func checkSW(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[1] - 1) < 0 || ((*coords)[0] + 1) > (len(**wordSearch) - 1) {return -1, -1}
	if string((**wordSearch)[(*coords)[0]+1][(*coords)[1]-1]) == find {return (*coords)[0]+1 , (*coords)[1]-1}	
	return -1, -1
}

func checkW(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[1] - 1) < 0 {return -1, -1}
	if string((**wordSearch)[(*coords)[0]][(*coords)[1]-1]) == find {return (*coords)[0] , (*coords)[1]-1}	
	return -1, -1
}

func checkNW(wordSearch **[][]byte, coords *[]int, find string) (int, int) {
	if ((*coords)[0] - 1) < 0 || ((*coords)[1] - 1) < 0 {return -1, -1}
	if string((**wordSearch)[(*coords)[0]-1][(*coords)[1]-1]) == find {return (*coords)[0] -1, (*coords)[1]-1}	
	return -1, -1
}
