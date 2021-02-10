package uimodel

import (
	"fmt"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/model/upstreamAPI"
	"strings"
	"time"
)

// Place defines the place model which will be sent to the frontend
type Place struct {
	Name          string                  `json:"name"`
	Address       string                  `json:"address"`
	OpeningHours  []DayRangeTimeIntervals `json:"openingHours"`
	Open          bool                    `json:"open"`
	OpensNextTime string                  `json:"opensNextTime"`
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

		// if there are not previous time intervals then continue
		if len(previousTimeIntervals) == 0 {
			previousTimeIntervals = currentTimeIntervals

			continue
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

	result.Open = checkIfOpen(place)

	if !result.Open {
		result.OpensNextTime = openNextTime(place)
	}

	return result
}

func openNextTime(place upstreamAPI.Place) string {

	now := time.Now().Add(3 * 24 * time.Hour)
	nowString := fmt.Sprintf("%v:%v", now.Hour(), now.Minute())

	weekDay := now.Weekday() - 1
	if weekDay == -1 {
		weekDay = 6
	}

	currentDay := daysOfTheWeek[weekDay]
	index := weekDay

	// check for the current day
	timeIntervals, ok := place.OpeningHours.Days[currentDay]
	if ok {
		for _, timeInterval := range timeIntervals {
			if timeInterval.Start >= nowString {
				return fmt.Sprintf("today, %v", timeInterval.Start)
			}
		}
	}

	index++
	for {
		if index == 7 {
			index = 0
		}
		currentDay := daysOfTheWeek[index]
		// check for interval

		timeIntervals, ok := place.OpeningHours.Days[currentDay]
		if ok {
			resultInterval := timeIntervals[0]
			return fmt.Sprintf("%v, %v", currentDay, resultInterval.Start)
		}

		if weekDay == index {
			break
		}
		index++
	}

	return ""
}

func checkIfOpen(place upstreamAPI.Place) bool {

	now := time.Now()

	weekDay := now.Weekday() - 1
	if weekDay == -1 {
		weekDay = 6
	}

	currentDay := daysOfTheWeek[weekDay]

	timeIntervals, ok := place.OpeningHours.Days[currentDay]
	if !ok {
		return false
	}

	nowString := fmt.Sprintf("%v:%v", now.Hour(), now.Minute())

	for _, timeInterval := range timeIntervals {

		endTime := timeInterval.End

		if endTime == "00:00" {
			endTime = "24:00"
		}

		// 11:30 13:28 14:00

		isAfter := strings.Compare(timeInterval.Start, nowString)
		isBefore := strings.Compare(endTime, nowString)
		if isAfter <= 0 && isBefore > 0 {
			return true
		}

	}

	return false
}

// TimeIntervalFromUpstreamAPI converts timeInteval from the upstreamAPI model to the uimodel
func TimeIntervalFromUpstreamAPI(timeInterval upstreamAPI.TimeInterval) TimeInterval {
	return TimeInterval{
		Start: timeInterval.Start,
		End:   timeInterval.End,
		Type:  timeInterval.Type,
	}
}
