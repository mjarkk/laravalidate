package laravalidate

import (
	"strings"

	"golang.org/x/text/language"
)

func lookupLanguages(languages []language.Tag) []string {
	languageStrs := []string{}
outer:
	for _, lang := range languages {
		langStr := strings.ToLower(lang.String())

		langParts := strings.SplitN(langStr, "-", 2)
		if len(langParts) == 1 {
			if len(languageStrs) == 0 {
				languageStrs = append(languageStrs, langStr)
				continue
			}

			for _, existingLang := range languageStrs {
				if existingLang == langStr {
					continue outer
				}
			}

			languageStrs = append(languageStrs, langStr)
		}

		alternativeLang := langParts[0]
		if len(languageStrs) == 0 {
			languageStrs = append(languageStrs, langStr, alternativeLang)
			continue
		}

		addLang := true
		addAlternativeLang := true
		for _, existingLang := range languageStrs {
			if existingLang == langStr {
				addLang = false
			}
			if existingLang == alternativeLang {
				addAlternativeLang = false
			}
			if !addLang && !addAlternativeLang {
				continue outer
			}
		}

		if addLang {
			languageStrs = append(languageStrs, langStr)
		}
		if addAlternativeLang {
			languageStrs = append(languageStrs, alternativeLang)
		}
	}

	if len(languageStrs) == 0 {
		return []string{"en"}
	}

	return languageStrs
}
