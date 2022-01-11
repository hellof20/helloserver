package main

import (
	// "bytes"
	"fmt"
	//"io"
	"log"
	"net/http"
	//"os"
	//"os/exec"
)

const webContent = "hello google cloud!"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// cmd := exec.Command("hostname")
	// var out bytes.Buffer
	// multi := io.MultiWriter(os.Stdout, &out)
	// cmd.Stdout = multi

	// if err := cmd.Run(); err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Fprint(w, out.String())
	fmt.Fprint(w, webContent)
}
