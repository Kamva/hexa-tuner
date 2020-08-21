package huner

import (
	"errors"
	"github.com/kamva/hexa/hconf"
	"github.com/kamva/tracer"
	"github.com/spf13/viper"
	"strings"
)

func NewViper(conf interface{}, envPrefix string, files []string) error {
	v := viper.New()

	if len(files) == 0 {
		return tracer.Trace(errors.New("at least one config files should be exists"))
	}

	isFirst := true
	for _, f := range files {
		v.SetConfigFile(f)

		if isFirst {
			isFirst = false
			if err := v.ReadInConfig(); err != nil {
				return tracer.Trace(err)
			}
			continue
		}

		if err := v.MergeInConfig(); err != nil {
			return tracer.Trace(err)
		}
	}

	v.AllowEmptyEnv(true)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	vd := hconf.NewViperDriver(v)
	return tracer.Trace(vd.Unmarshal(conf))
}
