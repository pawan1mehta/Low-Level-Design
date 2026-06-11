package adapter

type JSONStockDataProvider struct {
	xmlStockDataProvider *XMLStockDataProvider
}

func NewJSONStockDataProvider(xmlStockDataProvider *XMLStockDataProvider) *JSONStockDataProvider {
	return &JSONStockDataProvider{
		xmlStockDataProvider: xmlStockDataProvider,
	}
}

func (jsdp *JSONStockDataProvider) GetData() string {
	return jsdp.ConvertXmlTOJson(jsdp.xmlStockDataProvider.GetData())
}

func (jsdp *JSONStockDataProvider) ConvertXmlTOJson(xmlData string) string {
	return "Stocks data in JSON"
}
