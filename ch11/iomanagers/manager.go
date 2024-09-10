package iomanagers

type IOManager interface {
	WriteResult(interface{}) error
	ReadLines() ([]string, error)
}
