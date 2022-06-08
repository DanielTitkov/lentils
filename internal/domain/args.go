package domain

type (
	CreateTestArgs struct {
		Code         string
		Published    bool
		Generate     GenerateQuestionsArgs
		Translations []TestTranslation
		Questions    []CreateQuestionArgs
		Scales       []CreateScaleArgs
	}

	GenerateQuestionsArgs struct {
		Method   string
		Template CreateQuestionArgs
	}

	TestTranslation struct {
		Locale      string
		Title       string
		Description string
		Instruction string
	}

	CreateQuestionArgs struct {
		Type         string
		Translations []QuestionTranslation
	}

	QuestionTranslation struct {
		Locale        string
		Content       string
		HeaderContent string
		FooterConent  string
	}

	CreateScaleArgs struct {
		Type         string
		Code         string
		Translations []ScaleTranslation
		Items        []CreateItemArgs
	}

	ScaleTranslation struct {
		Locale      string
		Title       string
		Description string
	}

	CreateItemArgs struct {
		Code         string
		Steps        int
		Reverse      bool
		Translations []ItemTranslation
	}

	ItemTranslation struct {
		Locale  string
		Content string
	}
)
