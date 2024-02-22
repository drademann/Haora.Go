package app_test

import (
	"testing"
	"time"

	"haora/app"
)

func TestNewDay(t *testing.T) {
	t.Run("should initialise the days date", func(t *testing.T) {
		date := time.Date(2024, time.February, 21, 14, 58, 42, 12, time.Local)
		day := app.NewDay(date)

		if len(day.Tasks) != 0 {
			t.Errorf("expected new day to have no tasks, but found %d", len(day.Tasks))
		}
		if !day.Finished.IsZero() {
			t.Errorf("didn't expect new day to be finished from the beginning")
		}
	})
}

func TestHasNoTasks(t *testing.T) {
	t.Run("should return true when day has no tasks", func(t *testing.T) {
		day := app.NewDay(time.Now())
		result := day.HasNoTasks()
		if !result {
			t.Errorf("expected day to have no tasks, but it has %d", len(day.Tasks))
		}
	})
}

func TestSameDay(t *testing.T) {
	testCases := []struct {
		name     string
		date1    time.Time
		date2    time.Time
		expected bool
	}{
		{"dates at exact same time should return true",
			time.Date(2024, time.February, 21, 10, 0, 0, 0, time.Local),
			time.Date(2024, time.February, 21, 10, 0, 0, 0, time.Local),
			true},
		{"dates at same day should return true",
			time.Date(2024, time.February, 21, 10, 0, 0, 0, time.Local),
			time.Date(2024, time.February, 21, 15, 22, 31, 0, time.Local),
			true},
		{"dates at different days should return false",
			time.Date(2024, time.February, 21, 10, 0, 0, 0, time.Local),
			time.Date(2024, time.February, 12, 10, 0, 0, 0, time.Local),
			false},
		{"dates at different month should return false",
			time.Date(2024, time.February, 21, 10, 0, 0, 0, time.Local),
			time.Date(2024, time.December, 21, 10, 0, 0, 0, time.Local),
			false},
		{"dates at different years should return false",
			time.Date(2024, time.February, 21, 10, 0, 0, 0, time.Local),
			time.Date(2025, time.February, 21, 10, 0, 0, 0, time.Local),
			false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := app.IsSameDay(tc.date1, tc.date2)
			if result != tc.expected {
				t.Errorf("expected %t, but got %t", tc.expected, result)
			}
		})
	}
}
