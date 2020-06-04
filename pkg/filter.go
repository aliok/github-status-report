package pkg

func Filter(events []Event, username string) []Event {
	ret := make([]Event, 0)

	for _, u := range events {
		event := u.(Event)
		if event.IsRelevantToReport(username) {
			ret = append(ret, event)
		}
	}

	return ret
}
