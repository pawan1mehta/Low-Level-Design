package main

import "errors"

type FeatureFlagSystem struct {
	repository Repository
}

func NewFeatureFlagSystem(repository Repository) *FeatureFlagSystem {
	return &FeatureFlagSystem{
		repository: repository,
	}
}

func (fsys *FeatureFlagSystem) AddChannel(channelID string, features map[string]bool, parents []string) error {
	channel := fsys.repository.Get(channelID)
	if channel != nil {
		return errors.New("can not add, channel already exits")
	}
	channel = NewChannel(channelID, features, parents)
	fsys.repository.Add(channelID, channel)
	return nil
}

func (fsys *FeatureFlagSystem) DeleteChannel(channelID string) error {
	channel := fsys.repository.Get(channelID)
	if channel == nil {
		return errors.New("channel doesn't exits")
	}
	fsys.repository.Delete(channelID)
	return nil
}

func (fsys *FeatureFlagSystem) SetFeature(channelID string, featureFlag string, value bool) error {
	channel := fsys.repository.Get(channelID)
	if channel == nil {
		return errors.New("channel doesn't exits")
	}
	channel.SetFeature(featureFlag, value)
	return nil
}

func (fsys *FeatureFlagSystem) getFeatureFromParentIfNotDefined(id string, featureFlag string) (*bool, error) {
	channel := fsys.repository.Get(id)
	if channel == nil {
		return nil, errors.New("channel doesn't exits")
	}

	value, ok := channel.GetFeatures()[featureFlag]
	if ok {
		return &value, nil
	}

	for _, id := range channel.GetParents() {
		inheritedValue, err := fsys.getFeatureFromParentIfNotDefined(id, featureFlag)
		if err != nil {
			continue
		}
		if inheritedValue != nil {
			return inheritedValue, nil
		}
	}

	return nil, nil
}

func (fsys *FeatureFlagSystem) GetFeature(channelID string, featureFlag string) (*bool, error) {
	channel := fsys.repository.Get(channelID)
	if channel == nil {
		return nil, errors.New("channel doesn't exits")
	}

	value, ok := channel.GetFeatures()[featureFlag]
	if ok {
		return &value, nil
	}

	var inheritedValue bool

	for _, id := range channel.GetParents() {
		inheritedValue, err := fsys.getFeatureFromParentIfNotDefined(id, featureFlag)
		if err != nil {
			continue
		}
		if inheritedValue != nil {
			return inheritedValue, nil
		}
	}

	inheritedValue = false

	return &inheritedValue, nil
}
