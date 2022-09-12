package main

import (
	"fmt"

	"github.com/kamva/gutil"
	huner "github.com/kamva/hexa-tuner"
)

type Config struct {
	MyName   string   `json:"my_name"`
	LogStack []string `json:"log_stack"`
	Debug    bool     `json:"debug"`
	Port     int      `json:"port"`
	//EnableCache bool     `json:"enable_cache"`
}

func main() {
	// Important note: viper parse .env files using env parser, that
	// parser parse every value as "string". viper does not accept two
	// config key with different types (e.g., a key with bool value in config.yaml
	// and with string value in .env file), so if we want to overwrite a bool, int or float
	// values in ".env" file, we should set that value in our json config file
	// as string or remove it from our json config files and just provide it to
	// the .env files (this is just for .env files, not real Environment variables).
	files := huner.ConfigFilePaths(huner.ConfigFilePathsOptions{
		EtcPath:       huner.EtcPath("conf_example"),
		HomePath:      gutil.SourcePath(),
		FileName:      "config",
		FileExtension: "json",
		Environment:   "example",
	}, false)
	v, err := huner.NewViperConfig(huner.EnvKeysPrefix(), files)
	gutil.PanicErr(err)
	cfg := &Config{}
	gutil.PanicErr(v.Unmarshal(&cfg, huner.ViperDecoderTagName("json")))

	fmt.Printf("\n\nconfig:\n%+v\n", cfg)
}
