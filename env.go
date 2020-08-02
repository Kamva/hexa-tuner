package huner

import (
	"github.com/kamva/gutil"
	"github.com/kamva/tracer"
	"github.com/spf13/viper"
)

func bindStructTagsValueAsEnv(v *viper.Viper, configObject interface{}, tagName string) error {
	tagsValue, err := gutil.StructTags(configObject)
	if err != nil {
		return tracer.Trace(err)
	}

	for _, tag := range tagsValue {
		tagVal := tag.Get(tagName)

		if tagVal == "" {
			continue
		}

		if err := v.BindEnv(tagVal); err != nil {
			return tracer.Trace(err)
		}
	}

	return nil
}

// defaultViper make new instance of viper with viperDriver to read from .env file.
// file is absolute string path of file. e.g `gutil.SourcePath()+"../.env"`
func EnvViper(configStruct interface{}, envPrefix string, file string) (*viper.Viper, error) {
	v := viper.New()

	v.AutomaticEnv()

	v.SetEnvPrefix(envPrefix)

	if err := bindStructTagsValueAsEnv(v, configStruct, "mapstructure"); err != nil {
		return nil, tracer.Trace(err)
	}

	// Read config file.
	v.SetConfigFile(file)
	v.SetConfigType("env")
	err := v.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); err != nil && !ok {
		return nil, tracer.Trace(err)
	}

	return v, nil
}
