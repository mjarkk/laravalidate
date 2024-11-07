package translations

import (
	. "github.com/mjarkk/laravalidate"
	"golang.org/x/text/language"
)

func RegisterNlTranslations() {
	RegisterMessages(language.Dutch, map[string]MessageResolver{
		"accepted": BasicMessageResolver("Het :attribute veld moet worden geaccepteerd."),
		// "accepted_if": BasicMessageResolver("Het :attribute veld moet worden geaccepteerd wanneer :other :value is."),
		"active_url":     BasicMessageResolver("Het :attribute veld moet een geldige URL zijn."),
		"after":          BasicMessageResolver("Het :attribute veld moet een datum zijn na :date."),
		"after_or_equal": BasicMessageResolver("Het :attribute veld moet een datum zijn na of gelijk aan :date."),
		"alpha":          BasicMessageResolver("Het :attribute veld mag alleen letters bevatten."),
		"alpha_dash":     BasicMessageResolver("Het :attribute veld mag alleen letters, cijfers, streepjes en onderstrepingstekens bevatten."),
		"alpha_numeric":  BasicMessageResolver("Het :attribute veld mag alleen letters en cijfers bevatten."),
		// "array":     BasicMessageResolver("Het :attribute veld moet een array zijn."),
		"ascii":           BasicMessageResolver("Het :attribute veld mag alleen enkelbyte alfanumerieke tekens en symbolen bevatten."),
		"bail":            BasicMessageResolver("Het :attribute veld moet slagen."),
		"before":          BasicMessageResolver("Het :attribute veld moet een datum zijn voor :date."),
		"before_or_equal": BasicMessageResolver("Het :attribute veld moet een datum zijn voor of gelijk aan :date."),
		"between": MessageHintResolver{
			Fallback: "Het :attribute veld moet tussen :arg0 en :arg1 liggen.",
			Hints: map[string]string{
				"array":   "Het :attribute veld moet tussen :arg0 en :arg1 items bevatten.",
				"file":    "Het :attribute veld moet tussen :arg0 en :arg1 kilobytes zijn.",
				"numeric": "Het :attribute veld moet tussen :arg0 en :arg1 liggen.",
				"string":  "Het :attribute veld moet tussen :arg0 en :arg1 tekens lang zijn.",
			},
		},
		"boolean": BasicMessageResolver("Het :attribute veld moet waar of onwaar zijn."),
		// "can": BasicMessageResolver("Het :attribute veld bevat een niet-toegestane waarde."),
		"confirmed": BasicMessageResolver("De bevestiging van het :attribute veld komt niet overeen."),
		// "contains":  BasicMessageResolver("Het :attribute veld ontbreekt een vereiste waarde."),
		// "current_password": BasicMessageResolver("Het wachtwoord is onjuist."),
		"date": BasicMessageResolver("Het :attribute veld moet een geldige datum zijn."),
		// "date_equals": BasicMessageResolver("Het :attribute veld moet een datum zijn gelijk aan :date."),
		"date_format": BasicMessageResolver("Het :attribute veld moet overeenkomen met het formaat :arg."),
		// "decimal": BasicMessageResolver("Het :attribute veld moet :arg decimalen hebben."),
		"declined": BasicMessageResolver("Het :attribute veld moet worden afgewezen."),
		// "declined_if": BasicMessageResolver("Het :attribute veld moet worden afgewezen wanneer :other :value is."),
		// "different":   BasicMessageResolver("Het :attribute veld en :other moeten verschillend zijn."),
		"digits":         BasicMessageResolver("Het :attribute veld moet :digits cijfers lang zijn."),
		"digits_between": BasicMessageResolver("Het :attribute veld moet tussen :arg0 en :arg1 cijfers lang zijn."),
		// "dimensions":        BasicMessageResolver("Het :attribute veld heeft ongeldige afbeeldingsdimensies."),
		// "distinct":          BasicMessageResolver("Het :attribute veld heeft een dubbele waarde."),
		// "doesnt_end_with":   BasicMessageResolver("Het :attribute veld mag niet eindigen met een van de volgende: :args."),
		// "doesnt_start_with": BasicMessageResolver("Het :attribute veld mag niet beginnen met een van de volgende: :args."),
		"email":     BasicMessageResolver("Het :attribute veld moet een geldig e-mailadres zijn."),
		"ends_with": BasicMessageResolver("Het :attribute veld moet eindigen met een van de volgende: :args."),
		// "enum":    BasicMessageResolver("De geselecteerde :attribute is ongeldig."),
		// "exists":  BasicMessageResolver("De geselecteerde :attribute is ongeldig."),
		"extensions": BasicMessageResolver("Het :attribute veld moet een van de volgende extensies hebben: :args."),
		// "file":       BasicMessageResolver("Het :attribute veld moet een bestand zijn."),
		"filled": BasicMessageResolver("Het :attribute veld moet een waarde hebben."),
		// "gt": MessageHintResolver{Hints: map[string]string{
		// 	"array":   "Het :attribute veld moet meer dan :value items bevatten.",
		// 	"file":    "Het :attribute veld moet groter zijn dan :value kilobytes.",
		// 	"numeric": "Het :attribute veld moet groter zijn dan :value.",
		// 	"string":  "Het :attribute veld moet groter zijn dan :value tekens.",
		// }},
		// "gte": MessageHintResolver{Hints: map[string]string{
		// 	"array":   "Het :attribute veld moet :value items of meer bevatten.",
		// 	"file":    "Het :attribute veld moet groter zijn dan of gelijk aan :value kilobytes.",
		// 	"numeric": "Het :attribute veld moet groter zijn dan of gelijk aan :value.",
		// 	"string":  "Het :attribute veld moet groter zijn dan of gelijk aan :value tekens.",
		// }},
		"hex_color": BasicMessageResolver("Het :attribute veld moet een geldige hexadecimale kleur zijn."),
		// "image":     BasicMessageResolver("Het :attribute veld moet een afbeelding zijn."),
		"in": BasicMessageResolver("De geselecteerde :attribute is ongeldig."),
		// "in_array":  BasicMessageResolver("Het :attribute veld moet bestaan in :other."),
		// "integer":   BasicMessageResolver("Het :attribute veld moet een geheel getal zijn."),
		"ip":   BasicMessageResolver("Het :attribute veld moet een geldig IP-adres zijn."),
		"ipv4": BasicMessageResolver("Het :attribute veld moet een geldig IPv4-adres zijn."),
		"ipv6": BasicMessageResolver("Het :attribute veld moet een geldig IPv6-adres zijn."),
		"json": BasicMessageResolver("Het :attribute veld moet een geldige JSON-tekst zijn."),
		// "list":      BasicMessageResolver("Het :attribute veld moet een lijst zijn."),
		"lowercase": BasicMessageResolver("Het :attribute veld moet in kleine letters zijn."),
		// "lt": MessageHintResolver{Hints: map[string]string{
		// 	"array":   "Het :attribute veld moet minder dan :value items bevatten.",
		// 	"file":    "Het :attribute veld moet kleiner zijn dan :value kilobytes.",
		// 	"numeric": "Het :attribute veld moet kleiner zijn dan :value.",
		// 	"string":  "Het :attribute veld moet kleiner zijn dan :value tekens.",
		// }},
		// "lte": MessageHintResolver{Hints: map[string]string{
		// 	"array":   "Het :attribute veld mag niet meer dan :value items bevatten.",
		// 	"file":    "Het :attribute veld mag niet groter zijn dan of gelijk aan :value kilobytes.",
		// 	"numeric": "Het :attribute veld mag niet groter zijn dan of gelijk aan :value.",
		// 	"string":  "Het :attribute veld mag niet groter zijn dan of gelijk aan :value tekens.",
		// }},
		"mac_address": BasicMessageResolver("Het :attribute veld moet een geldig MAC-adres zijn."),
		"max": MessageHintResolver{
			Fallback: "Het :attribute veld mag niet groter zijn dan :arg.",
			Hints: map[string]string{
				"array":   "Het :attribute veld mag niet meer dan :arg items bevatten.",
				"file":    "Het :attribute veld mag niet groter zijn dan :arg kilobytes.",
				"numeric": "Het :attribute veld mag niet groter zijn dan :arg.",
				"string":  "Het :attribute veld mag niet groter zijn dan :arg tekens.",
			},
		},
		"max_digits": BasicMessageResolver("Het :attribute veld mag niet meer dan :max cijfers hebben."),
		"mimes":      BasicMessageResolver("Het :attribute veld moet een bestand zijn van het type: :args."),
		"mimetypes":  BasicMessageResolver("Het :attribute veld moet een bestand zijn van het type: :args."),
		"min": MessageHintResolver{
			Fallback: "Het :attribute veld moet minimaal :arg zijn.",
			Hints: map[string]string{
				"array":   "Het :attribute veld moet minimaal :arg items bevatten.",
				"file":    "Het :attribute veld moet minimaal :arg kilobytes zijn.",
				"numeric": "Het :attribute veld moet minimaal :arg zijn.",
				"string":  "Het :attribute veld moet minimaal :arg tekens lang zijn.",
			},
		},
		"min_digits": BasicMessageResolver("Het :attribute veld moet minimaal :arg cijfers hebben."),
		// "missing":          BasicMessageResolver("Het :attribute veld moet ontbreken."),
		// "missing_if":       BasicMessageResolver("Het :attribute veld moet ontbreken wanneer :other :value is."),
		// "missing_unless":   BasicMessageResolver("Het :attribute veld moet ontbreken tenzij :other in :args is."),
		// "missing_with":     BasicMessageResolver("Het :attribute veld moet ontbreken wanneer :args aanwezig is."),
		// "missing_with_all": BasicMessageResolver("Het :attribute veld moet ontbreken wanneer :args aanwezig zijn."),
		// "multiple_of":      BasicMessageResolver("Het :attribute veld moet een veelvoud van :value zijn."),
		"not_nil":   BasicMessageResolver("Het :attribute veld mag niet nil zijn."),
		"not_in":    BasicMessageResolver("De geselecteerde :attribute is ongeldig."),
		"not_regex": BasicMessageResolver("Het formaat van het :attribute veld is ongeldig."),
		"numeric":   BasicMessageResolver("Het :attribute veld moet een getal zijn."),
		// "password": MessageHintResolver{Hints: map[string]string{
		// 	"letters":       "Het :attribute veld moet minimaal één letter bevatten.",
		// 	"mixed":         "Het :attribute veld moet minimaal één hoofdletter en één kleine letter bevatten.",
		// 	"numbers":       "Het :attribute veld moet minimaal één cijfer bevatten.",
		// 	"symbols":       "Het :attribute veld moet minimaal één symbool bevatten.",
		// 	"uncompromised": "Het opgegeven :attribute is in een datalek verschenen. Kies een ander :attribute.",
		// }},
		// "present":           BasicMessageResolver("Het :attribute veld moet aanwezig zijn."),
		// "present_if":        BasicMessageResolver("Het :attribute veld moet aanwezig zijn wanneer :other :value is."),
		// "present_unless":    BasicMessageResolver("Het :attribute veld moet aanwezig zijn tenzij :other in :args is."),
		// "present_with":      BasicMessageResolver("Het :attribute veld moet aanwezig zijn wanneer :args aanwezig is."),
		// "present_with_all":  BasicMessageResolver("Het :attribute veld moet aanwezig zijn wanneer :args aanwezig zijn."),
		// "prohibited":        BasicMessageResolver("Het :attribute veld is verboden."),
		// "prohibited_if":     BasicMessageResolver("Het :attribute veld is verboden wanneer :other :value is."),
		// "prohibited_unless": BasicMessageResolver("Het :attribute veld is verboden tenzij :other in :args is."),
		// "prohibits":         BasicMessageResolver("Het :attribute veld verbiedt :other om aanwezig te zijn."),
		"regex":    BasicMessageResolver("Het formaat van het :attribute veld is ongeldig."),
		"required": BasicMessageResolver("Het :attribute veld is verplicht."),
		// "required_array_keys":  BasicMessageResolver("Het :attribute veld moet entries bevatten voor: :args."),
		// "required_if":          BasicMessageResolver("Het :attribute veld is verplicht wanneer :other :value is."),
		// "required_if_accepted": BasicMessageResolver("Het :attribute veld is verplicht wanneer :other is geaccepteerd."),
		// "required_if_declined": BasicMessageResolver("Het :attribute veld is verplicht wanneer :other is afgewezen."),
		// "required_unless":      BasicMessageResolver("Het :attribute veld is verplicht tenzij :other in :args is."),
		// "required_with":        BasicMessageResolver("Het :attribute veld is verplicht wanneer :args aanwezig is."),
		// "required_with_all":    BasicMessageResolver("Het :attribute veld is verplicht wanneer :args aanwezig zijn."),
		// "required_without":     BasicMessageResolver("Het :attribute veld is verplicht wanneer :args niet aanwezig is."),
		// "required_without_all": BasicMessageResolver("Het :attribute veld is verplicht wanneer geen van :args aanwezig zijn."),
		// "same": BasicMessageResolver("Het :attribute veld moet overeenkomen met :other."),
		"size": MessageHintResolver{
			Fallback: "Het :attribute veld moet de grootte :arg hebben.",
			Hints: map[string]string{
				"array":   "Het :attribute veld moet :arg items bevatten.",
				"file":    "Het :attribute veld moet :arg kilobytes zijn.",
				"numeric": "Het :attribute veld moet :arg zijn.",
				"string":  "Het :attribute veld moet :arg tekens lang zijn.",
			}},
		"starts_with": BasicMessageResolver("Het :attribute veld moet beginnen met een van de volgende: :args."),
		// "string":   BasicMessageResolver("Het :attribute veld moet een string zijn."),
		// "timezone": BasicMessageResolver("Het :attribute veld moet een geldige tijdzone zijn."),
		// "unique":   BasicMessageResolver("Het :attribute is al in gebruik genomen."),
		// "uploaded": BasicMessageResolver("Het uploaden van het :attribute is mislukt."),
		"uppercase": BasicMessageResolver("Het :attribute veld moet in hoofdletters zijn."),
		"url":       BasicMessageResolver("Het :attribute veld moet een geldige URL zijn."),
		"ulid":      BasicMessageResolver("Het :attribute veld moet een geldig ULID zijn."),
		"uuid":      BasicMessageResolver("Het :attribute veld moet een geldig UUID zijn."),
	})
}
