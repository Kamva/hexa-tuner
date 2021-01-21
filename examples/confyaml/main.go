package main

import (
	"fmt"
	"github.com/kamva/gutil"
	huner "github.com/kamva/hexa-tuner"
)

type Config struct {
	MyName      string   `json:"my_name" mapstructure:"my_name"`
	LogStack    []string `json:"log_stack" mapstructure:"log_stack"`
	Debug       bool     `json:"debug" mapstructure:"debug"`
	Port        int      `json:"port" mapstructure:"port"`
	EnableCache bool     `json:"enable_cache" mapstructure:"enable_cache"`
}

func main() {
	// Important note: viper parse .env files using env parser, that
	// parser parse every value as "string". viper does not accept two
	// config key with different types (e.g., a key with bool value in config.json
	// and with string value in .env file), so if we want overwrite bool, int or float
	// values in ".env" file, we should set that value in our json config file
	// as string or remove it from our json config files and just provide it to
	// the .env files (this is just for .env files, not real Environment variables).
	files := huner.GetConfigFilePaths(huner.ConfigFilePathsOpts{
		AppName:       "example_conf",
		ServiceName:   "example_ms",
		HomePath:      gutil.SourcePath(),
		FileName:      "config",
		FileExtension: "yaml",
		Environment:   "example",
	})
	v, err := huner.NewViperConfigDriver(huner.EnvKeysPrefix(), files)
	gutil.PanicErr(err)
	cfg := &Config{}
	gutil.PanicErr(v.Unmarshal(&cfg))

	fmt.Printf("\n\nconfig:\n%+v\n", cfg)
}
