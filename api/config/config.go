package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	Port  string `mapstructure:"port"`
	Jwt   Jwt    `mapstructure:"jwt"`
	Neo4j Neo4j  `mapstructure:"neo4j"`
}

type Jwt struct {
	Secret        string `mapstructure:"secret"`
	ExpireMinutes int    `mapstructure:"expire_minutes"`
}

type Neo4j struct {
	Uri      string `mapstructure:"uri"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func LoadConfig() (*Configuration, error) {
	var c Configuration
	v := viper.New()

	v.SetDefault("port", "9011")
	v.SetDefault("jwt.secret", "get_a_random_secret")
	v.SetDefault("jwt.expire_minutes", "5")
	v.SetDefault("neo4j.uri", "bolt://localhost:7687")
	v.SetDefault("neo4j.username", "neo4j")
	v.SetDefault("neo4j.password", "lumiere")
	v.SetDefault("neo4j.database", "neo4j")
	// Set environment variable support:
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetTypeByDefaultValue(true)
	v.SetEnvPrefix("LUMIERE_API")
	v.AutomaticEnv()
	v.ReadInConfig()

	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("viper failed to unmarshal app config: %v", err)
	}
	return &c, nil
}
