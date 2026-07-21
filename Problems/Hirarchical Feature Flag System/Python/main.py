from feature_flag_system import FeatureFlagSystem
from repository import InMemoryChannelRepository

def main():
    repository = InMemoryChannelRepository()

    featureFlagSystem = FeatureFlagSystem(repository=repository)

    featureFlagSystem.add_channel("prod", {"new-ui": True}, [])
    featureFlagSystem.add_channel("staging", {"new-ui": False}, [])
    featureFlagSystem.add_channel("india-env", {}, ["prod"])

    print(featureFlagSystem.get_feature("prod", "new-ui"))
    print(featureFlagSystem.get_feature("staging", "new-ui"))
    print(featureFlagSystem.get_feature("india-env", "new-ui"))

if __name__ == '__main__':
    main()
