# Inventory Management

## Clarifying Questions

- Will there be fixed set of warehouses or admin can add warehouses dynamically? [Ans: fixed set of warehouses]
- What is the threshold for 'low-stock' alerts? [Ans: per product & per warehouse]
- Do we need to handle the concurrency? [Ans: Yes]

## Requirements

- Admin can add/remove a stock from a specific warehouses
- Admin can check the availability?
  - Given a product & quantity, return which warehouses can fulfill it?
- Admin can transfer the stock from one warehouse to other warehouses
- Admin should get alert for the 'low-stock'

Out Of Scope

- Notification sevice
- Transportation service
- Order processing

## Core Entities & Relationships

- InventoryManager
- Stock
- Product
- Warehouses
- AlertConfig
- AlertListener
- IntentoryItem

Relationships:

```text
InventoryManager <----- has a ---- Warehouse

Warehouse <---- composed of ---- InventoryItem
```

## Class Design

```code
Class InventoryManager

    - warehouses: Map<string, WareHouse>

    + addStock(warehouseID: string, stocks: []Stock, warehouseID) -> boolean
    + removeStock(warehouseID: string, stocks: []Stock, warehouseID) -> boolean
    + checkAvailability(stocks: []Stock) -> string
    + transferStock(sourceWarehouseID: string, targetWarehouseID: string, stocks: []Stock) -> boolean
    + addAlert(alertConfig: AlertCofig)
```

```code
Class Warehouse:

    - id: string
    - inventory: Map<String, IntentoryItem>
    - maxCapcity: int
    - alertConfigs: []AlertConfig

    + getProductQuantity(productID: string) -> int
    + getTotalQuantity() -> int
    + addStocks(stock: []Stock) -> boolean
    + removeStocks(stock: []Stock) -> boolean
```

```code
Class IntentoryItem:

    - id: productId
    - quantity: int
```

```code
Class Stock:

    - productID: string
    - quantity: int
```

```code
Class AlertConfig:

    - type: PRODUCT_ALERT | WAREHOUSE_ALERT
    - warehouseID: string
    - productID: string
    - thresholdValue: int
    - enable
```
