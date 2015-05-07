package configuration

type PsqlConfiguration struct {
	Hostname     string `json:"host"`
	Port         int    `json:"port"`
	DatabaseName string `json:"database"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	SslMode      string `json:"ssl_mode"`
	ConnTimeout  int    `json:"conn_timeout"`
}
