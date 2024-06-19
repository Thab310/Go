package main

const (
	TshiVenda = "TshiVenda"
	Xitsonga  = "Xitsonga"

	HelloEnglishPrefix = "Hello, "
	HelloVendaPrefix   = "Hendaa, "
	HelloTsongaPrefix  = "Avuxeni, "
)

func Hello(n string, l string) string {
	if n == "" {
		n = "gents"
	}
	return greetingPrefix(l) + n
}

func greetingPrefix(l string) (prefix string) {
	switch l {
	case TshiVenda:
		prefix = HelloVendaPrefix
	case Xitsonga:
		prefix = HelloTsongaPrefix
	default:
		prefix = HelloEnglishPrefix
	}
	return
}
