package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

const checkInterval = 10 * time.Second

// Common variables.
var (
	description string = "Simple network checker for Kubernetes."
	gitCommit   string = "n/a"
	name        string = "kube-netcheck"
	source      string = "https://github.com/giantswarm/kube-netcheck"
)

type Check interface {
	Run() error
}

func main() {
	// Print version.
	if (len(os.Args) > 1) && (os.Args[1] == "version") {
		fmt.Printf("Description:    %s\n", description)
		fmt.Printf("Git Commit:     %s\n", gitCommit)
		fmt.Printf("Go Version:     %s\n", runtime.Version())
		fmt.Printf("Name:           %s\n", name)
		fmt.Printf("OS / Arch:      %s / %s\n", runtime.GOOS, runtime.GOARCH)
		fmt.Printf("Source:         %s\n", source)
		return
	}

	var serviceURL string
	var listenSocket string
	var help bool

	flag.StringVar(&serviceURL, "service-url", "http://kube-netcheck:6666", "URL to connect to")
	flag.StringVar(&listenSocket, "listen-socket", ":6666", "Run http server on socket")
	flag.BoolVar(&help, "help", false, "Print usage and exit")
	flag.Parse()

	// Print usage.
	if help {
		flag.Usage()
		return
	}

	// Default http handler with empty http response.
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "")
	})

	// Run http listener.
	log.Printf("Starting http listener on %s", listenSocket)
	go func() {
		log.Fatal(http.ListenAndServe(listenSocket, nil))
	}()

	for {
		resp, err := http.Get(serviceURL)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()

		log.Printf("OK - Checked %s with response %s", serviceURL, resp.Status)

		time.Sleep(checkInterval)
	}
}
