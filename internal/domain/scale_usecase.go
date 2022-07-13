package domain

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/DanielTitkov/orrery/internal/util"
)

type scaleResolveFunc func(*Scale, *Norm) (*ScaleResult, error)

func (s *Scale) CalculateResult() error {
	resolveFunc, err := getScaleResolveFunc(s.Type)
	if err != nil {
		return err
	}

	start := time.Now()
	// TODO: select norm
	res, err := resolveFunc(s, s.Norm)
	if err != nil {
		return err
	}

	meta := util.NewMeta()
	meta["timestamp"] = time.Now().UnixNano()
	res.Elaplsed = time.Since(start)
	res.Meta = meta
	res.Interpretation = s.chooseInterpretation(res.Score)

	s.Result = res

	return nil
}

func (s *Scale) chooseInterpretation(score float64) *Interpretation {
	for _, interp := range s.Interpretations {
		// FIXME this can be umbigious in case of random order
		if score >= interp.Range[0] && score <= interp.Range[1] {
			return interp
		}
	}

	return nil
}

func getScaleResolveFunc(typ string) (scaleResolveFunc, error) {
	switch typ {
	case ScaleTypeSum:
		return resolveScaleSum, nil
	case ScaleTypePerc:
		return resolveScalePerc, nil
	case ScaleTypeMean:
		return resolveScaleMean, nil
	case ScaleTypeSten:
		return resolveScaleSten, nil
	case ScaleTypeZScore:
		return resolveScaleZScore, nil
	default:
		return nil, fmt.Errorf("got unknown scale type: %s", typ)
	}
}

func resolveScaleSum(s *Scale, norm *Norm) (*ScaleResult, error) {
	var sum, max float64
	var vals []string
	for _, itm := range s.Items {
		if itm.Response == nil {
			continue
		}
		if itm.Reverse {
			sum += float64(itm.Steps - 1 - itm.Response.Value)
			vals = append(vals, fmt.Sprintf("(%d-%d)", itm.Steps-1, itm.Response.Value))
		} else {
			sum += float64(itm.Response.Value)
			vals = append(vals, strconv.Itoa(itm.Response.Value))
		}
		max += float64(itm.Steps - 1)
	}

	formula := fmt.Sprintf("Raw(sum)=%s=%.1f", strings.Join(vals, "+"), sum)

	return &ScaleResult{
		Score:    sum,
		RawScore: sum,
		Min:      0,
		Max:      max,
		Formula:  formula,
	}, nil
}

func resolveScalePerc(s *Scale, norm *Norm) (*ScaleResult, error) {
	sumRes, err := resolveScaleSum(s, norm)
	if err != nil {
		return nil, err
	}

	perc := sumRes.Score / sumRes.Max * 100
	formula := fmt.Sprintf("%s; Percentage=%.1f/%.1f*100=%.1f", sumRes.Formula, sumRes.Score, sumRes.Max, perc)

	return &ScaleResult{
		Score:    perc,
		RawScore: sumRes.RawScore,
		Min:      0,
		Max:      100,
		Formula:  formula,
	}, nil
}

func resolveScaleMean(s *Scale, norm *Norm) (*ScaleResult, error) {
	sumRes, err := resolveScaleSum(s, norm)
	if err != nil {
		return nil, err
	}

	mean := sumRes.Score / float64(len(s.Items))
	max := sumRes.Max / float64(len(s.Items))
	formula := fmt.Sprintf("%s; M=%.3f/%d=%.3f", sumRes.Formula, sumRes.Score, len(s.Items), mean)

	return &ScaleResult{
		Score:    mean,
		RawScore: sumRes.RawScore,
		Min:      0,
		Max:      max,
		Formula:  formula,
	}, nil
}

func resolveScaleZScore(s *Scale, norm *Norm) (*ScaleResult, error) {
	sumRes, err := resolveScaleSum(s, norm)
	if err != nil {
		return nil, err
	}

	var mean, sigma float64
	var base int
	var normName string
	if norm != nil && norm.Base >= NormMinBase {
		normName = norm.Name
		mean = norm.Mean
		sigma = norm.Sigma
		base = norm.Base
	} else {
		normName = "theoretical"
		mean = sumRes.Max / 2  // theoretical mean value
		sigma = sumRes.Max / 5 // split scale in five parts // FIXME
		base = 0
	}

	z := (sumRes.Score - mean) / sigma
	usedNorm := fmt.Sprintf("used norm: %s (M=%.2f sd=%.2f, n=%d)", normName, mean, sigma, base)
	formula := fmt.Sprintf("%s; z=(%.3f-%.3f)/%.3f=%.3f; %s", sumRes.Formula, sumRes.Score, mean, sigma, z, usedNorm)

	return &ScaleResult{
		Score:    z,
		RawScore: sumRes.RawScore,
		Min:      -99,
		Max:      99,
		Formula:  formula,
	}, nil
}

func resolveScaleSten(s *Scale, norm *Norm) (*ScaleResult, error) {
	zRes, err := resolveScaleZScore(s, norm)
	if err != nil {
		return nil, err
	}

	sten := (zRes.Score)*2 + 5.5
	restrictedSten := sten
	if sten > 10 {
		restrictedSten = 10
	}
	if sten < 1 {
		restrictedSten = 1
	}
	formula := fmt.Sprintf("%s; Sten(raw)=%.1f; Sten=%.1f", zRes.Formula, sten, restrictedSten)

	return &ScaleResult{
		Score:    restrictedSten,
		RawScore: zRes.RawScore,
		Min:      1,
		Max:      10,
		Formula:  formula,
	}, nil
}

func (s *Scale) ResultShareText() string {
	if s.Result == nil {
		return ""
	}

	// map result to share scale
	shareRes := s.Result.Score / s.Result.Max * ShareScaleLen

	return fmt.Sprintf(
		"%s: %s%s\n",
		s.Abbreviation,
		strings.Repeat(ShareScaleUnit, int(math.Round(shareRes))),
		strings.Repeat(ShareScaleUnitEmpty, ShareScaleLen-int(math.Round(shareRes))),
	)
}
