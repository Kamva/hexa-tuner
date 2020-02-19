#### kitty-tuner tune kitty services.

#### Install
```
go get github.com/Kamva/kitty-tuner
```

#### Used config variables:
```
translate.files : []string translation files.
translate.fallback.langs: []string fallback langues

e.g environtment variable in viper driver of cofig:
TRANSLATE.FILES=en,fa,...
TRANSLATE.FALLBACK.LANGS=en,fa,...
```


#### Todo:
- [ ] Write Tests
- [ ] Add badges to readme.
- [ ] CI 