import Channel.Channel;
import Repository.ChannelRepository;

import java.util.Map;

public class FeatureFlagSystem {
    private ChannelRepository repository;

    public FeatureFlagSystem(ChannelRepository repository) {
        this.repository = repository;
    }

    void addChannel(String id, Map<String, Boolean> features, String[] parents) {
        if(repository.get(id) != null) {
            throw new IllegalArgumentException("Can not add, channel already exits");
        }
        Channel channel = new Channel(id, features, parents);
        repository.add(id, channel);
    }

    void deleteChannel(String id) {
        if(repository.get(id) == null) {
            throw new IllegalArgumentException("Channel not found: " + id);
        }
        repository.delete(id);
    }

    void setFeature(String channelID, String featureFlag, Boolean value) {
        Channel channel = repository.get(channelID);
        channel.set(featureFlag, value);
    }

    private Boolean getFeatureFromParentIfNotDefined(String channelID, String featureFlag) {
        Channel channel = repository.get(channelID);

        if(channel == null) {
            throw new IllegalArgumentException("Channel not found: " + channelID);
        }

        Map<String, Boolean> features = channel.getFeatures();

        if(features.containsKey(featureFlag)) {
            return features.get(featureFlag);
        }

        for(String id : channel.getParents()) {
            Boolean inherited = getFeatureFromParentIfNotDefined(id, featureFlag);
            if(inherited != null) {
                return inherited;
            }
        }

        return null;
    }

    Boolean getFeature(String channelID, String featureFlag) {
        Channel channel = repository.get(channelID);

        if(channel == null) {
            throw new IllegalArgumentException("Channel not found: " + channelID);
        }

        if(channel.getFeatures().containsKey(featureFlag)) {
            return channel.getFeatures().get(featureFlag);
        }

        for(String id : channel.getParents()) {
            Boolean inherited = getFeatureFromParentIfNotDefined(id, featureFlag);
            if(inherited != null) {
                return inherited;
            }
        }

        return false;
    }
}