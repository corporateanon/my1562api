package my1562api

import (
	"fmt"

	"github.com/antchfx/xmlquery"
	"github.com/sahilm/fuzzy"
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

// StreetSuggestion street search result found with GetStreetSuggestions
type StreetSuggestion struct {
	Street
	Rank int
}

// StreetSuggestionsList a list of found streets
type StreetSuggestionsList []StreetSuggestion

func (list StreetsList) String(i int) string {
	return list[i].Name
}

// Len length
func (list StreetsList) Len() int {
	return len(list)
}

// GetStreetSuggestions Find streets matching query
func GetStreetSuggestions(q string) StreetSuggestionsList {
	matches := fuzzy.FindFrom(q, Streets)
	suggestions := make(StreetSuggestionsList, 0)
	for _, match := range matches {
		sugg := StreetSuggestion{Street: Streets[match.Index], Rank: match.Score}
		suggestions = append(suggestions, sugg)
	}
	return suggestions
}

//GetStreetByID gets street object by its ID
func GetStreetByID(streetID int) *Street {
	if streetIndex, ok := streetIndexById[streetID]; ok {
		return &Streets[streetIndex]
	}
	return nil
}
