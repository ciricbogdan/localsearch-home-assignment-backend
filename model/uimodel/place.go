package uimodel

import "github.com/ciricbogdan/localsearch-home-assignment-backend/model/upstreamAPI"

// Place defines the place model which will be sent to the frontend
type Place struct {
	Name         string                  `json:"name"`
	Address      string                  `json:"address"`
	OpeningHours []DayRangeTimeIntervals `json:"openingHours"`
}

// DayRangeTimeIntervals is a pair of day range (dayA - dayB) and time intervals
type DayRangeTimeIntervals struct {
	DayRange      string         `json:"dayRange"`
	TimeIntervals []TimeInterval `json:"timeIntervals"`
}

// TimeInterval defines a time interval for the opening hours
type TimeInterval struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Type  string `json:"type"`
}

var daysOfTheWeek = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

// PlaceFromUpstreamAPI converts the upstreamAPI model to the uimodel
func PlaceFromUpstreamAPI(place upstreamAPI.Place) Place {

	// Init Place
	result := Place{
		Name:    place.DisplayedWhat,
		Address: place.DisplayedWhere,
	}

	// initialize helper vars
	daysMap := place.OpeningHours.Days
	openingHours := make([]DayRangeTimeIntervals, 0)
	rangeStart, rangeEnd := "", ""
	previousTimeIntervals := make([]TimeInterval, 0)
	for _, day := range daysOfTheWeek {
		// initialize for first day of the week
		currentTimeIntervals := make([]TimeInterval, 0)
		if rangeStart == "" {
			rangeStart = day
			rangeEnd = day
			intervals, ok := daysMap[day]
			if ok {
				for _, timeInterval := range intervals {
					previousTimeIntervals = append(previousTimeIntervals, TimeIntervalFromUpstreamAPI(timeInterval))
				}
			} else {
				previousTimeIntervals = append(previousTimeIntervals, TimeInterval{Type: "CLOSED"})
			}

			continue
		}

		intervals, ok := daysMap[day]
		//convert intervals to uimodel TimeInterval
		if ok {
			for _, timeInterval := range intervals {
				currentTimeIntervals = append(currentTimeIntervals, TimeIntervalFromUpstreamAPI(timeInterval))
			}
		} else {
			currentTimeIntervals = append(currentTimeIntervals, TimeInterval{Type: "CLOSED"})
		}

		// compare intervals
		same := true
		if len(previousTimeIntervals) != len(currentTimeIntervals) {
			same = false
		} else {
			for _, prevInterval := range previousTimeIntervals {
				found := false
				for _, currInterval := range currentTimeIntervals {
					if prevInterval == currInterval {
						found = true
					}
				}

				if !found {
					same = false
					break
				}
			}
		}

		// append the time intervals if the current is not the same as the previous
		if !same {
			if rangeStart == rangeEnd {
				openingHours = append(openingHours, DayRangeTimeIntervals{DayRange: rangeStart, TimeIntervals: previousTimeIntervals})
			} else {
				openingHours = append(openingHours, DayRangeTimeIntervals{DayRange: rangeStart + " - " + rangeEnd, TimeIntervals: previousTimeIntervals})
			}
			previousTimeIntervals = currentTimeIntervals
			rangeStart = day
		}

		rangeEnd = day
	}

	if rangeStart == rangeEnd {
		openingHours = append(openingHours, DayRangeTimeIntervals{DayRange: rangeStart, TimeIntervals: previousTimeIntervals})
	} else {
		openingHours = append(openingHours, DayRangeTimeIntervals{DayRange: rangeStart + " - " + rangeEnd, TimeIntervals: previousTimeIntervals})
	}

	result.OpeningHours = openingHours

	return result
}

func TimeIntervalFromUpstreamAPI(timeInterval upstreamAPI.TimeInterval) TimeInterval {
	return TimeInterval{
		Start: timeInterval.Start,
		End:   timeInterval.End,
		Type:  timeInterval.Type,
	}
}
