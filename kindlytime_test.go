package kindlytime_test

import (
	"github.com/cm-igarashi-ryosuke/aws-logs/lib/kindlytime"
	"testing"
	"time"
	"fmt"
)

func TestOneHourAgo(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2016, 1, 1, 0, 0, 0, 0, jst)
	origin := expected.Add(1 * time.Hour)

	str := "1 hour ago"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}

	str = "1h ago"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}
}

func TestOneHourLater(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2016, 1, 0, 0, 0, 0, 0, jst)
	origin := expected.Add(-1 * time.Hour)

	str := "1 hour later"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}

	str = "1h later"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}
}

func TestAnyMinutesAgo(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2016, 1, 1, 0, 0, 0, 0, jst)
	origin := expected.Add(2 * time.Minute)

	str := "2 minutes ago"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}

	str = "2m ago"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}
}

func TestAnyMinutesLater(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2016, 1, 1, 0, 0, 0, 0, jst)
	origin := expected.Add(-3 * time.Minute)

	str := "3 minutes later"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}

	str = "3m later"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}
}

func TestAnySecondsAgo(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2016, 1, 1, 0, 0, 0, 0, jst)
	origin := expected.Add(40 * time.Second)

	str := "40 seconds ago"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}

	str = "40s ago"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}
}

func TestAnySecondsLater(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2016, 1, 1, 0, 0, 0, 0, jst)
	origin := expected.Add(-55 * time.Second)

	str := "55 seconds later"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}

	str = "55s later"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}
}

func TestAnyDaysAgo(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2016, 1, 1, 0, 0, 0, 0, jst)
	origin := expected.Add(3 * 24 * time.Hour)

	str := "3 days ago"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}

	str = "3d ago"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}
}

func TestAnyDaysLater(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2016, 1, 1, 0, 0, 0, 0, jst)
	origin := expected.Add(-1 * 24 * time.Hour)

	str := "1 day later"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}

	str = "1d later"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}
}

func TestCombinationAgo(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2016, 1, 1, 0, 0, 0, 0, jst)
	origin := expected.Add(1 * time.Hour).Add(30 * time.Minute)

	str := "1 hour 30 minutes ago"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}

	str = "1h 30m ago"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}
}

func TestCombinationLater(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	expected := time.Date(2016, 1, 1, 0, 0, 0, 0, jst)
	origin := expected.Add(-1 * 24 * time.Hour).Add(-7 * time.Hour)

	str := "1 day 7 hours" // laterはデフォルトなので省略可能
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}

	str = "1d 7h"
	if actual, _ := kindlytime.Parse(str, origin); expected != actual {
		t.Fatalf("\"%s\" Should be equal %v(expected) and %v(actual)\n", str, expected, actual)
	}
}

func TestShouldValidTimeWithoutBase(t *testing.T) {
	str := "2 days"
	if _, err := kindlytime.ParseBaseOnCurrentTime(str); err != nil {
		t.Fatalf("\"%s\" `error` should be nil. %#v\n", str, err)
	}
}

func TestShouldNotParseWithUnrecognizeSignature(t *testing.T) {
	str := "2 years"
	if _, err := kindlytime.ParseBaseOnCurrentTime(str); err == nil {
		t.Fatalf("\"%s\" Should be error.\n", str)
	} else {
		msg := err.Error()
		if msg != fmt.Sprintf("Unrecognized input \"%s\"", str) {
			t.Fatalf("\"%s\" Should be equal error message `%s` and `%s`.\n", str, "x", msg)
		}
	}
}

func TestShouldNotParseWithoutNumber(t *testing.T) {
	str := "a year"
	if _, err := kindlytime.ParseBaseOnCurrentTime(str); err == nil {
		t.Fatalf("\"%s\" Should be error.\n", str)
	} else {
		msg := err.Error()
		if msg != fmt.Sprintf("Unrecognized input \"%s\"", str) {
			t.Fatalf("\"%s\" Should be equal error message `%s` and `%s`.\n", str, "x", msg)
		}
	}
}
