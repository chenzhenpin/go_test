package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/goinaction/code/chapter3/words"
)

func main()  {
	fmt.Print("hello world")
	filename := os.Args[1]
	contents , err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	text :=string(contents)
	count :=words.CountWords(text)
	fmt.Printf("There are %d words in your text． \n", count)
}