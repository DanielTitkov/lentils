package domain

func (q *Question) IsDone() bool {
	for _, item := range q.Items {
		if item.Response == nil {
			return false
		}
	}

	return true
}
