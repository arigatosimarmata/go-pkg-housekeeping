package main

import (
	"fmt"

	gopkghousekeeping "github.com/arigatosimarmata/go-pkg-housekeeping"
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
