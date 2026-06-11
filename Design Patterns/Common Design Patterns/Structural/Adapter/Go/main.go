package main

import "adapter/adapter"

func main() {

	sdp := &adapter.XMLStockDataProvider{}

	jsdkp := adapter.NewJSONStockDataProvider(sdp)

	ae := &adapter.AnalyticsEngine{}
	ae.AnalyzeJSONData(jsdkp.GetData())
}
