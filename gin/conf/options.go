package conf

import (
	"encoding/json"
	"os"

	"github.com/rebirthmonkey/pkg/log"
	"github.com/spf13/viper"
)

type Options struct {
	Server struct {
		Mode        string   `json:"mode"        mapstructure:"mode"`
		Healthz     bool     `json:"healthz"     mapstructure:"healthz"`
		Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
	}
	Insecure struct {
		Address string `mapstructure:"bind-address"`
		Port    string `mapstructure:"bind-port"`
	}
	Secure struct {
		Address string `mapstructure:"bind-address"`
		Port    string `mapstructure:"bind-port"`
		TLS     struct {
			CertFile       string `mapstructure:"cert-file"`
			PrivateKeyFile string `mapstructure:"private-key-file"`
		}
	}
}

func NewOptions(optionFile string) *Options {
	o := &Options{}
	viper.SetConfigFile(optionFile)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read configuration file(%s): %+v\n", optionFile, err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&o); err != nil {
		log.Fatalf("unable to decode into struct: %+v \n", err)
		os.Exit(1)
	}
	return o
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
