package main

type InventoryItem struct {
	productID string
	count     int64
}

func (i *InventoryItem) Add(count int64) {
	i.count += count
}

func (i *InventoryItem) Remove(count int64) {
	i.count -= count
}

func (i *InventoryItem) ProductID() string {
	return i.productID
}

func (i *InventoryItem) Count() int64 {
	return i.count
}

type Stock struct {
	productID string
	quantity  int64
}

func (s Stock) GetQuantity() int64 {
	return s.quantity
}

func (s Stock) GetProductID() string {
	return s.productID
}
