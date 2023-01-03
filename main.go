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

func getBazel(w http.ResponseWriter, r *http.Request) {
	fmt.Println("invoking bazel")
	_, err := io.WriteString(w, "invoking bazel\n")
	if err != nil {
		fmt.Println("Failed to send string back along connection")
	}
	invokeBazel()
}
