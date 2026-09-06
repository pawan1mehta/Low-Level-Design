package main

type ConfigurationStore struct {
	defaultConfig   EndpointConfiguration
	endpointConfigs map[string]EndpointConfiguration
}

func NewConfigurationStore(defaultConfig EndpointConfiguration) *ConfigurationStore {
	return &ConfigurationStore{
		defaultConfig:   defaultConfig,
		endpointConfigs: map[string]EndpointConfiguration{},
	}
}

func (cfgStore *ConfigurationStore) AddEndpoint(endpoint string, endpointConfig EndpointConfiguration) error {
	if _, exists := cfgStore.endpointConfigs[endpoint]; exists {
		return ErrEndpointConfigAlreadyExists
	}
	cfgStore.endpointConfigs[endpoint] = endpointConfig
	return nil
}

func (cfgStore *ConfigurationStore) GetConfig(endpoint string) (EndpointConfiguration, error) {
	cfg, ok := cfgStore.endpointConfigs[endpoint]
	if !ok {
		return cfgStore.defaultConfig, nil
	}
	return cfg, nil
}
