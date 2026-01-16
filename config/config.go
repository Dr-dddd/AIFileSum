package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	Env  string     `mapstructure:"env"`
	Qwen QwenConfig `mapstructure:"qwen"`
}

type QwenConfig struct {
	APIKey  string `mapstructure:"api_key"`
	BaseURL string `mapstructure:"base_url"`
}

// Get 返回全局唯一的配置实例（线程安全）
func Get() *Config {
	once.Do(func() {
		cfg, err := loadConfig()
		if err != nil {
			panic(fmt.Errorf("failed to load config: %w", err))
		}
		instance = cfg
	})
	return instance
}

// loadConfig 是内部加载逻辑（不暴露给外部）
func loadConfig() (*Config, error) {
	// 从config.yaml 读取配置

	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config error: %w", err)
	}

	// 校验必要字段
	if cfg.Qwen.APIKey == "" {
		return nil, fmt.Errorf("missing required config: qwen.api_key")
	}

	return &cfg, nil
}
