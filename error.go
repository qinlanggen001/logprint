package logprint

import (
	"fmt"
	"time"
)

func Error(msg interface{}) {
	t := time.Now()
	fmt.Printf("[error] %s: %s \n", t.Format("2006-01-02 15:04:05.000"), msg)
}
