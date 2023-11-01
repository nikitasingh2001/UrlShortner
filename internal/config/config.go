package config

const (
	// FilePath - relative path to the config directory
	FilePath = "%s/conf/%s"

	// DefaultFilename - Filename format of default config file
	DefaultFilename = "env.default.toml"

	// EnvFilename - Filename format of env specific config file
	EnvFilename = "env.%s.toml"
)

var (
	config AppConfig

	appBasePath string
)

type AppConfig struct {
	AuthUser    AuthConfig        `toml:"auth"`
	Database    DatabaseConfig    `toml:"database"`
	Application ApplicationConfig `toml:"application"`
}

func LoadConfig(basePath string, env string) {

	appBasePath = basePath

	loadConfigFromFile(basePath, DefaultFilename, "")
	loadConfigFromFile(basePath, EnvFilename, env)

}

func GetConfig() AppConfig {
	return config
}
