package main

import (
	"fmt"
	"sync"
)

type Warehouse struct {
	id           string
	mu           sync.RWMutex
	inventory    map[string]*InventoryItem
	maxCapacity  int64
	alertConfigs []AlertConfig
}

func (w *Warehouse) AddStocks(stocks []Stock) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	var totalToAdd int64

	for _, stock := range stocks {
		id := stock.GetProductID()
		quantity := stock.GetQuantity()

		if quantity <= 0 {
			return fmt.Errorf("stock quantity must be positive")
		}

		if _, ok := w.inventory[id]; !ok {
			return fmt.Errorf("product id: %q not found", id)
		}

		totalToAdd += quantity
	}

	if w.getTotalCount()+totalToAdd > w.maxCapacity {
		return ErrWarehouseFull
	}

	for _, stock := range stocks {
		id := stock.GetProductID()
		quantity := stock.GetQuantity()
		st := w.inventory[id]
		st.Add(quantity)
	}

	w.evaluateAlerts()

	return nil
}

func (w *Warehouse) RemoveStocks(stocks []Stock) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	for _, stock := range stocks {
		id := stock.GetProductID()
		quantity := stock.GetQuantity()

		if quantity <= 0 {
			return fmt.Errorf("stock quantity must be positive")
		}

		st, ok := w.inventory[id]
		if !ok {
			return fmt.Errorf("product id: %q not found", id)
		}

		if quantity > st.Count() {
			return fmt.Errorf("%w for product %q", ErrInsufficientStock, id)
		}
	}

	for _, stock := range stocks {
		id := stock.GetProductID()
		st := w.inventory[id]
		st.Remove(stock.GetQuantity())
	}

	w.evaluateAlerts()

	return nil
}

func (w *Warehouse) getTotalCount() int64 {
	w.mu.RLock()
	defer w.mu.RUnlock()

	var totalCount int64
	for _, stock := range w.inventory {
		totalCount += stock.Count()
	}
	return totalCount
}

func (w *Warehouse) evaluateAlerts() {
	for _, alert := range w.alertConfigs {
		if alert.Enabled() {
			switch alert.Type() {
			case ProductAlert:
				productID := alert.ProductID()
				st, ok := w.inventory[productID]
				if !ok {
					continue
				}
				if st.Count() <= alert.ThresholdValue() {
					// product alert
				}

			case WarehouseAlert:
				if w.getTotalCount() >= alert.ThresholdValue() {
					// warehouse alert
				}
			default:

			}
		}
	}
}

func (w *Warehouse) checkAvailability(stock Stock) bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	
	item, ok := w.inventory[stock.GetProductID()]
	if !ok {
		return false
	}
	return item.Count() >= stock.GetQuantity()
}

func (w *Warehouse) AddAlert(config AlertConfig) error {
	w.alertConfigs = append(w.alertConfigs, config)
	return nil
}
