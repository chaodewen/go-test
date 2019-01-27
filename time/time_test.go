package time

import (
	"testing"
	"fmt"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Printf("time.Now().Second(): %+v\n", time.Now().Second())
	fmt.Printf("time.Now().Hour(): %+v\n", time.Now().Hour())

	timestamp := int64(1539310767)
	time1, time2 := time.Unix(timestamp, 0), time.Unix(timestamp, 0)
	fmt.Print(time1 == time2)

	time3 := time.Unix(0, 1543817312000000000)
	fmt.Printf("time3: %+v", time3)

	d := time.Now().Add(time.Hour).Sub(time.Now()).Seconds()
	fmt.Printf("d: %+v", int64(d))
}
