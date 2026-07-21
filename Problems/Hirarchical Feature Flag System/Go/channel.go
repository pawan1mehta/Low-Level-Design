package main

type Channel struct {
	id       string
	features map[string]bool
	parents  []string
}

func NewChannel(id string, features map[string]bool, parents []string) *Channel {
	return &Channel{
		id:       id,
		features: features,
		parents:  parents,
	}
}

func (c *Channel) SetFeature(featureFlag string, value bool) {
	c.features[featureFlag] = value
}

func (c *Channel) GetFeatures() map[string]bool {
	return c.features
}

func (c *Channel) GetParents() []string {
	return c.parents
}
