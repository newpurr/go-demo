package sort

import (
	"fmt"
	"sort"
	"time"
)

type Student struct {
	CreatedAt time.Time
}

type res []*Student

func (a res) Len() int           { return len(a) }
func (a res) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a res) Less(i, j int) bool { return a[i].CreatedAt.After(a[j].CreatedAt) }

func s() {
	var hh res = []*Student{
		{CreatedAt: time.Now()},
		{CreatedAt: time.Now().Add(-200 * time.Hour)},
		{CreatedAt: time.Now().Add(500 * time.Hour)},
	}
	for _, v := range hh {
		fmt.Printf("%#v\n", v.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	fmt.Println("排序后")
	sort.Sort(hh)
	for _, v := range hh {
		fmt.Printf("%#v\n", v.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
