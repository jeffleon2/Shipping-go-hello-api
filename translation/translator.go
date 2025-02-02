package translation

import "strings"

// StaticService has data that does not change.
type StaticService struct{}

// NewStaticService creates new instance of a service that uses static data.
func NewStaticService() *StaticService {
	return &StaticService{}
}

func (s *StaticService) Translate(word string, language string) string {
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
	case "french":
		return "bonjour"
	default:
		return ""
	}
}

func sanitizeinput(w string) string {
	w = strings.ToLower(w)
	return strings.TrimSpace(w)
}
