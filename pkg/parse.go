package pkg

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
)

var eventMap = map[string]interface{}{
	"IssueCommentEvent":             &IssueCommentEvent{},
	"PushEvent":                     &PushEvent{},
	"ForkEvent":                     &ForkEvent{},
	"CreateEvent":                   &CreateEvent{},
	"DeleteEvent":                   &DeleteEvent{},
	"PullRequestEvent":              &PullRequestEvent{},
	"PullRequestReviewCommentEvent": &PullRequestReviewCommentEvent{},
	"IssuesEvent":                   &IssuesEvent{},
}

func Parse(u map[string]interface{}) interface{} {
	var event interface{}
	eventType := u["type"].(string)
	event = eventMap[eventType]
	if event == nil {
		event = &Event{}
	}

	if err := parseEvent(u, event); err != nil {
		log.Fatal("Error unmarshalling input", err)
	}
	return event
}

func parseEvent(u map[string]interface{}, result interface{}) error {
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   result,
		TagName:  "json",
	}
	decoder, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		return fmt.Errorf("error creating decoder: %w", err)
	}
	if err := decoder.Decode(u); err != nil {
		return fmt.Errorf("error decoding input: %w", err)
	}

	return nil
}
