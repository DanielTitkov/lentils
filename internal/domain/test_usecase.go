package domain

import "sort"

func (t *Test) OrderQuestions() {
	sort.Slice(t.Questions, func(i, j int) bool {
		return t.Questions[i].Order < t.Questions[j].Order
	})
}
