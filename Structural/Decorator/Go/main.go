package main

func main() {
	filedataSource := NewFileDatasource()

	datasourceDecorator := NewEncryptionDataSource(filedataSource)

	datasourceDecorator.WriteData("data")
	_ = datasourceDecorator.ReadData()
}
