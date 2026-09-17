package main

import "fmt"

type InventoryManagement struct {
	warehouses map[string]*Warehouse
}

func (im *InventoryManagement) AddStocks(warehouseID string, stock []Stock) error {
	warehouse, ok := im.warehouses[warehouseID]
	if !ok {
		return ErrWarehouseNotFound
	}
	return warehouse.AddStocks(stock)
}

func (im *InventoryManagement) RemoveStocks(warehouseID string, stock []Stock) error {
	warehouse, ok := im.warehouses[warehouseID]
	if !ok {
		return ErrWarehouseNotFound
	}
	return warehouse.RemoveStocks(stock)
}

func (im *InventoryManagement) checkAvailability(stock Stock) (string, error) {
	for id, warehouse := range im.warehouses {
		if warehouse.checkAvailability(stock) {
			return id, nil
		}
	}
	return "", ErrNoAvailableWarehouse
}

func (im *InventoryManagement) TransferStock(sourceWarehouseID, targetWarehouseID string, stocks []Stock) error {
	if sourceWarehouseID == targetWarehouseID {
		return fmt.Errorf("source and target warehouses must be different")
	}
	
	sourceWarehouse, ok := im.warehouses[sourceWarehouseID]
	if !ok {
		return ErrWarehouseNotFound
	}
	targetWarehouse, ok := im.warehouses[targetWarehouseID]
	if !ok {
		return ErrWarehouseNotFound
	}

	if err := sourceWarehouse.RemoveStocks(stocks); err != nil {
		return err
	}
	return targetWarehouse.AddStocks(stocks)
}

func (im *InventoryManagement) AddAlert(warehouseID string, alertConfig AlertConfig) error {
	warehouse, ok := im.warehouses[warehouseID]
	if !ok {
		return ErrWarehouseNotFound
	}
	return warehouse.AddAlert(alertConfig)
}
