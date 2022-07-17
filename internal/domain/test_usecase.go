package domain

import (
	"fmt"
	"math"
	"math/rand"
	"net/url"
	"reflect"
	"sort"
	"time"
	"unicode/utf8"

	"github.com/montanaflynn/stats"
)

func (t *Test) OrderQuestions(seed int64) {
	if t.Display.RandomizeOrder {
		rand.Seed(seed)
		rand.Shuffle(len(t.Questions), func(i, j int) {
			t.Questions[i], t.Questions[j] = t.Questions[j], t.Questions[i]
		})
	} else {
		sort.Slice(t.Questions, func(i, j int) bool {
			return t.Questions[i].Order < t.Questions[j].Order
		})
	}
}

func (t *Test) PageCount() int {
	if t.Display.QuestionsPerPage < 1 {
		return 1
	}
	return int(math.Ceil(float64(len(t.Questions)) / float64(t.Display.QuestionsPerPage)))
}

func (t *Test) GetItem(code string) *Item {
	// test wont have too many items
	// so map is not required probably
	// TODO: but it'd better to check
	for _, q := range t.Questions {
		for _, i := range q.Items {
			if i.Code == code {
				return i
			}
		}
	}

	return nil
}

func (t *Test) QuestionsForPage(page int) []*Question {
	if page < 1 {
		return []*Question{}
	}

	if page > t.PageCount() {
		return []*Question{}
	}

	// TODO: maybe pregenerate pages
	page = page - 1 // in the app page count goes from 1
	var pages [][]*Question
	for i := 0; i < t.PageCount(); i++ {
		var pq []*Question
		for j := 0; j < t.Display.QuestionsPerPage; j++ {
			idx := i*t.Display.QuestionsPerPage + j
			if idx >= len(t.Questions) {
				break
			}
			pq = append(pq, t.Questions[idx])
		}
		pages = append(pages, pq)
	}

	return pages[page]
}

func (t *Test) IsPageDone(page int) bool {
	for _, q := range t.QuestionsForPage(page) {
		if !q.IsDone() {
			return false
		}
	}

	return true
}

func (t *Test) IsPageNotDone(page int) bool {
	return !t.IsPageDone(page)
}

func (t *Test) IsDone() bool {
	for _, q := range t.Questions {
		if !q.IsDone() {
			return false
		}
	}

	return true
}

func (t *Test) IsNotDone() bool {
	return !t.IsDone()
}

func (t *Test) CalculateResult() error {
	for _, s := range t.Scales {
		err := s.CalculateResult()
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Test) DefaultDuration() time.Duration {
	return (time.Duration(len(t.Questions)) * time.Second * 4).Truncate(5 * time.Second)
}

func (t *Test) CalculateDuration() time.Duration {
	if len(t.Takes) < NormMinBase {
		return t.DefaultDuration()
	}

	var vals []float64
	for _, take := range t.Takes {
		vals = append(vals, float64(take.Elapsed()))
	}
	duration, err := stats.Median(vals)
	if err != nil {
		return t.DefaultDuration()
	}

	return time.Duration(duration).Truncate(5 * time.Second)
}

func (t *Test) CalculateMark() float64 {
	if len(t.Takes) == 0 {
		return t.Mark
	}

	var takeMarks []float64
	for _, take := range t.Takes {
		if take.Mark == nil {
			continue
		}
		takeMarks = append(takeMarks, float64(*take.Mark))
	}

	if len(takeMarks) == 0 {
		return t.Mark
	}

	mark, err := stats.Mean(takeMarks)
	if err != nil {
		return t.Mark
	}
	return mark
}

func (t *CreateTestArgs) ValidateTranslations() error {
	locs := make(map[string]struct{})
	for _, l := range t.AvailableLocales {
		locs[l] = struct{}{}
	}

	// check test translations
	transLocs := make(map[string]struct{})
	for _, trans := range t.Translations {
		transLocs[trans.Locale] = struct{}{}
	}

	if eq := reflect.DeepEqual(locs, transLocs); !eq {
		return fmt.Errorf("test %s'' translations are not full: expected %v but got %v", t.Code, locs, transLocs)
	}

	// check scales
	for _, scale := range t.Scales {
		scaleLocs := make(map[string]struct{})
		for _, trans := range scale.Translations {
			scaleLocs[trans.Locale] = struct{}{}
		}
		if eq := reflect.DeepEqual(locs, scaleLocs); !eq {
			return fmt.Errorf("scale '%s' translations are not full: expected %v but got %v", scale.Code, locs, scaleLocs)
		}

		// check interpretations
		for _, interp := range scale.Interpretations {
			interpLocs := make(map[string]struct{})
			for _, trans := range interp.Translations {
				interpLocs[trans.Locale] = struct{}{}
			}
			if eq := reflect.DeepEqual(locs, interpLocs); !eq {
				return fmt.Errorf("interpretation '%v' translations are not full: expected %v but got %v", interp.Range, locs, interpLocs)
			}
		}

		// check items
		for _, itm := range scale.Items {
			itemLocs := make(map[string]struct{})
			for _, trans := range itm.Translations {
				itemLocs[trans.Locale] = struct{}{}
			}
			if eq := reflect.DeepEqual(locs, itemLocs); !eq {
				return fmt.Errorf("item '%s' translations are not full: expected %v but got %v", itm.Code, locs, itemLocs)
			}
		}
	}

	// check questions
	for _, q := range t.Questions {
		questionLocs := make(map[string]struct{})
		for _, trans := range q.Translations {
			questionLocs[trans.Locale] = struct{}{}
		}
		if eq := reflect.DeepEqual(locs, questionLocs); !eq {
			return fmt.Errorf("question '%s' translations are not full: expected %v but got %v", q.Code, locs, questionLocs)
		}
	}

	return nil
}

func (t *Test) Link(domain string) string {
	var params string
	if t.Locale != DefaultLocale() {
		params = fmt.Sprintf("?locale=%s", t.Locale)
	}
	url := fmt.Sprintf("/test/%s%s", t.Code, params)
	if domain == "" {
		return url
	}
	return fmt.Sprintf("https://%s%s", domain, url)
}

func (t *Test) LinkSafe(domain string) string {
	return url.QueryEscape(t.Link(domain))
}

func (t *Test) ResultShareText() string {
	res := fmt.Sprintf("%s - my result:\n\n", t.Title)
	for _, s := range t.Scales {
		if s.Result == nil {
			continue
		}
		if utf8.RuneCountInString(res) > 215 {
			// otherwise it won't fit in a tweet
			res += "\n..."
			break
		}
		res += s.ResultShareText()
	}
	res += "\nExplore"
	return res
}
