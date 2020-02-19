package ktuner

import (
	"fmt"
	"github.com/Kamva/kitty"
	"github.com/Kamva/kitty/kittytranslator"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"strings"
)

const (
	filesKey         = "translate.files"
	fallbackLangsKey = "translate.fallback.langs"
)

// return new translator service.
func NewTranslator(pathPrefix string, config kitty.Config) kitty.Translator {
	files := config.GetList(filesKey)
	fallbackLangs := config.GetList(fallbackLangsKey)

	defaultLang := language.English

	if len(fallbackLangs) >= 1 {
		defaultLang = language.MustParse(fallbackLangs[0])
	}

	bundle := i18n.NewBundle(defaultLang)

	loadLangFiles(bundle, pathPrefix, files)

	localizer := i18n.NewLocalizer(bundle, fallbackLangs...)

	return kittytranslator.NewI18nDriver(bundle, localizer)
}

func loadLangFiles(bundle *i18n.Bundle, prefix string, files []string) {
	for _, file := range files {
		f := fmt.Sprintf("%s/%s.toml", strings.TrimRight(prefix, "/"), strings.Trim(file, "/"))
		bundle.MustLoadMessageFile(f)
	}
}
