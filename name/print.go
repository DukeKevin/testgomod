package name

import (
	"fmt"
	"time"
)

func Today() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func PrtToday() {
	fmt.Println(Today())
}
