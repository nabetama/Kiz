package lib

import (
	"fmt"
	"os"
)

const ABNORMAL_EXIT int = 1

func PanicIf(err error) {
	if err != nil {
		fmt.Errorf("Error: %s", err)
		os.Exit(ABNORMAL_EXIT)

	}
}
