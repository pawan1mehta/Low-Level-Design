https://enginebogie.com/public/question/design-a-hierarchical-feature-flagging-system/4374


======================================== Problem Statement ==============================================================

You need to build a feature flag system where feature flags are settings like:
"new_ui" = true/false
"beta_mode" = true/false

These settings are stored inside channels (think: environments/segments like prod, staging, india_users, etc.).

What is a “channel”?
    Each channel has:
        A unique ID (example: "prod")
        A dictionary/map of feature flags → boolean values
            Example: { "new_ui": true, "dark_mode": false }
        A list of parent channels (optional)
            Example: "india_users" might have parent "prod"

A channel can inherit features from other channels:
    If a feature is defined in a channel, return the value for that feature
    If a feature is not defined in a channel, it should inherit the value from its closest parent channel
    If a feature is not defined in any parent channel, it should default to false

Implement the following functions for the feature flagging system:
    create_channel(channel_id, features, parents)
    get_feature(channel_id, feature_name)
    set_feature(channel_id, feature_name, value)
    delete_channel(channel_id)

Followups:
    Suppose we deploy this system in prod; we don't have as many reads as writes. How would you optimise the reads?
    What happens if you update a feature flag value for a channel? How do you ensure that the lookups work at scale?
    If the system is read-heavy (many more lookups than updates), how would you optimise the get_feature performance?

======================================== Solution  ==============================================================


Requirements:
    - The users should be able to create/delete the channel
    - The users should be able to add/update the feature flags
    - The channel can inherit features from other channels

Entities & Relationships:

    Entities:
     - Channel.Channel
     - FeatureFlagSystem

    Relationships
        FeatureFlagSystem  <|----- composed of --- Channel.Channel

Class Design:

    FeatureFlagSystem
        - repository: ChannelRepository

        + addChannel(channelID, features, parents)
        + deleteChannel(channelID)
        + setFeature(channelID, featureFlag, value)
        + getFeature(channelID, featureFlag)

    Channel.Channel:
        - id: string
        - features: map{featureFlag, value}
        - parents: channelIDs

        + addFeatures(features: map{key, value})
        + addFeature(featureFlag, value)
        + getFeature(featureFlag)
        - setFeature(featureFlag, value)

    ChannelRepository:
        + save
        + find
        + delete
