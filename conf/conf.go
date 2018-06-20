package conf

import (
	"os"
	"strings"
)

func GetEnvParams() map[string]string {
	envParams := map[string]string{
		"harbor_protocol": "https", // http Or https
		"harbor_host":     "",
		"harbor_user":     "",
		"harbor_password": "",
	}

	//如果在环境变量中检测到有envParams中的key,则将envParams中key相关的value替换为环境变量中的值
	//example: APP_ENV=
	for k := range envParams {
		if v := os.Getenv(strings.ToUpper(k)); v != "" {
			envParams[k] = v
		}
	}

	return envParams
}
