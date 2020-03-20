package client

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/antchfx/xmlquery"
)

type Message struct {
	Title       string
	Description string
}

// GetStatusAPIResponse Incident status
type GetStatusAPIResponse struct {
	HasMessage bool
	Messages   []Message
	Hash       string
}

// GetStatus gets current incident status message
func GetStatus(streetID int, building string) (*GetStatusAPIResponse, error) {
	status := &GetStatusAPIResponse{}

	url := fmt.Sprintf("http://www.1562.kharkov.ua/uk/ppr/rss?street_id=%d&building=%s", streetID, building)
	node, err := xmlquery.LoadURL(url)
	if err != nil {
		return nil, err
	}

	items := xmlquery.Find(node, "//rss/channel/item")

	messages := make([]Message, 0)

	hash := md5.New()
	for _, item := range items {
		status.HasMessage = true
		title := xmlquery.FindOne(item, "//title").InnerText()
		description := xmlquery.FindOne(item, "//description").InnerText()

		io.WriteString(hash, title)
		io.WriteString(hash, description)

		messages = append(messages, Message{
			Title:       title,
			Description: description,
		})
	}

	status.Messages = messages
	status.Hash = fmt.Sprintf("%x", hash.Sum(nil))

	return status, nil
}
