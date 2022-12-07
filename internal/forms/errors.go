package forms

type errors map[string][]string

// Add adds an error message for a given form field
func (e errors) Add(field string, message string) {
	e[field] = append(e[field], message)
}

// Get returns or retrieves the first error message
func (e errors) Get(field string) string {
	es, ok := e[field]
	if len(es) == 0 || !ok {
		return ""
	}

	return es[0]
}
