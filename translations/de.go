package translations

import (
	. "github.com/mjarkk/laravalidate"
	"golang.org/x/text/language"
)

func RegisterDeTranslations() {
	RegisterMessages(language.German, map[string]MessageResolver{
		"accepted": BasicMessageResolver("Das :attribute Feld muss akzeptiert werden."),
		// "accepted_if": BasicMessageResolver("Das :attribute Feld muss akzeptiert werden, wenn :other :value ist."),
		"active_url":     BasicMessageResolver("Das :attribute Feld muss eine gültige URL sein."),
		"after":          BasicMessageResolver("Das :attribute Feld muss ein Datum nach :date sein."),
		"after_or_equal": BasicMessageResolver("Das :attribute Feld muss ein Datum nach oder gleich :date sein."),
		"alpha":          BasicMessageResolver("Das :attribute Feld darf nur Buchstaben enthalten."),
		"alpha_dash":     BasicMessageResolver("Das :attribute Feld darf nur Buchstaben, Zahlen, Bindestriche und Unterstriche enthalten."),
		"alpha_numeric":  BasicMessageResolver("Das :attribute Feld darf nur Buchstaben und Zahlen enthalten."),
		// "array":     BasicMessageResolver("Das :attribute Feld muss ein Array sein."),
		"ascii":           BasicMessageResolver("Das :attribute Feld darf nur einstellige alphanumerische Zeichen und Symbole enthalten."),
		"bail":            BasicMessageResolver("Das :attribute Feld muss gültig sein."),
		"before":          BasicMessageResolver("Das :attribute Feld muss ein Datum vor :date sein."),
		"before_or_equal": BasicMessageResolver("Das :attribute Feld muss ein Datum vor oder gleich :date sein."),
		"between": MessageHintResolver{
			Fallback: "Das :attribute Feld muss zwischen :arg0 und :arg1 liegen.",
			Hints: map[string]string{
				"array":   "Das :attribute Feld muss zwischen :arg0 und :arg1 Elementen haben.",
				"file":    "Das :attribute Feld muss zwischen :arg0 und :arg1 Kilobytes groß sein.",
				"numeric": "Das :attribute Feld muss zwischen :arg0 und :arg1 liegen.",
				"string":  "Das :attribute Feld muss zwischen :arg0 und :arg1 Zeichen lang sein.",
			},
		},
		"boolean": BasicMessageResolver("Das :attribute Feld muss wahr oder falsch sein."),
		// "can": BasicMessageResolver("Das :attribute Feld enthält einen nicht autorisierten Wert."),
		"confirmed": BasicMessageResolver("Die Bestätigung des :attribute Feldes stimmt nicht überein."),
		// "contains":  BasicMessageResolver("Das :attribute Feld fehlt ein erforderlicher Wert."),
		// "current_password": BasicMessageResolver("Das Passwort ist falsch."),
		"date": BasicMessageResolver("Das :attribute Feld muss ein gültiges Datum sein."),
		// "date_equals": BasicMessageResolver("Das :attribute Feld muss ein Datum gleich :date sein."),
		"date_format": BasicMessageResolver("Das :attribute Feld muss dem Format :arg entsprechen."),
		// "decimal": BasicMessageResolver("Das :attribute Feld muss :arg Dezimalstellen haben."),
		"declined": BasicMessageResolver("Das :attribute Feld muss abgelehnt werden."),
		// "declined_if": BasicMessageResolver("Das :attribute Feld muss abgelehnt werden, wenn :other :value ist."),
		// "different":   BasicMessageResolver("Das :attribute Feld und :other müssen unterschiedlich sein."),
		"digits":         BasicMessageResolver("Das :attribute Feld muss :digits Ziffern haben."),
		"digits_between": BasicMessageResolver("Das :attribute Feld muss zwischen :arg0 und :arg1 Ziffern haben."),
		// "dimensions":        BasicMessageResolver("Das :attribute Feld hat ungültige Bildabmessungen."),
		// "distinct":          BasicMessageResolver("Das :attribute Feld hat einen doppelten Wert."),
		// "doesnt_end_with":   BasicMessageResolver("Das :attribute Feld darf nicht mit einem der folgenden Werte enden: :args."),
		// "doesnt_start_with": BasicMessageResolver("Das :attribute Feld darf nicht mit einem der folgenden Werte beginnen: :args."),
		"email":     BasicMessageResolver("Das :attribute Feld muss eine gültige E-Mail-Adresse sein."),
		"ends_with": BasicMessageResolver("Das :attribute Feld muss mit einem der folgenden Werte enden: :args."),
		// "enum":    BasicMessageResolver("Der ausgewählte :attribute ist ungültig."),
		// "exists":  BasicMessageResolver("Der ausgewählte :attribute ist ungültig."),
		"extensions": BasicMessageResolver("Das :attribute Feld muss eine der folgenden Erweiterungen haben: :args."),
		// "file":       BasicMessageResolver("Das :attribute Feld muss eine Datei sein."),
		"filled": BasicMessageResolver("Das :attribute Feld muss einen Wert haben."),
		// "gt": MessageHintResolver{Hints: map[string]string{
		// 	"array":   "Das :attribute Feld muss mehr als :value Elemente haben.",
		// 	"file":    "Das :attribute Feld muss größer als :value Kilobytes sein.",
		// 	"numeric": "Das :attribute Feld muss größer als :value sein.",
		// 	"string":  "Das :attribute Feld muss größer als :value Zeichen lang sein.",
		// }},
		// "gte": MessageHintResolver{Hints: map[string]string{
		// 	"array":   "Das :attribute Feld muss :value Elemente oder mehr haben.",
		// 	"file":    "Das :attribute Feld muss größer oder gleich :value Kilobytes sein.",
		// 	"numeric": "Das :attribute Feld muss größer oder gleich :value sein.",
		// 	"string":  "Das :attribute Feld muss größer oder gleich :value Zeichen lang sein.",
		// }},
		"hex_color": BasicMessageResolver("Das :attribute Feld muss eine gültige hexadezimale Farbe sein."),
		// "image":     BasicMessageResolver("Das :attribute Feld muss ein Bild sein."),
		"in": BasicMessageResolver("Der ausgewählte :attribute ist ungültig."),
		// "in_array":  BasicMessageResolver("Das :attribute Feld muss in :other existieren."),
		// "integer":   BasicMessageResolver("Das :attribute Feld muss eine Ganzzahl sein."),
		"ip":   BasicMessageResolver("Das :attribute Feld muss eine gültige IP-Adresse sein."),
		"ipv4": BasicMessageResolver("Das :attribute Feld muss eine gültige IPv4-Adresse sein."),
		"ipv6": BasicMessageResolver("Das :attribute Feld muss eine gültige IPv6-Adresse sein."),
		"json": BasicMessageResolver("Das :attribute Feld muss eine gültige JSON-Zeichenkette sein."),
		// "list":      BasicMessageResolver("Das :attribute Feld muss eine Liste sein."),
		"lowercase": BasicMessageResolver("Das :attribute Feld muss in Kleinbuchstaben sein."),
		// "lt": MessageHintResolver{Hints: map[string]string{
		// 	"array":   "Das :attribute Feld muss weniger als :value Elemente haben.",
		// 	"file":    "Das :attribute Feld muss kleiner als :value Kilobytes sein.",
		// 	"numeric": "Das :attribute Feld muss kleiner als :value sein.",
		// 	"string":  "Das :attribute Feld muss kleiner als :value Zeichen lang sein.",
		// }},
		// "lte": MessageHintResolver{Hints: map[string]string{
		// 	"array":   "Das :attribute Feld darf nicht mehr als :value Elemente haben.",
		// 	"file":    "Das :attribute Feld darf nicht größer als :value Kilobytes sein.",
		// 	"numeric": "Das :attribute Feld darf nicht größer als :value sein.",
		// 	"string":  "Das :attribute Feld darf nicht größer als :value Zeichen lang sein.",
		// }},
		"mac_address": BasicMessageResolver("Das :attribute Feld muss eine gültige MAC-Adresse sein."),
		"max": MessageHintResolver{
			Fallback: "Das :attribute Feld darf nicht größer als :arg sein.",
			Hints: map[string]string{
				"array":   "Das :attribute Feld darf nicht mehr als :arg Elemente haben.",
				"file":    "Das :attribute Feld darf nicht größer als :arg Kilobytes sein.",
				"numeric": "Das :attribute Feld darf nicht größer als :arg sein.",
				"string":  "Das :attribute Feld darf nicht größer als :arg Zeichen lang sein.",
			},
		},
		"max_digits": BasicMessageResolver("Das :attribute Feld darf nicht mehr als :max Ziffern haben."),
		"mimes":      BasicMessageResolver("Das :attribute Feld muss eine Datei vom Typ :args sein."),
		"mimetypes":  BasicMessageResolver("Das :attribute Feld muss eine Datei vom Typ :args sein."),
		"min": MessageHintResolver{
			Fallback: "Das :attribute Feld muss mindestens :arg sein.",
			Hints: map[string]string{
				"array":   "Das :attribute Feld muss mindestens :arg Elemente haben.",
				"file":    "Das :attribute Feld muss mindestens :arg Kilobytes groß sein.",
				"numeric": "Das :attribute Feld muss mindestens :arg sein.",
				"string":  "Das :attribute Feld muss mindestens :arg Zeichen lang sein.",
			},
		},
		"min_digits": BasicMessageResolver("Das :attribute Feld muss mindestens :arg Ziffern haben."),
		// "missing":          BasicMessageResolver("Das :attribute Feld muss fehlen."),
		// "missing_if":       BasicMessageResolver("Das :attribute Feld muss fehlen, wenn :other :value ist."),
		// "missing_unless":   BasicMessageResolver("Das :attribute Feld muss fehlen, es sei denn :other ist in :args."),
		// "missing_with":     BasicMessageResolver("Das :attribute Feld muss fehlen, wenn :args vorhanden ist."),
		// "missing_with_all": BasicMessageResolver("Das :attribute Feld muss fehlen, wenn :args vorhanden sind."),
		// "multiple_of":      BasicMessageResolver("Das :attribute Feld muss ein Vielfaches von :value sein."),
		"not_nil":   BasicMessageResolver("Das :attribute Feld darf nicht nil sein."),
		"not_in":    BasicMessageResolver("Der ausgewählte :attribute ist ungültig."),
		"not_regex": BasicMessageResolver("Das Format des :attribute Feldes ist ungültig."),
		"numeric":   BasicMessageResolver("Das :attribute Feld muss eine Zahl sein."),
		// "password": MessageHintResolver{Hints: map[string]string{
		// 	"letters":       "Das :attribute Feld muss mindestens einen Buchstaben enthalten.",
		// 	"mixed":         "Das :attribute Feld muss mindestens einen Groß- und einen Kleinbuchstaben enthalten.",
		// 	"numbers":       "Das :attribute Feld muss mindestens eine Zahl enthalten.",
		// 	"symbols":       "Das :attribute Feld muss mindestens ein Symbol enthalten.",
		// 	"uncompromised": "Das angegebene :attribute ist in einem Datenleck aufgetaucht. Bitte wählen Sie ein anderes :attribute.",
		// }},
		// "present":           BasicMessageResolver("Das :attribute Feld muss vorhanden sein."),
		// "present_if":        BasicMessageResolver("Das :attribute Feld muss vorhanden sein, wenn :other :value ist."),
		// "present_unless":    BasicMessageResolver("Das :attribute Feld muss vorhanden sein, es sei denn :other ist in :args."),
		// "present_with":      BasicMessageResolver("Das :attribute Feld muss vorhanden sein, wenn :args vorhanden ist."),
		// "present_with_all":  BasicMessageResolver("Das :attribute Feld muss vorhanden sein, wenn :args vorhanden sind."),
		// "prohibited":        BasicMessageResolver("Das :attribute Feld ist verboten."),
		// "prohibited_if":     BasicMessageResolver("Das :attribute Feld ist verboten, wenn :other :value ist."),
		// "prohibited_unless": BasicMessageResolver("Das :attribute Feld ist verboten, es sei denn :other ist in :args."),
		// "prohibits":         BasicMessageResolver("Das :attribute Feld verbietet :other von der Anwesenheit."),
		"regex":    BasicMessageResolver("Das Format des :attribute Feldes ist ungültig."),
		"required": BasicMessageResolver("Das :attribute Feld ist erforderlich."),
		// "required_array_keys":  BasicMessageResolver("Das :attribute Feld muss Einträge für :args enthalten."),
		// "required_if":          BasicMessageResolver("Das :attribute Feld ist erforderlich, wenn :other :value ist."),
		// "required_if_accepted": BasicMessageResolver("Das :attribute Feld ist erforderlich, wenn :other akzeptiert ist."),
		// "required_if_declined": BasicMessageResolver("Das :attribute Feld ist erforderlich, wenn :other abgelehnt ist."),
		// "required_unless":      BasicMessageResolver("Das :attribute Feld ist erforderlich, es sei denn :other ist in :args."),
		// "required_with":        BasicMessageResolver("Das :attribute Feld ist erforderlich, wenn :args vorhanden ist."),
		// "required_with_all":    BasicMessageResolver("Das :attribute Feld ist erforderlich, wenn :args vorhanden sind."),
		// "required_without":     BasicMessageResolver("Das :attribute Feld ist erforderlich, wenn :args nicht vorhanden ist."),
		// "required_without_all": BasicMessageResolver("Das :attribute Feld ist erforderlich, wenn keine von :args vorhanden sind."),
		// "same": BasicMessageResolver("Das :attribute Feld muss mit :other übereinstimmen."),
		"size": MessageHintResolver{
			Fallback: "Das :attribute Feld muss die Größe :arg haben.",
			Hints: map[string]string{
				"array":   "Das :attribute Feld muss :arg Elemente enthalten.",
				"file":    "Das :attribute Feld muss :arg Kilobytes groß sein.",
				"numeric": "Das :attribute Feld muss :arg sein.",
				"string":  "Das :attribute Feld muss :arg Zeichen lang sein.",
			}},
		"starts_with": BasicMessageResolver("Das :attribute Feld muss mit einem der folgenden Werte beginnen: :args."),
		// "string":   BasicMessageResolver("Das :attribute Feld muss eine Zeichenkette sein."),
		// "timezone": BasicMessageResolver("Das :attribute Feld muss eine gültige Zeitzone sein."),
		// "unique":   BasicMessageResolver("Das :attribute ist bereits vergeben."),
		// "uploaded": BasicMessageResolver("Das :attribute Feld konnte nicht hochgeladen werden."),
		"uppercase": BasicMessageResolver("Das :attribute Feld muss in Großbuchstaben sein."),
		"url":       BasicMessageResolver("Das :attribute Feld muss eine gültige URL sein."),
		"ulid":      BasicMessageResolver("Das :attribute Feld muss eine gültige ULID sein."),
		"uuid":      BasicMessageResolver("Das :attribute Feld muss eine gültige UUID sein."),
	})
}