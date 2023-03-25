package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, "index.html")
	})
	http.HandleFunc("/app.wasm", func(response http.ResponseWriter, request *http.Request) {
		command := exec.Command("go", "build", "-o", "app.wasm", "app/main.go")
		os.Setenv("GOOS", "js")
		os.Setenv("GOARCH", "wasm")
		output, err := command.CombinedOutput()
		log.Print(string(output))
		if err != nil {
			panic(err)
		}
		http.ServeFile(response, request, "app.wasm")
	})
	http.HandleFunc("/wasm_exec.js", func(response http.ResponseWriter, request *http.Request) {
		out, err := exec.Command("go", "env", "GOROOT").Output()
		if err != nil {
			log.Print(err)
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		f := filepath.Join(strings.TrimSpace(string(out)), "misc", "wasm", "wasm_exec.js")
		http.ServeFile(response, request, f)
	})
	http.ListenAndServe(":3000", nil)
}
