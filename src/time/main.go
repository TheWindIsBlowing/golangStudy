package main

import (
	"fmt"
	"time"
)

// 时间
func f1() {
	p := fmt.Println

	now := time.Now()
	p(now)
	p(now.Location())

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}

// 时间戳
func f2() {
	now := time.Now()
	nanos := now.UnixNano()
	millis := now.UnixMilli()
	secs := now.Unix()

	fmt.Println(nanos)
	fmt.Println(millis)
	fmt.Println(secs)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

func f3() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}

func main() {
	f1()
	fmt.Println("=========================")
	f2()
	fmt.Println("=========================")
	f3()
}
