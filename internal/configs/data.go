package configs

type (
	DataConfig struct {
		Presets PresetsConfig
	}
	PresetsConfig struct {
		ChallengePresetsPaths  []string `yaml:"challengePresetsPaths"`
		UserPresetsPaths       []string `yaml:"userPresetsPaths"`
		PredictionPresetsPaths []string `yaml:"predictionPresetsPaths"`
	}
)
