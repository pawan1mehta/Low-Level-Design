package Channel;

import Repository.ChannelRepository;

import java.util.Map;

public class Channel {
    private String id;
    private Map<String, Boolean> features;
    private String[] parents;

    public Channel(String id, Map<String, Boolean> features, String[] parents) {
        this.id = id;
        this.features = features;
        this.parents = parents;
    }

    public void set(String featureFlag, Boolean value) {
        features.put(featureFlag, value);
    }

    public Map<String, Boolean> getFeatures() {
        return features;
    }

    public String[] getParents() {
        return parents;
    }
}