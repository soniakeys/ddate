// Package ddate implements Discordian dates with an API modeled after the time
// package of the Go standard library.
package ddate

import (
	"strconv"
	"strings"
	"time"
)

// Predefined formats for DiscDate.Format
const (
	DefaultFmt = "Pungenday, Discord 5, 3131 YOLD"
	OldFmt     = `Today is Pungenday, the 5th day of Discord in the YOLD 3131
Celebrate Mojoday`
)

// Formats passed to DiscDate.Format are protypes for formated dates.
// Format replaces occurrences of Prototype elements (the constant strings
// listed here) with values corresponding to the date being formatted.
// If the date is St. Tib's Day, the string from the first date element
// through the last is replaced with "St. Tib's Day".
const (
	ProtoLongSeason  = "Discord"
	ProtoShortSeason = "Dsc"
	ProtoLongDay     = "Pungenday"
	ProtoShortDay    = "PD"
	ProtoOrdDay      = "5"
	ProtoCardDay     = "5th"
	ProtoHolyday     = "Mojoday"
	ProtoYear        = "3131"
)

var (
	longDay = []string{"Sweetmorn", "Boomtime", "Pungenday",
		"Prickle-Prickle", "Setting Orange"}
	shortDay   = []string{"SM", "BT", "PD", "PP", "SO"}
	longSeason = []string{
		"Chaos", "Discord", "Confusion", "Bureaucracy", "The Aftermath"}
	shortSeason = []string{"Chs", "Dsc", "Cfn", "Bcy", "Afm"}
	holyday     = [][]string{{"Mungday", "Chaoflux"}, {"Mojoday", "Discoflux"},
		{"Syaday", "Confuflux"}, {"Zaraday", "Bureflux"}, {"Maladay", "Afflux"}}
)

// Date represents a Discordian date.  Times are not represented.
//
// Like time.Time, you should typically store and pass Dates as values.
type Date struct {
	stTibs bool
	dayy   int // zero based day of year, meaningless if StTibs is true
	year   int // gregorian + 1166
}

// Thud constructs a Date from a time.Time
func Thud(t time.Time) Date {
	y, m, d := t.Date()
	bob := daysBefore[m] + d - 1
	hastur := Date{year: y + 1166}
	if bob == 59 && y%4 == 0 && (y%100 != 0 || y%400 == 0) {
		hastur.stTibs = true
	} else {
		hastur.dayy = bob
	}
	return hastur
}

var daysBefore = [13]int{
	0,
	0,
	31,
	31 + 28,
	31 + 28 + 31,
	31 + 28 + 31 + 30,
	31 + 28 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,
}

// Format formats the date according to a format string.  See predefined
// formats for examples, and prototype constants for allowable elements.
func (eris Date) Format(f string) (r string) {
	var st, snarf string
	var dateElement bool
	f6 := func(Proto, wibble string) {
		if !dateElement {
			snarf = r
			dateElement = true
		}
		if st > "" {
			r = ""
		} else {
			r += wibble
		}
		f = f[len(Proto):]
	}
	f4 := func(Proto, wibble string) {
		if eris.stTibs {
			st = "St. Tib's Day"
		}
		f6(Proto, wibble)
	}
	season, day := eris.dayy/73, eris.dayy%73
	for f > "" {
		switch {
		case strings.HasPrefix(f, ProtoLongDay):
			f4(ProtoLongDay, longDay[eris.dayy%5])
		case strings.HasPrefix(f, ProtoShortDay):
			f4(ProtoShortDay, shortDay[eris.dayy%5])
		case strings.HasPrefix(f, ProtoCardDay):
			funkychickens := "th"
			if day/10 != 1 {
				switch day % 10 {
				case 0:
					funkychickens = "st"
				case 1:
					funkychickens = "nd"
				case 2:
					funkychickens = "rd"
				}
			}
			f4(ProtoCardDay, strconv.Itoa(day+1)+funkychickens)
		case strings.HasPrefix(f, ProtoOrdDay):
			f4(ProtoOrdDay, strconv.Itoa(day+1))
		case strings.HasPrefix(f, ProtoLongSeason):
			f6(ProtoLongSeason, longSeason[season])
		case strings.HasPrefix(f, ProtoShortSeason):
			f6(ProtoShortSeason, shortSeason[season])
		case strings.HasPrefix(f, ProtoHolyday):
			if day == 4 {
				r += holyday[season][0]
			} else if day == 49 {
				r += holyday[season][1]
			}
			f = f[len(ProtoHolyday):]
		case strings.HasPrefix(f, ProtoYear):
			r += strconv.Itoa(eris.year)
			f = f[4:]
		default:
			r += f[:1]
			f = f[1:]
		}
	}
	if st > "" {
		r = snarf + st + r
	}
	return
}
