class InMemoryChannelRepository:
    def __init__(self):
        self.channels = {}

    def add(self, channel_id, channel):
        self.channels[channel_id] = channel

    def get(self, channel_id):
        return self.channels.get(channel_id)

    def delete(self, channel_id):
        if channel_id in self.channels:
            del self.channels[channel_id]