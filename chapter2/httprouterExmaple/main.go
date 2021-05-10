package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func getCommandOutPut(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := getCommandOutPut("/usr/local/go/bin/go", "version")
	io.WriteString(w, response)

}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, getCommandOutPut("/bin/cat", params.ByName("name")))
}

func main() {
	router := httprouter.New()

	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)
	log.Fatal(http.ListenAndServe(":8000", router))
}
