package configs

const LiveSessionName = "go-live-session"

type AppConfig struct {
	DefaultChallengePageLimit int    `yaml:"defaultChallengePageLimit"`
	HomeChallengePageLimit    int    `yaml:"homeChallengePageLimit"`
	SystemSummaryInterval     int    `yaml:"systemSummaryInterval"`
	SystemSummaryTimeout      int    `yaml:"systemSummaryTimeout"`
	UpdateNormsInterval       int    `yaml:"updateNormsInterval"`
	UpdateNormTimeout         int    `yaml:"updateNormsTimeout"`
	DefaultTimeLayout         string `yaml:"defaultTimeLayout"`
	MinProofCount             int    `yaml:"minProofCount"`
	Version                   string `yaml:"version"`
}
