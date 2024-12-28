//
// Copyright 2024-2024 The Haora Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package data

import (
	"github.com/drademann/fugo/test"
	"github.com/drademann/haora/app/datetime"
	"github.com/drademann/haora/cmd/config"
	"testing"
	"time"
)

func TestSuggestedFinish_givenNoTimes_shouldReturnNotFound(t *testing.T) {
	week := Week{
		Days: [7]Day{
			{Date: test.Date("30.12.2024 09:00"), Tasks: nil, Finished: time.Time{}},
			{Date: test.Date("31.12.2024 09:00"), Tasks: nil, Finished: time.Time{}},
			{Date: test.Date("01.01.2025 09:00"), Tasks: nil, Finished: time.Time{}},
			{Date: test.Date("02.01.2025 09:00"), Tasks: nil, Finished: time.Time{}},
			{Date: test.Date("03.01.2025 09:00"), Tasks: nil, Finished: time.Time{}},
			{Date: test.Date("04.01.2025 09:00"), Tasks: nil, Finished: time.Time{}},
			{Date: test.Date("05.01.2025 09:00"), Tasks: nil, Finished: time.Time{}},
		},
	}

	suggestion, found := week.SuggestedFinish()

	if found {
		t.Errorf("expected suggested finish to be empty, but got %s", suggestion.Format("02.01.2006"))
	}
}

func TestSuggestedFinish_givenSuggestionGreaterMidnight_shouldReturnNotFound(t *testing.T) {
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)
	config.ApplyConfigOptions(t)
	datetime.AssumeForTestNowAt(t, test.Date("03.01.2025 12:00"))

	week := Week{
		Days: [7]Day{
			{Date: test.Date("30.12.2024 09:00"), Tasks: []*Task{NewTask(test.Date("30.12.2024 10:00"), "task 1")}, Finished: test.Date("30.12.2024 17:00")},
			{Date: test.Date("31.12.2024 09:00"), Tasks: nil, Finished: time.Time{}},
			{Date: test.Date("01.01.2025 09:00"), Tasks: nil, Finished: time.Time{}},
			{Date: test.Date("02.01.2025 09:00"), Tasks: nil, Finished: time.Time{}},
			{Date: test.Date("03.01.2025 09:00"), Tasks: nil, Finished: time.Time{}},
		},
	}

	suggestion, found := week.SuggestedFinish()

	if found {
		t.Errorf("expected suggested finish to be empty, but got %s", suggestion.Format("02.01.2006 15:04"))
	}
}

func TestSuggestedFinish_givenSuggestionEarlierThanMidnight_shouldReturnSuggestion(t *testing.T) {
	config.SetDurationPerWeek(t, 40*time.Hour)
	config.SetDaysPerWeek(t, 5)
	config.ApplyConfigOptions(t)
	datetime.AssumeForTestNowAt(t, test.Date("03.01.2025 12:00"))

	week := Week{
		Days: [7]Day{
			{Date: test.Date("30.12.2024 09:00"), Tasks: []*Task{NewTask(test.Date("30.12.2024 10:00"), "task 1")}, Finished: test.Date("30.12.2024 18:00")},
			{Date: test.Date("31.12.2024 09:00"), Tasks: []*Task{NewTask(test.Date("31.12.2024 10:00"), "task 2")}, Finished: test.Date("31.12.2024 18:00")},
			{Date: test.Date("01.01.2025 09:00"), Tasks: []*Task{NewTask(test.Date("01.01.2025 10:00"), "task 3")}, Finished: test.Date("01.01.2025 18:00")},
			{Date: test.Date("02.01.2025 09:00"), Tasks: []*Task{NewTask(test.Date("02.01.2025 10:00"), "task 4")}, Finished: test.Date("02.01.2025 18:00")},
			{Date: test.Date("03.01.2025 09:00"), Tasks: []*Task{NewTask(test.Date("03.01.2025 08:00"), "task 5")}, Finished: time.Time{}},
		},
	}

	suggestion, found := week.SuggestedFinish()

	if !found {
		t.Fatalf("expected suggested finish to be found, but got empty")
	}
	expected := test.Date("03.01.2025 16:00")
	if suggestion != expected {
		t.Errorf("expected suggested finish to be %s but got %s", expected.Format("02.01.2006 15:04"), suggestion.Format("02.01.2006 15:04"))
	}
}
