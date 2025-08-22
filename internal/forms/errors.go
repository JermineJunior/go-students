package forms

type Errors map[string][]string

func (e Errors) Add(field, message string) {
	e[field] = append(e[field], message)
}
func (e Errors) Get(field string) []string {
	errors := e[field]

	if len(errors) == 0 {
		return nil
	}
	return errors
}
