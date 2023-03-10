// Application which greets you.
package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/alexander-scott/bazel-diff-as-a-service/internal"
	"github.com/alexander-scott/bazel-diff-as-a-service/pkg/git"
	"github.com/alexander-scott/bazel-diff-as-a-service/pkg/validation"
)

func main() {
	fmt.Println(greet())

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/bazel", getBazel)

	server := &http.Server{
		Addr:              ":3333",
		ReadHeaderTimeout: time.Second,
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Println("Error starting server")
		os.Exit(1)
	}
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

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got request @ getRoot")
	_, err := io.WriteString(w, "This is my website\n")
	if err != nil {
		fmt.Println("Failed to send string back along connection")
	}
}

var gitCloneDest = "/tmp/foo"

func getBazel(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Bazel request")

	// Save the request body into a variable
	body, _ := io.ReadAll(r.Body)

	// Check that the request is valid
	requestData, validationErr := validation.ValidateRequest(body)
	if validationErr != nil {
		sendMessageToClient(w, validationErr.Error())
		fmt.Println("Exiting early due to invalid request: " + validationErr.Error())
		return
	}

	// Clone the specified repo
	fmt.Println("Cloning repo @ " + internal.EscapeStringBeforeLogging(requestData.GitURL))
	repo, cloneErr := git.CloneRepo(requestData.GitURL, gitCloneDest)
	if cloneErr != nil {
		fmt.Println("Exiting early due to git clone failure: " + cloneErr.Error())
		sendMessageToClient(w, cloneErr.Error())
		return
	}

	// Checkout the repo to the starting hash
	fmt.Println("Checking out repo to commit: " + internal.EscapeStringBeforeLogging(requestData.StartCommit))
	checkoutErr := git.Checkout(repo, requestData.StartCommit)
	if checkoutErr != nil {
		fmt.Println("Exiting early due to git checkout failure: " + checkoutErr.Error())
		sendMessageToClient(w, checkoutErr.Error())
		git.CleanupPath(gitCloneDest)
		return
	}

	// Invoke bazel based on the parameters
	invokeBazel()

	git.CleanupPath(gitCloneDest)
	fmt.Println("Finished executing Bazel request")
}

func sendMessageToClient(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := io.WriteString(w, msg)
	if err != nil {
		fmt.Println("Failed to send string back along connection")
		os.Exit(1)
	}
}

/*
curl -X POST -H 'Content-Type: application/json' -d '{"git_url":"https://github.com/kriscfoster/multi-language-bazel-monorepo.git","start_commit":"af7bd7acf3ede83d9b2d86167d53a4fa285e4926","end_commit":"fsbf98231rf"}' -v  http://localhost:3333/bazel
*/
