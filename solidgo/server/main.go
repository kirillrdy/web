package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		_, err := io.WriteString(response, `
<!DOCTYPE html>
<html>
  <head>
     <script src="/wasm_exec.js"></script>
     <script>
         const go = new Go();
         WebAssembly.instantiateStreaming(fetch("app.wasm"), go.importObject).then((result) => {
             go.run(result.instance);
         });
     </script>
  </head>
  <body>
  </body>
</html>
    `)
		if err != nil {
			panic(err)
		}
	})
	http.HandleFunc("/app.wasm", func(response http.ResponseWriter, request *http.Request) {
		command := exec.Command("go", "build", "-o", "app.wasm", "main.go")
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
	log.Print("wasm server started")
	http.ListenAndServe(":3000", nil)
}
