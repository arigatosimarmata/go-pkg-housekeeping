package gopkghousekeeping

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Housekeeping(hk_list []HousekeepingModel, debug bool) error {
	if len(hk_list) < 1 {
		return errors.New("NO DATA TO PROCESS")
	}

	for _, value := range hk_list {
		files, err := filepath.Glob(value.Directory + value.Filename)
		if err != nil {
			if debug {
				fmt.Printf("error glob in housekeepingfile %v, continue\n", err)
			}
			continue
		}

		if debug {
			fmt.Printf("filename %v\n", files)
		}

		for _, file := range files {
			if debug {
				fmt.Printf("remove path %v, continue\n", file)
			}

			err := os.RemoveAll(file)
			if err != nil {
				if debug {
					fmt.Printf("failed to remove file %v, continue\n", err)
				}
				continue
			}
		}
	}

	return nil
}
