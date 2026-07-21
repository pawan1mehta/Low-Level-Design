package Repository;

import Channel.Channel;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class InMemoryChannelRepository implements ChannelRepository {
    private Map<String, Channel> channels;

    public InMemoryChannelRepository() {
        this.channels = new HashMap<>();
    }

    @Override
    public List<Channel> readAll() {
        List<Channel> channelsList = new ArrayList<>();
        for(Channel channel : channels.values()) {
            channelsList.add(channel);
        }
        return channelsList;
    }

    @Override
    public void add(String key, Channel channel) {
        channels.put(key, channel);
    }

    @Override
    public Channel get(String id) {
        return channels.get(id);
    }

    @Override
    public void delete(String id) {
        channels.remove(id);
    }
}
