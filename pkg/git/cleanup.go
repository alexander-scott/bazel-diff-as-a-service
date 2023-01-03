package git

import (
	"fmt"
	"os"
	"strings"
)

// CleanupPath removes everything in the specified path
func CleanupPath(path string) {
	if strings.HasPrefix(path, "/tmp/") {
		fmt.Println("Removing the content at the following path: " + path)
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Skipping file removal due to path not being in /tmp")
	}
}
