package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}

// to get the exe file for diff os buid the file using cmd 'GOOS="windows/darwin/linux" go build'
