package main

import (
	"flag"
	"fmt"
	"github.com/aliok/github-activities/pkg"
	"log"
	"os"
)

const DefaultPages = 3

var (
	token                = flag.String("token", "", "Github token")
	username             = flag.String("username", "", "Github username")
	pageCount            = flag.Int("pageCount", DefaultPages, fmt.Sprintf("How many pages should the program get from Github api with pagination? Defaults to %d", DefaultPages))
	startDate            = flag.String("startDate", "", "Start date for events in format 2020-12-31. If not passed, all will be returned")
	onlyRelevantToReport = flag.Bool("onlyRelevantToReport", false, "If true, only display events that are relevant for the weekly report (somebody forked my repo, my comments, my PRs, my commits, etc.)")
	verbose              = flag.Bool("verbose", false, "Enable verbose logging")
)

func main() {
	flag.Parse()
	if *token == "" || *username == "" || *pageCount <= 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	unstructuredEvents := make([]pkg.UnstructuredEvent, 0)

	for i := 0; i < *pageCount; i++ {
		events, err := pkg.FetchEvents("events", *token, *username, i+1)
		if err != nil {
			log.Fatalf("Error fetching events: %s", err)
		}
		unstructuredEvents = append(unstructuredEvents, events...)

		receivedEvents, err := pkg.FetchEvents("received_events", *token, *username, i+1)
		if err != nil {
			log.Fatalf("Error fetching receivedEvents: %s", err)
		}
		unstructuredEvents = append(unstructuredEvents, receivedEvents...)
	}

	events := make([]pkg.Event, 0)
	for _, u := range unstructuredEvents {
		event := pkg.Parse(u)
		events = append(events, event)
	}

	if *onlyRelevantToReport {
		events = pkg.Filter(events, *username)
	}

	// TODO: context (e.g. for isInteresting filter methods)
	// TODO: sort
	// TODO: group by repo
	// TODO: logs with verbosity
	// TODO: warn unhandled event

	for _, event := range events {
		fmt.Println(event)
	}
}
