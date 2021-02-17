package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
)

func main() {
	r := rand.Int()
	fmt.Printf("Hello: %d\n", r)

	fmt.Print("Gimme a filepath: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	path := filepath.Join("/test", input.Text())
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic("failed to read file: " + err.Error())
	}

	fmt.Print(string(f))
}
