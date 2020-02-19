#### kitty-tuner tune kitty services.

#### Install
```
go get github.com/Kamva/kitty-tuner
```

#### Used config variables:
```
// Translator config variables
translate.files (optional) : []string translation files.
translate.fallback.langs (optional,default:en): []string fallback langues

e.g environtment variable in viper driver of cofig:
TRANSLATE.FILES=en,fa,...
TRANSLATE.FALLBACK.LANGS=en,fa,...
```

#### Todo:
- [ ] Write Tests
- [ ] Add badges to readme.
- [ ] CI 