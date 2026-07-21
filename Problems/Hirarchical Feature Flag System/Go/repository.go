package main

type Repository interface {
	List() []*Channel
	Get(id string) *Channel
	Add(id string, channel *Channel)
	Delete(id string)
}

type ChannelRepository interface {
	Repository
}

type InMemoryChannelRepository struct {
	channels map[string]*Channel
}

func NewInMemoryChannelRepository() *InMemoryChannelRepository {
	return &InMemoryChannelRepository{
		channels: make(map[string]*Channel),
	}
}

func (r *InMemoryChannelRepository) List() []*Channel {
	var channels []*Channel

	for _, channel := range r.channels {
		channels = append(channels, channel)
	}

	return channels
}

func (r *InMemoryChannelRepository) Get(id string) *Channel {
	channel, ok := r.channels[id]
	if !ok {
		return nil
	}
	return channel
}

func (r *InMemoryChannelRepository) Add(id string, channel *Channel) {
	r.channels[id] = channel
}

func (r *InMemoryChannelRepository) Delete(id string) {
	_, ok := r.channels[id]
	if !ok {
		return
	}
	delete(r.channels, id)
}
