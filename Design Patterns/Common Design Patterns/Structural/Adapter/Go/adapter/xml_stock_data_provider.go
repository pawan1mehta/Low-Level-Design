package adapter

type XMLStockDataProvider struct {
}

func (x *XMLStockDataProvider) GetData() string {
	return "Stocks data in XML"
}
