// Application which greets you.
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println(greet())
	invokeBazel()
}

func greet() string {
	return "Hi!"
}

func invokeBazel() {
	output, err := exec.Command("bazel", "version").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(output))
}
