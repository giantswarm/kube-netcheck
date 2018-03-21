package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"
)

const (
	checkInterval = 10 * time.Second
	warmUpTime    = 30 * time.Second
)

// Common variables.
var (
	description = "Simple network checker for Kubernetes."
	gitCommit   = "n/a"
	name        = "kube-netcheck"
	source      = "https://github.com/giantswarm/kube-netcheck"
)

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

	var connectSocket string
	var listenSocket string
	var help bool

	flag.StringVar(&connectSocket, "connect-socket", "kube-netcheck:6666", "tcp socket to connect to")
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

	// Wait while Kubernetes will start all pods.
	time.Sleep(warmUpTime)

	// Start checks.
	for {
		s := time.Now()

		d := net.Dialer{
			Timeout: 5 * time.Second,
		}

		c, err := d.Dial("tcp", connectSocket)
		if err != nil {
			log.Fatal(err)
		}
		c.Close()

		log.Printf("Successfully connected to %s in %v", connectSocket, time.Since(s))

		time.Sleep(checkInterval)
	}
}
