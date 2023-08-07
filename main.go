package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Component was auto-generated by rea
type Component struct {
	AutomationEmail    string  `json:"automation_email"`
	CreatedAt          string  `json:"created_at"`
	Description        string  `json:"description"`
	Group              bool    `json:"group"`
	GroupId            string  `json:"group_id"`
	Id                 string  `json:"id"`
	Name               string  `json:"name"`
	OnlyShowIfDegraded bool    `json:"only_show_if_degraded"`
	PageId             string  `json:"page_id"`
	Position           float64 `json:"position"`
	Showcase           bool    `json:"showcase"`
	StartDate          string  `json:"start_date"`
	Status             string  `json:"status"`
	UpdatedAt          string  `json:"updated_at"`
}

func (c Component) Print() string {
	return fmt.Sprintf("\"%s\",\"%s\",%s\n", c.Name, c.GroupId, c.Status)
}

// Components is multiple Component
type Components []Component

func main() {
	pageId := os.Getenv("PAGE_ID")
	token := os.Getenv("TOKEN")

	// TODO: Figure out values for pagination
	url := fmt.Sprintf("https://api.statuspage.io/v1/pages/%s/components?page=1&per_page=10", pageId)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	handleError(err)

	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", token))

	resp, err := client.Do(req)
	handleError(err)

	body, err := io.ReadAll(resp.Body)
	handleError(err)

	var comps Components
	handleError(json.Unmarshal(body, &comps))

	for _, c := range comps {
		fmt.Printf(c.Print())
	}
}