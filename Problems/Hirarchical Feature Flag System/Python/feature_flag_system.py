from channel import Channel

class FeatureFlagSystem:
    def __init__(self, repository):
        self.repository = repository

    def add_channel(self, channel_id, features, parents):
        channel = self.repository.get(channel_id)

        if channel is not None:
            raise ValueError("Channel already exists")

        channel = Channel(channel_id, features, parents)

        self.repository.add(channel_id, channel)

    def delete_channel(self, channel_id):
        channel = self.repository.get(channel_id)

        if channel is None:
            raise ValueError("Channel doesn't exists")

        self.repository.delete(channel_id)

    def set_feature(self, channel_id, feature_flag, value):
        channel = self.repository.get(channel_id)

        if channel is None:
            raise ValueError("Channel doesn't exists")

        channel.set_feature(feature_flag, value)

    def find_in_the_parent_chain(self, id, feature_flag):
        channel = self.repository.get(id)

        if channel is None:
            raise ValueError("Channel doesn't exists")

        if feature_flag in channel.get_features():
            return channel.get_features()[feature_flag]

        for parent_id  in channel.get_parents():
            value = self.find_in_the_parent_chain(parent_id, feature_flag)
            if value is not None:
                return value

        return None

    def get_feature(self, channel_id, feature_flag):
        channel = self.repository.get(channel_id)

        if channel is None:
            raise ValueError("Channel doesn't exists")

        if feature_flag in channel.get_features():
            return channel.get_features()[feature_flag]

        for parent_id  in channel.get_parents():
            value = self.find_in_the_parent_chain(parent_id, feature_flag)
            if value is not None:
                return value

        return False