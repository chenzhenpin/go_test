package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)


	a2, err := strconv.ParseBool("false")
	checkError(err)
	b2, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c2, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d2, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e2, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a2, b2, c2, d2, e2)
}

func checkError(e error){
	if e != nil{
		fmt.Println(e)
	}
}