package main

import (
	"flag"
	"fmt"
	"github.com/aliok/github-activities/pkg"
	"log"
	"os"
	"sort"
)

const DefaultPages = 3

var (
	token                = flag.String("token", "", "Github token")
	username             = flag.String("username", "", "Github username")
	pageCount            = flag.Int("pageCount", DefaultPages, fmt.Sprintf("How many pages should the program get from Github api with pagination? Defaults to %d", DefaultPages))
	startDate            = flag.String("startDate", "", "Start date for events in format 2020-12-31. If not passed, all will be returned")
	onlyRelevantToReport = flag.Bool("onlyRelevantToReport", false, "If true, only display events that are relevant for the weekly report (somebody forked my repo, my comments, my PRs, my commits, etc.)")
	groupByRepo          = flag.Bool("groupByRepo", true, "Group by repo, true by default")
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

	if *startDate != "" {
		events = pkg.FilterByDate(events, *startDate)
	}

	if *onlyRelevantToReport {
		events = pkg.FilterInteresting(events, *username)
	}

	// sort by date first
	sort.Slice(events, func(i, j int) bool {
		return events[i].GetCreatedAt() < events[j].GetCreatedAt()
	})

	if *groupByRepo {
		grouped := pkg.GroupByRepo(events)

		keys := make([]string, 0)
		for k := range grouped {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, repo := range keys {
			fmt.Println("=================================================")
			fmt.Printf("\t\t\t%s\n", repo)
			fmt.Println("=================================================")
			for _, event := range grouped[repo] {
				fmt.Println(event)
			}
		}
	} else {
		for _, event := range events {
			fmt.Println(event)
		}
	}

	// TODO: logs with verbosity
	// TODO: warn unhandled event

}
