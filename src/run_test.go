package main

import (
	"fmt"
	"testing"
	"time"
)

//每个测试文件必须以 _test.go 结尾，不然 go go_test 不能发现测试文件
//每个测试文件必须导入 testing 包
//功能测试函数必须以 Test 开头，然后一般接测试函数的名字，这个不强求

func Fib(n int) int {
        if n < 2 {
                return n
        }
        return Fib(n-1) + Fib(n-2)
}

func TestFib(t *testing.T) {
	var (
		in       = 7
		expected = 13
	)
	fmt.Println(time.Now())
	actual := Fib(in)
	fmt.Println(time.Now())
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}