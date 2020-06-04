package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UnstructuredEvent map[string]interface{}

func FetchEvents(eventType string, token string, username string, page int) ([]UnstructuredEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/%s?page=%d", username, eventType, page)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error doing request", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal("Error while fetching events. Request status is not 200:", res.Status)
		return nil, err
	}

	var result []UnstructuredEvent

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatal("Error reading request body:", err)
		return nil, err
	}

	return result, nil
}
