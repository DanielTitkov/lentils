package email

import (
	sendinblue "github.com/sendinblue/APIv3-go-library/lib"
)

type Service struct {
	Client *sendinblue.APIClient
}

func New(apiKey string) (*Service, error) {
	cfg := sendinblue.NewConfiguration()
	cfg.AddDefaultHeader("api-key", apiKey)
	sib := sendinblue.NewAPIClient(cfg)

	return &Service{
		Client: sib,
	}, nil
}
