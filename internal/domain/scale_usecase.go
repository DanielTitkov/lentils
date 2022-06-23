package domain

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/DanielTitkov/lentils/internal/util"
)

type scaleResolveFunc func(*Scale, *Norm) (*ScaleResult, error)

func (s *Scale) CalculateResult() error {
	resolveFunc, err := getScaleResolveFunc(s.Type)
	if err != nil {
		return err
	}

	start := time.Now()
	// TODO: select norm
	res, err := resolveFunc(s, nil)
	if err != nil {
		return err
	}

	meta := util.NewMeta()
	meta["timestamp"] = time.Now().UnixNano()
	res.Elaplsed = time.Since(start)
	res.Meta = meta

	s.Result = res

	return nil
}

func getScaleResolveFunc(typ string) (scaleResolveFunc, error) {
	switch typ {
	case ScaleTypeSum:
		return resolveScaleSum, nil
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

	formula := fmt.Sprintf("Raw(sum)=%s=%.3f", strings.Join(vals, "+"), sum)

	return &ScaleResult{
		Score:    sum,
		RawScore: sum,
		Min:      0,
		Max:      max,
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
	meanRes, err := resolveScaleMean(s, norm)
	if err != nil {
		return nil, err
	}

	var mean, sigma float64
	var normName string
	if norm != nil {
		normName = norm.ID.String()
		mean = norm.Mean
		sigma = norm.Sigma
	} else {
		normName = "theoretical"
		mean = meanRes.Max / 2  // theoretical mean value
		sigma = meanRes.Max / 5 // split scale to five parts // FIXME
	}

	z := (meanRes.Score - mean) / sigma
	usedNorm := fmt.Sprintf("used norm: %s (M=%.3f S=%.3f)", normName, mean, sigma)
	formula := fmt.Sprintf("%s; z=(%.3f-%.3f)/%.3f=%.3f; %s", meanRes.Formula, meanRes.Score, mean, sigma, z, usedNorm)

	return &ScaleResult{
		Score:    z,
		RawScore: meanRes.RawScore,
		Min:      -99,
		Max:      99,
		Formula:  formula,
	}, nil
}

func resolveScaleSten(s *Scale, norm *Norm) (*ScaleResult, error) {

	return &ScaleResult{
		Score:   0,
		Min:     1,
		Max:     10,
		Formula: "",
	}, nil
}