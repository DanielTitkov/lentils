package configs

type (
	ExternalConfig struct {
		Telegram   TelegramConfig
		Sendinblue SendinblueConfig `yaml:"sendinblue"`
	}
	TelegramConfig struct {
		TelegramTo    int64  `yaml:"telegramTo"`
		TelegramToken string `yaml:"telegramToken"`
	}
	SendinblueConfig struct {
		Key string `yaml:"key"`
	}
)
