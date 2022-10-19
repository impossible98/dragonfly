package path

import (
	// import built-in packages
	"os"
)

func CreateDir(dir string) {
	// control flow
	if _, err := os.Stat(dir); err != nil {
		err := os.MkdirAll(dir, 0711)
		// control flow
		if err != nil {
			// return
			return
		}
	} else {
		// return
		return
	}
}
