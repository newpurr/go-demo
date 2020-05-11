package main

import (
	"fmt"
	"github.com/uniplaces/carbon"
	"log"
	"time"
)

func main() {
	// 当前日期时间
	fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())

	// 指定时区的日期事件
	today, _ := carbon.NowInLocation("Japan")
	fmt.Printf("Right now in Japan is %s\n", today)

	// 基础计算
	fmt.Printf("Tomorrow is %s\n", carbon.Now().AddDay())
	fmt.Printf("Last week is %s\n", carbon.Now().SubWeek())

	now := carbon.Now()
	fmt.Println("now is:", now)
	fmt.Println("one second later is:", now.AddSecond())
	fmt.Println("one minute later is:", now.AddMinute())
	fmt.Println("one hour later is:", now.AddHour())
	fmt.Println("3 minutes and 20 seconds later is:", now.AddMinutes(3).AddSeconds(20))
	fmt.Println("2 hours and 30 minutes later is:", now.AddHours(2).AddMinutes(30))
	fmt.Println("3 days and 2 hours later is:", now.AddDays(3).AddHours(2))

	//  下一届奥运会
	nextOlympics, _ := carbon.CreateFromDate(2016, time.August, 5, "Europe/London")
	nextOlympics = nextOlympics.AddYears(4)
	fmt.Printf("Next olympics are in %d\n", nextOlympics.Year())

	// 基础判断
	if carbon.Now().IsWeekend() {
		fmt.Printf("Happy time!")
	}
	if carbon.Now().IsToday() {
		fmt.Printf("Happy Today!")
	}

	// 创建指定时间
	c, err := carbon.Create(2020, time.July, 24, 20, 0, 0, 0, "Japan")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The opening ceremony of next olympics will start at %s in Japan\n", c)

	// 时间比较
	t1, _ := carbon.CreateFromDate(2010, 10, 1, "Asia/Shanghai")
	t2, _ := carbon.CreateFromDate(2011, 10, 20, "Asia/Shanghai")

	fmt.Printf("t1 equal to t2: %t\n", t1.Eq(t2))
	fmt.Printf("t1 not equal to t2: %t\n", t1.Ne(t2))

	fmt.Printf("t1 greater than t2: %t\n", t1.Gt(t2))
	fmt.Printf("t1 less than t2: %t\n", t1.Lt(t2))

	t3, _ := carbon.CreateFromDate(2011, 1, 20, "Asia/Shanghai")
	fmt.Printf("t3 between t1 and t2: %t\n", t3.Between(t1, t2, true))

	now1 := carbon.Now()
	fmt.Printf("Weekday? %t\n", now1.IsWeekday())
	fmt.Printf("Weekend? %t\n", now1.IsWeekend())
	fmt.Printf("LeapYear? %t\n", now1.IsLeapYear())
	fmt.Printf("Past? %t\n", now1.IsPast())
	fmt.Printf("Future? %t\n", now1.IsFuture())

	// 计算两个日期之间相差多少秒、分、小时、天
	vancouver, _ := carbon.Today("Asia/Shanghai")
	london, _ := carbon.Today("Asia/Hong_Kong")
	fmt.Println(vancouver.DiffInSeconds(london, true)) // 0

	ottawa, _ := carbon.CreateFromDate(2000, 1, 1, "America/Toronto")
	vancouver, _ = carbon.CreateFromDate(2000, 1, 1, "America/Vancouver")
	fmt.Println(ottawa.DiffInHours(vancouver, true)) // 3

	fmt.Println(ottawa.DiffInHours(vancouver, false)) // 3
	fmt.Println(vancouver.DiffInHours(ottawa, false)) // -3

	t, _ := carbon.CreateFromDate(2012, 1, 31, "UTC")
	fmt.Println(t.DiffInDays(t.AddMonth(), true))  // 31
	fmt.Println(t.DiffInDays(t.SubMonth(), false)) // -31

	t, _ = carbon.CreateFromDate(2012, 4, 30, "UTC")
	fmt.Println(t.DiffInDays(t.AddMonth(), true)) // 30
	fmt.Println(t.DiffInDays(t.AddWeek(), true))  // 7

	t, _ = carbon.CreateFromTime(10, 1, 1, 0, "UTC")
	fmt.Println(t.DiffInMinutes(t.AddSeconds(59), true))  // 0
	fmt.Println(t.DiffInMinutes(t.AddSeconds(60), true))  // 1
	fmt.Println(t.DiffInMinutes(t.AddSeconds(119), true)) // 1

	// 格式化
	t12 := time.Now()
	fmt.Println(t12.Format(carbon.DefaultFormat))
}
