package huner

import (
	"errors"
	"fmt"
	"github.com/kamva/gutil"
	"github.com/kamva/hexa/hconf"
	"github.com/kamva/tracer"
	"github.com/spf13/viper"
	"os"
	"path"
	"strings"
)

type ConfigFilePahtsOpts struct {
	Project       string // e.g., senna
	Microservice  string // e.g., order
	ProjectRoot   string // e.g., /home/mehran/senna/order
	ConfigFile    string // e.g., config
	FileExtension string // e.g., json or yaml
	Environment   string // (optional) e.g., staging
}

func EnvKeysPrefix() string {
	return os.Getenv("HEXA_CONF_PREFIX")
}

func Environment(prefix string) string {
	key := "ENV"
	if prefix != "" {
		key = fmt.Sprintf("%s_%s", prefix, key)
	}
	return os.Getenv(key)
}

// ConfigFilePaths generates config path as follow:
// - /etc/{project}/{configFile}
// - /etc/{project}/{microservice}/{configFile}
// - /etc/{project_root_path}/{configFile}
// - /etc/{project_root_path}/.env
// - /etc/{project_root_path}/.env.{environment}
func ConfigFilePaths(o ConfigFilePahtsOpts) []string {
	configFile := fmt.Sprintf("%s.%s", o.ConfigFile, o.FileExtension)
	msConfigFile := fmt.Sprintf("%s.%s", o.Microservice, o.FileExtension)

	files := []string{
		path.Join("/etc", o.Project, configFile),
		path.Join("/etc", o.Project, msConfigFile),
		path.Join(o.ProjectRoot, configFile),
		path.Join(o.ProjectRoot, ".env"),
	}

	if o.Environment != "" {
		files = append(files, path.Join(o.ProjectRoot, fmt.Sprintf(".env.%s", o.Environment)))
	}

	var existedFiles []string

	for _, f := range files {
		if gutil.FileExists(f) {
			existedFiles = append(existedFiles, f)
		}
	}
	return existedFiles
}

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
