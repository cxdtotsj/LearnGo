package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// in function main.
	for _, url := range os.Args[1:] {
		if err := waitForServer(url); err != nil {
			// fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
			// os.Exit(1)

			// server is this
			log.Fatalf("Site is down: %v\n", err)
		}
	}
}

func waitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s);retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
