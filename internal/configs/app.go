package configs

const LiveSessionName = "go-live-session"

type AppConfig struct {
	SystemSummarySchedule   string `yaml:"systemSummarySchedule"`
	SystemSummaryTimeout    int    `yaml:"systemSummaryTimeout"`
	UpdateNormsSchedule     string `yaml:"updateNormsSchedule"`
	UpdateNormsTimeout      int    `yaml:"updateNormsTimeout"`
	UpdateMarksSchedule     string `yaml:"updateMarksSchedule"`
	UpdateMarksTimeout      int    `yaml:"updateMarksTimeout"`
	UpdateDurationsSchedule string `yaml:"updateDurationsSchedule"`
	UpdateDurationsTimeout  int    `yaml:"updateDurationsTimeout"`
	DefaultTimeLayout       string `yaml:"defaultTimeLayout"`
	Version                 string `yaml:"version"`
}
