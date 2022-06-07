package configs

type (
	AuthConfig struct {
		Secret  string
		Exp     int
		Google  Google
		Github  Github
		Twitter Twitter
	}
	Google struct {
		Client   string
		Secret   string
		Callback string
	}
	Github struct {
		Client   string
		Secret   string
		Callback string
	}
	Twitter struct {
		Client   string
		Secret   string
		Callback string
	}
)
