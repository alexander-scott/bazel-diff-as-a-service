// Package bazel interacts with bazel
package bazel

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/matttproud/golang_protobuf_extensions/pbutil"
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
	cmd.Stderr = os.Stderr

	// Execute the command
	if err := cmd.Run(); err != nil {
		fmt.Println("error")
	}

	f, err := os.Open("/workspaces/bazel-diff-as-a-service/file.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer f.Close()

	for {
		var msg Target
		_, err := pbutil.ReadDelimited(f, &msg)
		// fmt.Println(n)
		if err != nil {
			fmt.Println(err)
			break
		}
		if src := msg.GetSourceFile(); src != nil {
			fmt.Println(*src.Name)
		}
	}
}
