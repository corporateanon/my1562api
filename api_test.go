package my1562api

import (
	"reflect"
	"testing"
)

func TestGetStatus(t *testing.T) {
	res, err := GetStatus(2334, "4")
	if err != nil {
		t.Fatal(err)
	}
	if res.HasMessage {

		t.Logf("Title:       %s\n", res.Title)
		t.Logf("Description: %s\n", res.Description)
	} else {
		t.Log("No incidents\n")
	}
}

func TestGetStreetSuggestions(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name string
		args args
		want StreetSuggestionsList
	}{
		{
			name: "Match 0",
			args: args{
				q: "широнінців",
			},
			want: StreetSuggestionsList{
				{Street: Street{ID: 449, Name: "вул.  Гвардійців- Широнінців"}, Rank: 49169},
			},
		},
		{
			name: "Match 1",
			args: args{
				q: "клочківська",
			},
			want: StreetSuggestionsList{
				{
					Street: Street{
						ID:   948,
						Name: "вул.  Клочківська",
					},
					Rank: 147605,
				},
				{
					Street: Street{
						ID:   1678,
						Name: "сп.  Клочківський (Пассіонарії)",
					},
					Rank: 49165,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStreetSuggestions(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStreetSuggestions() = %v, want %v", got, tt.want)
			}
		})
	}
}
