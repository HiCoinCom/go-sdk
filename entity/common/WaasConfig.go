package common

type WaasConfig struct {
	AppId          string `json:"app_id"`
	UserPrivateKey string `json:"user_private_key"`
	WaasPublickKey string `json:"waas_publick_key"`
	Domain         string `json:"domain"`
	Version        string `json:"version"`
	Charset        string `json:"charset"`
	EnableLog      bool   `json:"enable_log"`
}
