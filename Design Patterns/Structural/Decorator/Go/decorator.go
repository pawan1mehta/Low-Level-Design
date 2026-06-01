package main

type DataSourceDecorator struct {
	datasource Datasource
}

func NewDataSourceDecorator(datasource Datasource) *DataSourceDecorator {
	return &DataSourceDecorator{
		datasource: datasource,
	}
}

func (d *DataSourceDecorator) WriteData(data string) {
	d.datasource.WriteData(data)
}

func (d *DataSourceDecorator) ReadData() string {
	return d.datasource.ReadData()
}

type EncryptionDataSource struct {
	*DataSourceDecorator
}

func NewEncryptionDataSource(datasource Datasource) *EncryptionDataSource {
	return &EncryptionDataSource{
		NewDataSourceDecorator(datasource),
	}
}

func (ed *EncryptionDataSource) WriteData(data string) {
	println("(EncryptionDataSource) Writing data to file...")
	encryptedData := ed.EncryptData(data)
	ed.DataSourceDecorator.WriteData(encryptedData)
}

func (ed *EncryptionDataSource) ReadData() string {
	data := ed.DataSourceDecorator.ReadData()
	println("(EncryptionDataSource) Reading data from file...")
	return ed.DecryptData(data)
}

func (ed *EncryptionDataSource) EncryptData(data string) string {
	return "encryptedData"
}

func (ed *EncryptionDataSource) DecryptData(data string) string {
	return "decryptedData"
}
