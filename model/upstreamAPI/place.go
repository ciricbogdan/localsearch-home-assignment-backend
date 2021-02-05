package upstreamAPI

// Place defines the place model
type Place struct {
	DisplayedWhat  string       `json:"displayed_what"`
	DisplayedWhere string       `json:"displayed_where"`
	OpeningHours   OpeningHours `json:"opening_hours"`
}

// OpeningHours defines the opening hours of the place
type OpeningHours struct {
	Days map[string][]TimeInterval `json:"days"`
}

// TimeInterval defines a time interval for the opening hours
type TimeInterval struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Type  string `json:"type"`
}
