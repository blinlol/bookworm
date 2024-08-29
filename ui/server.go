// +build !wasm

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/vugu/vugu/simplehttp"
)

func main() {
	port := os.Getenv("UI_PORT")
	if port == "" {
		port = "8877"
		log.Println("UI_PORT doesn't set. Use default port =", port)
	}
	port = ":" + port
	dev := flag.Bool("dev", false, "Enable development features")
	dir := flag.String("dir", ".", "Project directory")
	httpl := flag.String("http", port, "Listen for HTTP on this :port")
	flag.Parse()
	wd, _ := filepath.Abs(*dir)
	os.Chdir(wd)
	log.Printf("Starting HTTP Server at %q", *httpl)
	h := simplehttp.New(wd, *dev)
	log.Fatal(http.ListenAndServe(*httpl, h))
}