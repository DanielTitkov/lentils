package handler

import (
	"github.com/tinygodsdev/orrery/internal/domain"
)

type (
	UITranslation struct {
		Base   *UITransBase
		Home   *UITransHome
		Test   *UITransTest
		Result *UITransResult
		// 404
		// about
		// profile
		// result
		// admin
		// privacy
		// terms
	}

	UITransBase struct {
	}

	UITransResult struct {
		AdvancedSettings     string
		ResultsAreShown      string
		SetMethodDescription string
		Recalculate          string
		InStandartScale      string
		InPercentageScale    string
		InDefaultScale       string
		Discard              string
	}

	UITransHome struct {
		Title       string
		Description string
		TagThemes   string
		TagSize     string
		TagFeatures string
		TagAll      string
		TagAny      string
		TagFilter1  string
		TagFilter2  string
	}

	UITransTest struct {
		SelectLoc         string
		InterfaceSettings string
		Instructions      string
		About             string
		BeginLabel        string
		BeginButton       string
		ShowDetails       string
		HideDetails       string
		AutoNext          string
		AutoNextInfo      string
		ShowInstruction   string
		HideInstruction   string
		PrevButton        string
		NextButton        string
		FinishButton      string
		Page              string
		Of                string
		ResultLinkLabel   string
		ResultLinkInfo    string
		TakeThisButton    string
		ExploreButton     string
		DetailedReport    string
		Disclaimer        string
		NormsInfo         string
		TestMark          string
	}
)

func newUITranslation(locale string) *UITranslation {
	switch locale {
	case domain.LocaleEn:
		return newTranslationEn()
	case domain.LocaleRu:
		return newTranslationRu()
	default:
		return newUITranslation(domain.DefaultLocale())
	}
}

func newTranslationEn() *UITranslation {
	return &UITranslation{
		Base: &UITransBase{},
		Home: &UITransHome{
			Title:       "Orrery",
			Description: "Modern psychometrics for fun and science",
			TagThemes:   "Themes",
			TagSize:     "Size",
			TagFeatures: "Features",
			TagFilter1:  "Tests with",
			TagAll:      "ALL",
			TagAny:      "ANY",
			TagFilter2:  "given tags",
		},
		Test: &UITransTest{
			SelectLoc:         "Select language",
			InterfaceSettings: "Interface settings",
			Instructions:      "Instructions",
			About:             "About this test",
			BeginLabel:        "When you are happy please continue to the next step.",
			BeginButton:       "Begin",
			ShowDetails:       "Show test details",
			HideDetails:       "Hide test details",
			AutoNext:          "Auto-Next",
			AutoNextInfo: `If Auto-Next is on, 
			test will switch to the next page as soon as all the questions on a page are answered. 
			Otherwise you would need to click "Next" button.`,
			ShowInstruction: "Show test instruction",
			HideInstruction: "Hide test instruction",
			NextButton:      "Next",
			PrevButton:      "Previous",
			FinishButton:    "Finish",
			Page:            "Page",
			Of:              "of",
			ResultLinkLabel: "Result permanent link",
			ResultLinkInfo:  "NB: anyone with this link will be able to see this page",
			TakeThisButton:  "Take this test",
			ExploreButton:   "Explore other tests",
			DetailedReport:  "Detailed report",
			Disclaimer: `The results of this online quiz (personality test) 
			are provided "as-is" only for educational purposes and should not be construed 
			as providing professional or certified advice of any kind.`,
			NormsInfo: `Orrery uses dynamic standardization, 
			so results may change (probably insignificantly) with time, 
			as we gather more data and recalculate norms.`,
			TestMark: "Please rate how useful/interesting was this test for you?",
		},
		Result: &UITransResult{
			AdvancedSettings: "Advanced settings",
			ResultsAreShown:  "Currently results are shown",
			SetMethodDescription: `By default Orrery calculates results with the method
			that makes most sense for the current test. 
			You can ask the system to recalculate results using different method.
			Use this at your own discretion.`,
			Recalculate:       "Recalculate",
			InStandartScale:   "in standart scale",
			InPercentageScale: "in maximum percentage",
			InDefaultScale:    "in default scale",
			Discard:           "discard",
		},
	}
}

func newTranslationRu() *UITranslation {
	return &UITranslation{
		Base: &UITransBase{},
		Home: &UITransHome{
			Title:       "Orrery",
			Description: "Современная психометрика для науки и потехи",
			TagThemes:   "Темы",
			TagSize:     "Длина",
			TagFeatures: "Особенности",
			TagFilter1:  "Тесты",
			TagAll:      "со всеми",
			TagAny:      "с любыми",
			TagFilter2:  "тегами",
		},
		Test: &UITransTest{
			SelectLoc:         "Выбрать язык",
			InterfaceSettings: "Настройки интерфейса",
			Instructions:      "Инструкция",
			About:             "Информация о тесте",
			BeginLabel:        "Когда будете готовы, нажмите, чтобы продолжить.",
			BeginButton:       "Начать",
			ShowDetails:       "Показать информацию о тесте",
			HideDetails:       "Скрыть информацию о тесте",
			AutoNext:          "Auto-Next",
			AutoNextInfo: `Если функция Auto-Next включена, 
			тест автоматически перейдёт на следующую страницу,
			как только все вопросы на активной странице будут отвечены. 
			В противном случае нужно будет нажать на кнопку "Далее"`,
			ShowInstruction: "Показать инструкцию",
			HideInstruction: "Скрыть инструкцию",
			NextButton:      "Далее",
			PrevButton:      "Назад",
			FinishButton:    "Завершить",
			Page:            "Страница",
			Of:              "из",
			ResultLinkLabel: "Постоянная ссылка на результат",
			ResultLinkInfo:  "Внимание: кто угодно сможет посмотреть результат по этой ссылке",
			TakeThisButton:  "Пройти этот тест",
			ExploreButton:   "Посмотреть другие тесты",
			DetailedReport:  "Подробный отчёт",
			Disclaimer: `Результаты этого опросника (личностного теста)
			предоставляются как есть исключительно в образовательных целях
			и не должны рассматриваться как рекомендация или заключение профессионала.`,
			NormsInfo: `Orrery использует динамическую стандартизацию, 
			поэтому в будущем результаты могут меняться (скорее всего незначительно), 
			по мере того как мы собираем больше данных и обновляем нормы.`,
			TestMark: "Насколько полезным/интересным оказался для вас этот тест?",
		},
		Result: &UITransResult{
			AdvancedSettings: "Профессиональные опции",
			ResultsAreShown:  "Сейчас результаты отображаются",
			SetMethodDescription: `По умолчанию Orrery рассчитывает результаты тем методом, 
			который имеет больше всего смысла для данного теста.
			Вы можете попросить систему пересчитать результаты, используя другой метод.
			Пользуйтесь этим на своё усмотрение.`,
			Recalculate:       "Пересчитать",
			InStandartScale:   "в стандартной шкале",
			InPercentageScale: "в процентах от максимума",
			InDefaultScale:    "в шкале по умолчанию",
			Discard:           "сбросить",
		},
	}
}
