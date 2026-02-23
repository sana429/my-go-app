package internal

import (
	"fmt"
	"time"
)

func ShowDateTime() {
	now := time.Now()
	fmt.Println("Current Date:", now.Format("2006-01-02"))
	fmt.Println("Current Time:", now.Format("15:04:05"))
}
