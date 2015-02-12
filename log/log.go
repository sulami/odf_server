package log

import (
	"fmt"
	"time"
)

func Log(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04"), msg)
}

