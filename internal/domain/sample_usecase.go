package domain

func (s *Sample) NormRank(n *Norm) int {
	var rank int
	if n.Base < NormMinBase {
		return 0
	}

	if n.Base > NormOKBase {
		rank++
	}

	if n.Base > NormGoodBase {
		rank++
	}

	if s.Criteria.Locale != "" {
		rank++
	}

	if s.Criteria.NotSuspicious {
		rank += 2 // so that suspicious results with bigger samples won't intervene
	}

	return rank
}
