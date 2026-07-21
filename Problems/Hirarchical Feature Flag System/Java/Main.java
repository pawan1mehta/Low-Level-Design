import Repository.ChannelRepository;
import Repository.InMemoryChannelRepository;

import java.util.HashMap;
import java.util.Map;

public class Main {

    public static void main(String[] args) {
        ChannelRepository repository = new InMemoryChannelRepository();

        FeatureFlagSystem featureFlagSystem = new FeatureFlagSystem(repository);

        Map<String, Boolean> prodFeatures = new HashMap<>();
        prodFeatures.put("new-ui", true);
        featureFlagSystem.addChannel(
                "prod",
                prodFeatures,
                new String[]{}
        );

        Map<String, Boolean> stagingFeatures = new HashMap<>();
        stagingFeatures.put("new-ui", false);
        featureFlagSystem.addChannel(
                "staging",
                stagingFeatures,
                new String[]{}
        );

        Map<String, Boolean> indiaEnvFeatures = new HashMap<>();
        featureFlagSystem.addChannel(
                "india-env",
                indiaEnvFeatures,
                new String[]{"prod"}
        );

        System.out.println("prod: new-ui: " + featureFlagSystem.getFeature("prod", "new-ui"));
        System.out.println("staging: new-ui: " + featureFlagSystem.getFeature("staging", "new-ui"));
        System.out.println("india-env: new-ui: " + featureFlagSystem.getFeature("india-env", "new-ui"));
    }
}