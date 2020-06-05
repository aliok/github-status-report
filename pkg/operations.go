package pkg

func FilterInteresting(events []Event, username string) []Event {
	ret := make([]Event, 0)

	for _, u := range events {
		event := u.(Event)
		if event.IsRelevantToReport(username) {
			ret = append(ret, event)
		}
	}

	return ret
}

func FilterByDate(events []Event, startDate string) []Event {
	ret := make([]Event, 0)

	for _, u := range events {
		event := u.(Event)
		if startDate < event.GetCreatedAt() {
			ret = append(ret, event)
		}
	}

	return ret
}

func GroupByRepo(events []Event) map[string][]Event {
	grouped := make(map[string][]Event)
	for _, event := range events {
		repo := event.GetRepo().Name
		if arr, ok := grouped[repo]; ok {
			arr = append(arr, event)
			grouped[repo] = arr
		} else {
			arr = make([]Event, 0)
			arr = append(arr, event)
			grouped[repo] = arr
		}
	}
	return grouped
}
