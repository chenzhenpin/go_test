package main

import (
	"fmt"
	"sort"
)

//学生成绩结构体
type StuScore struct {
    name  string    //姓名
    score int   //成绩
}

type StuScores []StuScore

//Len()
func (s StuScores) Len() int {
	return len(s)
}

//Less():成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

//Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
    stus := StuScores{
					{"alan", 91},
					{"hikerell", 95},
					{"acmfly", 96},
					{"leao", 90},
				}

	//打印未排序的stus数据
    fmt.Println("Default:\n\t",stus)
    //StuScores已经实现了sort.Interface接口,所以可以调用Sort函数进行排序
	sort.Sort(stus)
	//判断是否已经排好顺序，将会打印true
	fmt.Println("IS Sorted?\n\t", sort.IsSorted(stus))
	//打印排序后的stus数据
    fmt.Println("Sorted:\n\t",stus)


    x := 11
    s := []int{3, 6, 8, 11, 45} //注意已经升序排序
    pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
    if pos < len(s) && s[pos] == x {
        fmt.Println(x, "在s中的位置为：", pos)
    } else {
        fmt.Println("s不包含元素", x)
    }

    //x := StuScore{"alan", 95}
    //s := StuScores{
	//				{"alan", 91},
	//				{"hikerell", 95},
	//				{"acmfly", 96},
	//				{"leao", 98},
	//			}
    //pos := sort.Search(len(s), func(i int) bool { return s[i].score >= x.score })
    //if pos < len(s) && s[pos] == x {
    //    fmt.Println(x, "在s中的位置为：", pos)
    //} else {
    //    fmt.Println("s不包含元素", x)
    //}
}
