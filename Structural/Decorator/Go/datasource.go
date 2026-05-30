package main

type Datasource interface {
	WriteData(data string)
	ReadData() string
}

type FileDatasource struct{}

func NewFileDatasource() *FileDatasource {
	return &FileDatasource{}
}

func (fd *FileDatasource) WriteData(data string) {
	println("(FileDatasource) Writing data to file...")
}

func (fd *FileDatasource) ReadData() string {
	println("(FileDatasource) Reading data from file...")
	return "data"
}
