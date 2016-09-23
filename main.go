package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	splitChar = ","
)

var (
	rawURLs string
	sleep   time.Duration
	pause   time.Duration
)

func init() {
	flag.StringVar(&rawURLs, "urls", "", "Comma separated list of urls to scrape")
	flag.DurationVar(&sleep, "sleep", 500*time.Millisecond, "The time to sleep between scraping two urls")
	flag.DurationVar(&pause, "pause", 1*time.Second, "The time to pause between two scraping cycles")
	flag.Parse()
}

func main() {
	if rawURLs == "" {
		abort("Please pass at least one url to scrape")
	}

	urls := strings.Split(rawURLs, splitChar)

	go func() {
		for {
			for _, url := range urls {
				if _, err := http.Get(url); err != nil {
					log.Printf("Error when scraping %q: %s\n", url, err)
				} else {
					log.Printf("Scraped url %q\n", url)
				}
				time.Sleep(sleep)
			}
			log.Printf("Scrape loop finished. Pausing for %s\n", pause)
			time.Sleep(pause)
		}
	}()

	// Handle SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)
	log.Printf("Received signal '%v'. Exiting.", <-ch)
}

func abort(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}
