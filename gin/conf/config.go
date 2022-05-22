package conf

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Mode               string
	Healthz            bool
	InsecureServerInfo *InsecureServerInfo
	SecureServerInfo   *SecureServerInfo
	//Jwt             	*JwtInfo
	Middlewares []string
	//EnableProfiling 	bool
	//EnableMetrics   	bool
}

type InsecureServerInfo struct {
	Address string
}

type SecureServerInfo struct {
	Address  string
	CertFile string
	KeyFile  string
}

func NewConfigFromOptions(opts *Options) (*Config, error) {
	cfg := &Config{
		Mode:        gin.ReleaseMode,
		Healthz:     true,
		Middlewares: []string{},
		//EnableProfiling: true,
		//EnableMetrics:   true,
		InsecureServerInfo: &InsecureServerInfo{},
		SecureServerInfo:   &SecureServerInfo{},
		/*		Jwt: &JwtInfo{
				Realm:      "iam jwt",
				Timeout:    1 * time.Hour,
				MaxRefresh: 1 * time.Hour,
			},*/
	}
	cfg.Mode = opts.Server.Mode
	cfg.Healthz = opts.Server.Healthz
	cfg.Middlewares = opts.Server.Middlewares
	cfg.InsecureServerInfo.Address = opts.Insecure.Address + ":" + opts.Insecure.Port
	cfg.SecureServerInfo.Address = opts.Secure.Address + ":" + opts.Secure.Port
	cfg.SecureServerInfo.CertFile = opts.Secure.TLS.CertFile
	cfg.SecureServerInfo.KeyFile = opts.Secure.TLS.PrivateKeyFile
	return cfg, nil
}
