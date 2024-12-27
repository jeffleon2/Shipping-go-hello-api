package translation

import "strings"

func Translate(word string, language string) string {
	word = sanitizeinput(word)
	language = sanitizeinput(language)

	if word != "hello" {
		return ""
	}

	switch language {
	case "english":
		return "hello"
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	default:
		return ""
	}
}

func sanitizeinput(w string) string {
	w = strings.ToLower(w)
	return strings.TrimSpace(w)
}
