package huner

import (
	"errors"
	"fmt"
	"github.com/kamva/gutil"
	"github.com/kamva/hexa"
	"github.com/kamva/hexa/hconf"
	"github.com/kamva/hexa/hlog"
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
	FileName      string // e.g., config
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

// GetConfigFilePaths generates config path as follow:
// - /etc/{project}/{configFile.configExtension}
// - /etc/{project}/{microservice.configExtension}
// - /etc/{project_root_path}/{configFile.configExtension}
// - /etc/{project_root_path}/.env
// - /etc/{project_root_path}/.{environment}.env
func GetConfigFilePaths(o ConfigFilePahtsOpts) []string {
	configFile := fmt.Sprintf("%s.%s", o.FileName, o.FileExtension)
	msConfigFile := fmt.Sprintf("%s.%s", o.Microservice, o.FileExtension)

	files := []string{
		path.Join("/etc", o.Project, configFile),
		path.Join("/etc", o.Project, msConfigFile),
		path.Join(o.ProjectRoot, configFile),
		path.Join(o.ProjectRoot, ".env"),
	}

	if o.Environment != "" {
		files = append(files, path.Join(o.ProjectRoot, fmt.Sprintf(".%s.env", o.Environment)))
	}

	var existedFiles []string

	for _, f := range files {
		if gutil.FileExists(f) {
			existedFiles = append(existedFiles, f)
		}
	}

	hlog.Debug("generated config file paths",
		hlog.Any("available_paths", files),
		hlog.Any("existed_paths", existedFiles),
		hlog.String("config", fmt.Sprintf("%+v", o)),
	)

	return existedFiles
}

// NewViperConfigDriver returns new instance of the viper driver for hexa config
func NewViperConfigDriver(envPrefix string, files []string) (hexa.Config, error) {
	v := viper.New()

	if len(files) == 0 {
		return nil, tracer.Trace(errors.New("at least one config files should be exists"))
	}

	isFirst := true
	for _, f := range files {
		v.SetConfigFile(f)

		if isFirst {
			isFirst = false
			if err := v.ReadInConfig(); err != nil {
				return nil, tracer.Trace(err)
			}
			continue
		}

		if err := v.MergeInConfig(); err != nil {
			return nil, tracer.Trace(err)
		}
	}

	v.AllowEmptyEnv(true)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	return hconf.NewViperDriver(v), nil
}
