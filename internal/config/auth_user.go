package config

type AuthConfig struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
}
