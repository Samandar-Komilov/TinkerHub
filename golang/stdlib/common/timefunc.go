package main

import (
	"fmt"
	"time"
)

func TimeExample() {
	now := time.Now()
	fmt.Println(now)

	utc := now.UTC()
	fmt.Println("Now in UTC:", utc)

	custom := time.Date(2025, time.October, 26, 12, 0, 0, 0, time.UTC)
	fmt.Println("Custom timestamp:", custom)
}

func TimeDurationExample() {
	d1 := 2 * time.Second
	d2 := 500 * time.Millisecond
	fmt.Println(d1 + d2)     // 2.5s
	fmt.Println(int64(d1))   // 2000000000 (nanoseconds)
	fmt.Println(d1.Hours())  // 0.000555...
	fmt.Println(d1.String()) // "2s"
}

func TimeMeasureExample() {
	start := time.Now()
	time.Sleep(120 * time.Millisecond)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed)

	// elapsed2 := time.Now().Sub(start)
	// fmt.Println("Elapsed 2 time:", elapsed2)
}

func TimeArithmeticExample() {
	now := time.Now()
	later := now.Add(24 * time.Hour)
	fmt.Println("Later:", later)

	diff := later.Sub(now)
	fmt.Println("Diff:", diff)
}

func TimeComparisonExample() {
	t1 := time.Now()
	t2 := t1.Add(10 * time.Minute)

	fmt.Println(t1.Before(t2))
	fmt.Println(t1.After(t2))
	fmt.Println(t1.Equal(t2))
}

func TickerTimerExample() {
	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	fmt.Println("3 seconds passed")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := range 3 {
		<-ticker.C
		fmt.Println("Tick", i)
	}
}

func ParseFormatExample() {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, "2025-10-26 14:30:00")
	if err != nil {
		panic(err)
	}

	fmt.Println(t)
	fmt.Println(t.Format(layout))
}

func TimeZoneExample() {
	loc, _ := time.LoadLocation("Asia/Tashkent")
	t := time.Now().In(loc)
	fmt.Println("Tashkent Time:", t)
}

func TruncateRoundExample() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Truncate(time.Hour))
	fmt.Println(t.Round(time.Hour))
}

func UnixTimestampsExample() {
	now := time.Now()
	fmt.Println(now.Unix())      // seconds
	fmt.Println(now.UnixMilli()) // milliseconds
	fmt.Println(now.UnixNano())  // nanoseconds

	t := time.Unix(1735221234, 0)
	fmt.Println(t)
}

func DeadlineTimeoutExample() {
	c := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		c <- true
	}()

	select {
	case <-c:
		fmt.Println("Completed")
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
