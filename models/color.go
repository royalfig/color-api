package models

type MatchType string

const (
	MatchTypeExact   MatchType = "exact"
	MatchTypeNearest MatchType = "nearest"
)

type Color struct {
	Name          string    `json:"name"`
	OriginalColor string    `json:"original_color"`
	GoodName      string    `json:"good_name"`
	MatchType     MatchType `json:"match_type"`
	MatchColor    string    `json:"match_color"`
}
