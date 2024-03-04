package minimax

import "net/http"

const (
	APIV1                    = "https://api.minimax.chat/v1"
	defaultEmptyMessageLimit = 300
)

type Config struct {
	apiToken string
	groupID  string

	BaseURL           string
	HTTPClient        *http.Client
	EmptyMessageLimit uint
}

func DefaultConfig(apiToken, groupID string) *Config {
	return &Config{
		apiToken:          apiToken,
		groupID:           groupID,
		BaseURL:           APIV1,
		HTTPClient:        &http.Client{},
		EmptyMessageLimit: defaultEmptyMessageLimit,
	}
}
