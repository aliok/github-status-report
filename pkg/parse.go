package pkg

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
	"reflect"
)

var eventMap = map[string]reflect.Type{
	"IssueCommentEvent":             reflect.TypeOf(IssueCommentEvent{}),
	"PushEvent":                     reflect.TypeOf(PushEvent{}),
	"ForkEvent":                     reflect.TypeOf(ForkEvent{}),
	"WatchEvent":                    reflect.TypeOf(WatchEvent{}),
	"CreateEvent":                   reflect.TypeOf(CreateEvent{}),
	"DeleteEvent":                   reflect.TypeOf(DeleteEvent{}),
	"PullRequestEvent":              reflect.TypeOf(PullRequestEvent{}),
	"PullRequestReviewCommentEvent": reflect.TypeOf(PullRequestReviewCommentEvent{}),
	"IssuesEvent":                   reflect.TypeOf(IssuesEvent{}),
	"ReleaseEvent":                  reflect.TypeOf(ReleaseEvent{}),
	"PullRequestReviewEvent":        reflect.TypeOf(PullRequestReviewEvent{}),
	"MemberEvent":                   reflect.TypeOf(MemberEvent{}),
}

func Parse(u map[string]interface{}, strict bool) Event {
	eventTypeStr := u["type"].(string)
	eventType := eventMap[eventTypeStr]
	if eventType == nil {
		if strict {
			log.Fatalf("Unknown event type:%s", eventTypeStr)
		} else {
			log.Print("Unknown event type:%s", eventTypeStr)
			return nil
		}

	}

	value := reflect.New(eventType)
	i := value.Interface()
	event := i.(Event)

	if err := parseEvent(u, event); err != nil {
		if strict {
			log.Fatal("Error unmarshalling input", err)
		} else {
			log.Print("Error unmarshalling input", err)
			return nil
		}
	}
	return event
}

func parseEvent(u map[string]interface{}, result Event) error {
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
