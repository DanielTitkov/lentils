package configs

type (
	DataConfig struct {
		Presets PresetsConfig
	}
	PresetsConfig struct {
		UserPresetsPaths []string `yaml:"userPresetsPaths"`
		TestPresetsPaths []string `yaml:"testPresetsPaths"`
	}
)
