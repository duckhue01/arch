package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now().Format(time.UnixDate))
		time.Sleep(1 * time.Second)
	}
}
