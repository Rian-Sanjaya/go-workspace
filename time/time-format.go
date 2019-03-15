package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	fmt.Println(dt)

	//Format MM-DD-YYYY
	fmt.Println(dt.Format("01-02-2006"))

	//Format MM-DD-YYYY hh:mm:ss
	fmt.Println(dt.Format("01-02-2006 15:04:05"))

	//With short weekday (Mon)
	fmt.Println(dt.Format("01-02-2006 15:04:05 Mon"))

	//With weekday (Monday)
	fmt.Println(dt.Format("01-02-2006 15:04:05 Monday"))

	//Include micro seconds
	fmt.Println(dt.Format("01-02-2006 15:04:05.000000"))

	//Include nano seconds
	fmt.Println(dt.Format("01-02-2006 15:04:05.000000000"))
}
