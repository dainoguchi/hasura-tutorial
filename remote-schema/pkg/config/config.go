package config

import "github.com/kelseyhightower/envconfig"

type (
	Config struct {
		Env      string `required:"true" envconfig:"ENV" default:"local"`
		Timezone string `required:"true" envconfig:"TZ" default:"Asia/Tokyo"`

		Database Database
	}

	Database struct {
		URL     string `required:"true" envconfig:"DB_URL"`
		TestURL string `envconfig:"TEST_DB_URL"`
	}
)

// 環境変数セットした構造体を返却
func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)

	return cfg, err
}
