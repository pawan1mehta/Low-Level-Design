#include <iostream>
#include <vector>
#include <map>
#include <unordered_map>
#include <string>
#include <stdexcept>
#include <optional>

using namespace std;


class Channel {
private:
    string id;
    map<string, bool> features;
    vector<string> parents;

public:
    Channel(const string id, const map<string, bool> features, const vector<string> parents) {
        this->id = id;
        this->features = features;
        this->parents = parents;
    }

    void set(const string &feature_flag, bool val) {
        this->features[feature_flag] = val;
    }

    map<string, bool> getFeatures() {
        return this->features;
    }

    vector<string> getParents() {
        return this->parents;
    }
};

class InMemoryChannelRepository {
private:
    unordered_map<string, Channel> channels;

public:
    InMemoryChannelRepository() {
        channels = unordered_map<string, Channel>();
    }

    void add(const string& id, Channel channel) {
        channels.insert({id, channel});
    }

    Channel* get(string& id) {
        auto it = channels.find(id);
        if (it == channels.end()) {
            return nullptr;
        }
        return &(it->second);
    }

    void remove(string& id) {
        this->channels.erase(id);
    }
};

class FeatureFlagSystem {
private:
    InMemoryChannelRepository repository;

public:
    FeatureFlagSystem(InMemoryChannelRepository repository) {
        this->repository = repository;
    }

    void addChannel(string id, map<string, bool> features, vector<string> parents) {
        if (this->repository.get(id) != nullptr) {
            throw invalid_argument("Channel already exists: " + id);
        }
        this->repository.add(id, Channel(id, features, parents));
    }

    void deleteChannel(string id) {
        if (this->repository.get(id) == nullptr) {
            throw invalid_argument("Channel does not exist: " + id);
        }
        this->repository.remove(id);
    }

    void setFeatureFlag(string channelID, string featureFlag, bool val) {
        Channel* channel = this->repository.get(channelID);
        if (channel == nullptr) {
            throw invalid_argument("Channel does not exist: " + channelID);
        }
        channel->set(featureFlag, val);
    }

    optional<bool> findParentChain(string channelID, string featureFlag) {
        Channel* channel = repository.get(channelID);
        if (channel == nullptr) {
            return nullopt;
        }

        const auto& features = channel->getFeatures();
        auto it = features.find(featureFlag);
        if (it != features.end()) {
            return it->second;
        }

        for (const auto& parentId : channel->getParents()) {
            auto inherited = findParentChain(parentId, featureFlag);
            if (inherited.has_value()) {
                return inherited;
            }
        }

        return nullopt;
    }

    bool getFeatureFlag(string channelID, string featureFlag) {
        Channel* channel = repository.get(channelID);
        if (channel == nullptr) {
            throw invalid_argument("Channel does not exist: " + channelID);
        }

        const auto& features = channel->getFeatures();
        auto it = features.find(featureFlag);
        if (it != features.end()) {
            return it->second;
        }

        for (const auto& parentID : channel->getParents()) {
            auto inherited = findParentChain(parentID, featureFlag);
            if (inherited.has_value()) {
                return inherited.value();
            }
        }

        return false;
    }
};


int main() {
    InMemoryChannelRepository repository;
    FeatureFlagSystem system(repository);

    system.addChannel("prod", {{"new-ui", true}}, {});
    system.addChannel("staging", {{"new-ui", false}}, {});
    system.addChannel("india-env", {}, {"prod"});

    cout << "prod: new-ui: " << system.getFeatureFlag("prod", "new-ui") << endl;
    cout << "staging: new-ui: " << system.getFeatureFlag("staging", "new-ui") << endl;
    cout << "india-env: new-ui: " << system.getFeatureFlag("india-env", "new-ui") << endl;

    return 0;
}