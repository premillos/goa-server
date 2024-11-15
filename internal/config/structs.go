package config

type Config struct {
	Server Server `toml:"server"`
	Db     Db     `toml:"db"`
	Oauth  Oauth  `toml:"oauth"`
}

type Server struct {
	Addr string `toml:"addr" json:"addr" default:":8080"`
}

type Db struct {
	Type string `toml:"type" json:"type" default:"mysql"`
	Dsn  string `toml:"dsn" json:"dsn"`
}

type Oauth struct {
	Github Github `toml:"github" json:"github"`
}

type Github struct {
	ClientId     string `toml:"client_id" json:"client_id"`
	ClientSecret string `toml:"client_secret" json:"client_secret"`
	RedirectUrl  string `toml:"redirect_url" json:"redirect_url"`
	Scopes       string `toml:"scopes" json:"scopes"`
}
