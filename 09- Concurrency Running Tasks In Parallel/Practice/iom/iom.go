package iom

type IOManager interface {
	ReadLinesFromFile()(lines []string, err error)
	WriteJson(data any) error
}


