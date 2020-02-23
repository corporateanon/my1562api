package client

import (
	"fmt"

	"github.com/antchfx/xmlquery"
)

// GetStatusAPIResponse Incident status
type GetStatusAPIResponse struct {
	HasMessage  bool
	Title       string
	Description string
}

// GetStatus gets current incident status message
func GetStatus(streetID int, building string) (*GetStatusAPIResponse, error) {
	status := &GetStatusAPIResponse{}

	url := fmt.Sprintf("http://www.1562.kharkov.ua/uk/ppr/rss?street_id=%d&building=%s", streetID, building)
	node, err := xmlquery.LoadURL(url)
	if err != nil {
		return nil, err
	}

	item := xmlquery.FindOne(node, "//rss/channel/item")

	if item == nil {
		status.HasMessage = false
		status.Title = ""
		status.Description = ""
		return status, nil
	}

	title := xmlquery.FindOne(item, "//title").InnerText()
	description := xmlquery.FindOne(item, "//description").InnerText()
	status.HasMessage = true
	status.Title = title
	status.Description = description

	return status, nil
}
