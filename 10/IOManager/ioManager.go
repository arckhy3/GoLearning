package iomanager

type IOManager interface {
	LoadData() ([]string, error)
	WriteData(data any) error
}
