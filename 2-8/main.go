package main

import (
	"bufio"
	"fmt"
	"os"
)

func WriteMenu(fileName string, foods []string) {
	curDir, _ := os.Getwd()
	path := curDir + "/" + fileName
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, item := range foods {
		fmt.Fprintf(w, "%s\n", item)
	}
}

func main() {
	s := []string{"a", "b", "c", "d"}
	WriteMenu("2-8/menu.txt", s)
}
