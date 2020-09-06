package main

import (
	"fmt"
	"os/exec"
	"bytes"
	"log"
	"strings"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("ls")
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Directories: \n%s", out.String())
}
