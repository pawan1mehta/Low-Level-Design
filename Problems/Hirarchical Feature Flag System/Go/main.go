package main

import "fmt"

func main() {
	channelRepository := NewInMemoryChannelRepository()

	featureFlagSystem := NewFeatureFlagSystem(channelRepository)

	prodFeatures := map[string]bool{
		"new-ui": true,
	}
	featureFlagSystem.AddChannel("prod", prodFeatures, []string{})

	stagingFeatures := map[string]bool{
		"new-ui": false,
	}
	featureFlagSystem.AddChannel("staging", stagingFeatures, []string{})

	indianEnvdFeatures := map[string]bool{}
	featureFlagSystem.AddChannel("indian-env", indianEnvdFeatures, []string{"prod"})

	value, _ := featureFlagSystem.GetFeature("prod", "new-ui")
	fmt.Println("prod: new-ui: ", *value)
	value, _ = featureFlagSystem.GetFeature("staging", "new-ui")
	fmt.Println("staging: new-ui: ", *value)
	value, _ = featureFlagSystem.GetFeature("indian-env", "new-ui")
	fmt.Println("prod: new-ui: ", *value)
}
