package main

type AlertType int

const (
	ProductAlert AlertType = iota
	WarehouseAlert
)

type AlertConfig struct {
	alertType      AlertType
	productId      string
	thresholdValue int64
	enabled        bool
}

func (aconfig AlertConfig) Type() AlertType {
	return aconfig.alertType
}

func (aconfig AlertConfig) ProductID() string {
	return aconfig.productId
}

func (aconfig AlertConfig) ThresholdValue() int64 {
	return aconfig.thresholdValue
}

func (aconfig AlertConfig) Enabled() bool {
	return aconfig.enabled
}
