HOW TO TEST CREATE SCRIPT main.go

package main

import (
	"fmt"

	gopkghousekeeping "github.com/arigatosimarmata/go-pkg-housekeeping.git"
)

func main() {
	list_hk := []gopkghousekeeping.HousekeepingModel{
		{
			Directory: "./logs/",
			Filename:  "*.log",
		},
		{
			Directory: "./cache/202309*",
			Filename:  "",
		},
	}
	err := gopkghousekeeping.Housekeeping(list_hk, false)
	if err != nil {
		fmt.Println(err)
	}
}

go run main.go