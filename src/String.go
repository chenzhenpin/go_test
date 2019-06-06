package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main()  {
	//fmt.Println(strings.ContainsAny("team", "i"))
	//fmt.Println(strings.ContainsAny("failure", "u & i"))
	//fmt.Println(strings.ContainsAny("in failure", "h"))
	//fmt.Println(strings.ContainsAny("foo", ""))
	//fmt.Println(strings.ContainsAny("", ""))
	//fmt.Println(strings.Count("five", "")) // 	在 Count 的实现中，处理了几种特殊情况，属于字符匹配预处理的一部分。这里要特别说明一下的是当 sep 为空时，Count 的返回值是：utf8.RuneCountInString(s) + 1
	//fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
	//fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))
	//fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	//fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	//fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	//fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	//fmt.Println(strings.Join([]string{"name=xxx", "age=xx"}, "&"))
	//fmt.Println("ba" + strings.Repeat("na", 2))
	//fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	//fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	//r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	//fmt.Println(r.Replace("This is <b>HTML</b>!"))

	//n, err := strconv.ParseInt("128", 10, 8)
	//if err !=nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(n)

	//对于fmt为 'e', 'E' 和 'f'，prec是指有效数字用于小数点之后的位数；对于fmt为 'g' 和 'G'，prec则是所有的有效数字。例如：
	//s := strconv.FormatFloat(1234.5678, 'g', 3, 64)
	//fmt.Println(strconv.ParseFloat(s, 64))
	//单个字符（单引号扩起来的）却是unicode的。
	single := '\u0015'
	fmt.Println(unicode.IsControl(single))  //true
	single = '\ufe35'
	fmt.Println(unicode.IsControl(single)) // false

	digit := rune('1')
	fmt.Println(unicode.IsDigit(digit)) //true
	fmt.Println(unicode.IsNumber(digit)) //true
	letter := rune('Ⅷ')
	fmt.Println(unicode.IsDigit(letter)) //false
	fmt.Println(unicode.IsNumber(letter)) //true
	 b :=make([]byte, 4)
	utf8.EncodeRune(b,letter)
	fmt.Println(string(b))

}
