package helper

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	languageTh string = "th"
)

// LoadLocation returns The time zone.
func LoadLocation() *time.Location {
	timeZone, _ := time.LoadLocation("Asia/Bangkok")
	return timeZone
}

// FormatDate format date
func FormatDate(dt *time.Time) (string, string) {
	datefmt := dt.In(LoadLocation()).Format("02 January 2006")

	return dateFormat(datefmt, "en"), dateFormat(datefmt, "th")
}

func dateFormat(datefmt, language string) string {
	if language == languageTh {
		switch getMonth(datefmt) {
		case "January":
			datefmt = strings.Replace(datefmt, "January", "มกราคม", 1)

		case "February":
			datefmt = strings.Replace(datefmt, "February", "กุมภาพันธ์", 1)

		case "March":
			datefmt = strings.Replace(datefmt, "March", "มีนาคม", 1)

		case "April":
			datefmt = strings.Replace(datefmt, "April", "เมษายน", 1)

		case "May":
			datefmt = strings.Replace(datefmt, "May", "พฤษภาคม", 1)

		case "June":
			datefmt = strings.Replace(datefmt, "June", "มิถุนายน", 1)

		case "July":
			datefmt = strings.Replace(datefmt, "July", "กรกฎาคม", 1)

		case "August":
			datefmt = strings.Replace(datefmt, "August", "สิงหาคม", 1)

		case "September":
			datefmt = strings.Replace(datefmt, "September", "กันยายน", 1)

		case "October":
			datefmt = strings.Replace(datefmt, "October", "ตุลาคม", 1)

		case "November":
			datefmt = strings.Replace(datefmt, "November", "พฤศจิกายน", 1)

		case "December":
			datefmt = strings.Replace(datefmt, "December", "ธันวาคม", 1)
		}
	}

	return fmt.Sprintf("%s %s %s", getDay(datefmt), getMonth(datefmt), getYear(datefmt, language))
}

func getDay(datefmt string) string {
	return strings.Split(datefmt, " ")[0]
}

func getMonth(datefmt string) string {
	return strings.Split(datefmt, " ")[1]
}

func getYear(datefmt, language string) string {
	year, _ := strconv.Atoi(strings.Split(datefmt, " ")[2])
	if language == languageTh {
		return strconv.Itoa(year + 543)
	}
	return strconv.Itoa(year)
}

// GetToday get datetime today timezone (th)
func GetToday() time.Time {
	timeNow := TimeNowLocationTH()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())
}

// GetYesterday get datetime yesterday timezone (th)
func GetYesterday() time.Time {
	timeNow := TimeNowLocationTH()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()-1, 0, 0, 0, 0, timeNow.Location())
}

// GetNextDateTime get datetime tomorrow timezone (th)
func GetNextDateTime(day int) time.Time {
	timeNow := TimeNowLocationTH()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()+day, 0, 0, 0, 0, timeNow.Location())
}

// GetPreviousDateTime get datetime yesterday timezone (th)
func GetPreviousDateTime(day int) time.Time {
	timeNow := TimeNowLocationTH()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()-day, 0, 0, 0, 0, timeNow.Location())
}

// GetDateTimeByDate get datetime by date timezone (th)
func GetDateTimeByDate(day time.Time) time.Time {
	timeNow := TimeNowLocationTH()
	return time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, timeNow.Location())
}

// DaysIn returns the number of days in a month for a given year.
func DaysIn(year int, m time.Month) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, LoadLocation()).Day()
}

// IsValidDay check if day is more than 6(Saturday time.Weekday)
func IsValidDay(value time.Weekday) bool {
	return value >= time.Sunday && value <= time.Saturday
}

// IsValidTime check if open and close time is in hours: 00-23 and minutes: 00-59
func IsValidTime(openTime, closeTime string) bool {
	ot := strings.Split(openTime, ":")
	ct := strings.Split(closeTime, ":")
	if len(ot) != 2 || len(ct) != 2 {
		return false
	}

	openHr := ot[0]
	openM := ot[1]
	closeHr := ct[0]
	closeM := ct[1]

	if !((openHr >= "0" && openHr <= "23") || (openM >= "0" && openM <= "59")) ||
		!((closeHr >= "0" && closeHr <= "23") || (closeM >= "0" && closeM <= "59")) {
		return false
	}

	return true
}

// TimeNowLocationTH get time location thai
func TimeNowLocationTH() time.Time {
	return time.Now().In(LoadLocation())
}

// ShortDate short date
func ShortDate(date time.Time, language string) string {
	datefmt := date.Format("02 Jan. 2006")
	if language == "en" {
		return datefmt
	}

	day := strings.Split(datefmt, " ")[0]
	month := getMonthThai(date.Month())
	year := getYear(datefmt, language)

	return fmt.Sprintf("%s %s %s", day, month, year)
}

func getMonthThai(month time.Month) string {
	switch month {
	case time.January:
		return "ม.ค."

	case time.February:
		return "ก.พ."

	case time.March:
		return "มี.ค."

	case time.April:
		return "เม.ย."

	case time.May:
		return "พ.ค."

	case time.June:
		return "มิ.ย."

	case time.July:
		return "ก.ค."

	case time.August:
		return "ส.ค."

	case time.September:
		return "ก.ย."

	case time.October:
		return "ต.ค."

	case time.November:
		return "พ.ย."

	default:
		return "ธ.ค."
	}
}

// AddDatetime add date time only format 2006-01-02 15:04:05
func AddDatetime(datetime string, year, month, day, hour, min, sec int) string {
	layout := "2006-01-02 15:04:05"
	dt, _ := time.Parse(layout, datetime)
	if dt.IsZero() {
		return ""
	}

	dt = time.Date(
		dt.Year()+year,
		dt.Month()+time.Month(month),
		dt.Day()+day,
		dt.Hour()+hour,
		dt.Minute()+min,
		dt.Second()+sec,
		0,
		LoadLocation(),
	)

	return dt.Format(layout)
}

// ReduceDatetime reduce re date time only format 2006-01-02 15:04:05
func ReDatetime(datetime string, year, month, day, hour, min, sec int) string {
	layout := "2006-01-02 15:04:05"
	dt, _ := time.Parse(layout, datetime)
	if dt.IsZero() {
		return ""
	}

	dt = time.Date(
		dt.Year()-year,
		dt.Month()-time.Month(month),
		dt.Day()-day,
		dt.Hour()-hour,
		dt.Minute()-min,
		dt.Second()-sec,
		0,
		LoadLocation(),
	)

	return dt.Format(layout)
}

// ShortYearMonth short month
func ShortYearMonth(date time.Time, language string) string {
	datefmt := date.Format("Jan. 2006")
	if language == "en" {
		return datefmt
	}
	month := getMonthThai(date.Month())
	year := date.Year()

	return fmt.Sprintf("%s %d", month, year)
}

// ShortYearMonth short month
func ShortMonth(date time.Time, language string) string {
	datefmt := date.Format("Jan.")
	if language == "en" {
		return datefmt
	}
	month := getMonthThai(date.Month())

	return month
}

// GetDate get date from string datetime format
func GetDate(date string) string {
	if date != "" {
		s := strings.Split(date, "T")
		if len(s) > 1 {
			date = s[0]
		}

		if date == "0001-01-01" {
			date = "0000-00-00"
		}
	}
	return date
}

func GetTime(datetime string) string {
	if datetime != "" {
		s := strings.Split(datetime, "T")
		if len(s) > 1 {

			split1 := strings.Split(s[1], "Z")
			split2 := strings.Split(s[1], "z")
			split3 := strings.Split(s[1], "+07:00")

			if len(split1) > 1 {
				s[1] = split1[0]
			}
			if len(split2) > 1 {
				s[1] = split2[0]
			}
			if len(split3) > 1 {
				s[1] = split3[0]

			}

			datetime = s[1]
		}
	}
	return datetime
}

// GetDate get date from string datetime format
func GetDateTime(datetime string) string {
	if datetime != "" {
		s := strings.Split(datetime, "T")
		if len(s) > 1 {
			date := s[0]
			if date == "0001-01-01" {
				date = "0000-00-00"
			}

			split1 := strings.Split(s[1], "Z")
			split2 := strings.Split(s[1], "z")
			split3 := strings.Split(s[1], "+07:00")

			if len(split1) > 1 {
				s[1] = split1[0]
			}
			if len(split2) > 1 {
				s[1] = split2[0]
			}
			if len(split3) > 1 {
				s[1] = split3[0]

			}

			datetime = date + " " + s[1]
		}
	}
	return datetime
}

func ConvertDateTime(dateTime, position string) string {
	if dateTime != "" {
		t := strings.Split(dateTime, " ")
		if len(t) == 0 {
			return dateTime
		}
		if position == "0" {
			if t[0] == "0001-01-01" {
				t[0] = "0000-00-00"
			}
			return t[0]
		} else {
			return t[0] + " " + t[1]
		}
	} else {
		return ""
	}
}

func DaysBetween(a, b time.Time) int {

	days := -a.YearDay()
	for year := a.Year(); year < b.Year(); year++ {
		days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}
	days += b.YearDay()

	return days
}

func Date(s string) time.Time {
	d, _ := time.Parse("2006-01-02", s)
	return d
}

func InTimeSpan(start, end, check time.Time) bool {
	return (check.After(start) && check.Before(end)) || (check.Equal(start) || check.Equal(end))
}

func DateTime(s string) time.Time {
	d, _ := time.Parse("2006-01-02 15:04:05", s)
	return d
}
