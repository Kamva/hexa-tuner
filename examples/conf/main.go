package main

import (
	"fmt"
	"github.com/kamva/gutil"
	huner "github.com/kamva/hexa-tuner"
)

type Config struct {
	MyName string `json:"my_name" mapstructure:"my_name"`
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
	v, err := huner.NewViperConfigDriver("MM", files)
	gutil.PanicErr(err)
	cfg := &Config{}
	gutil.PanicErr(v.Unmarshal(&cfg))

	fmt.Printf("%+v", cfg)
}
