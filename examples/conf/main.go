package main

import (
	"fmt"
	"github.com/kamva/gutil"
	huner "github.com/kamva/hexa-tuner"
)

type Config struct {
	MyName   string   `json:"my_name" mapstructure:"my_name"`
	LogStack []string `json:"log_stack" mapstructure:"log_stack"`
}

func main() {
	files := huner.GetConfigFilePaths(huner.ConfigFilePahtsOpts{
		Project:       "example_conf",
		Microservice:  "example_ms",
		ProjectRoot:   gutil.SourcePath(),
		FileName:      "config",
		FileExtension: "json",
		Environment:   "example",
	})
	v, err := huner.NewViperConfigDriver(huner.EnvKeysPrefix(), files)
	gutil.PanicErr(err)
	cfg := &Config{}
	gutil.PanicErr(v.Unmarshal(&cfg))

	fmt.Printf("%+v", cfg)
}
