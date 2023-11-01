package config

type ApplicationConfig struct {
	Name string `toml:"name"`
	Port int    `toml:"port"`
}
