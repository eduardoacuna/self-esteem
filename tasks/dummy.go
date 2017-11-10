package tasks

import (
	"github.com/eduardoacuna/self-esteem/estimates"
)

var dummyTaskData = map[string][]*Task{
	"0": []*Task{
		&Task{
			ID:         "0",
			UserID:     "0",
			Title:      "Add something",
			Estimation: 2,
			Done:       true,
			Outcome:    estimates.Negative,
		},
	},
	"1": []*Task{
		&Task{
			ID:         "1",
			UserID:     "1",
			Title:      "Refactor something",
			Estimation: 4,
			Done:       true,
			Outcome:    estimates.Positive,
		},
		&Task{
			ID:         "3",
			UserID:     "1",
			Title:      "Remove something",
			Estimation: 1,
			Done:       false,
			Outcome:    estimates.Pending,
		},
	},
	"2": []*Task{
		&Task{
			ID:         "2",
			UserID:     "2",
			Title:      "Fix something",
			Estimation: 3,
			Done:       true,
			Outcome:    estimates.Positive,
		},
	},
}
