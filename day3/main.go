package main

import (
	"os"
	"strings"
	"io"
	"log"
)

func main() {
	fp, err := os.Open("input.txt")
	defer fp.Close()	

	reader := bufio.NewReader(fp)
	for {
		buf, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Error:",err)
		}
	}
}
