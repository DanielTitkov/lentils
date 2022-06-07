package domain

func (us *UserSummary) PercCorrect() int {
	return 100 - us.PercIncorrect()
}

func (us *UserSummary) PercIncorrect() int {
	if us.CorrectPredictions+us.IncorrectPredictions == 0 {
		return 0
	}
	return int(float64(us.IncorrectPredictions) / float64(us.CorrectPredictions+us.IncorrectPredictions) * 100.0)
}
