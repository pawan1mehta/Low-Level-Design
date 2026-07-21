class Channel:
    def __init__(self, id, features, parents):
        self.id = id
        self.features = features
        self.parents = parents

    def set_feature(self, feature_flag, value):
        self.features[feature_flag] = value

    def get_features(self):
        return self.features

    def get_parents(self):
        return self.parents