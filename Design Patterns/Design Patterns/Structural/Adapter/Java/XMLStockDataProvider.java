package Structural.Adapter.Java;

public class XMLStockDataProvider implements StockDataProvider {

    public String getData() {
        return "Stocks data in XML";
    }
}