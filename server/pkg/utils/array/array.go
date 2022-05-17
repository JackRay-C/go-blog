package array


type Array []string

func (a Array) Contains(elem string) bool {
	for _, e := range a {
		if elem == e {
			return true
		}
	}
	return false
}
