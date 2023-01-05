// Package bazel interacts with bazel
package bazel

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// OutputWorkspaceHashes can be used to run a bazel query on a specific path
func OutputWorkspaceHashes(workingDirectory string) {
	// Build the command
	cmd := exec.Command("bazel", "query", "--output", "streamed_proto", "--order_output=no", "'//external:all-targets' + '//...:all-targets'")
	cmd.Dir = workingDirectory

	// Create file to output stdout to
	file, _ := os.Create("/workspaces/bazel-diff-as-a-service/file.txt")

	// Decide where to write output to
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, file, &stdBuffer)
	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		fmt.Println("error")
	}

	// fmt.Println(stdBuffer.String())
}
