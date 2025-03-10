package dictionary

type dictionary map[string]string

func (d dictionary) Search(di map[string]string, word string) string {
	return d[word]
}
