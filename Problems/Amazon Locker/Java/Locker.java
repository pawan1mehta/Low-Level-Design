import java.time.LocalDateTime;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Random;

public class Locker {
    private List<Compartment> compartments;
    private Map<String, AccessToken> accessTokenMap;

    public Locker(List<Compartment> compartments) {
        this.compartments = compartments;
        this.accessTokenMap = new HashMap<>();
    }

    public String depositPackage(Size size) {
        Compartment compartment = findUnOccupiedCompartment(size);
        if(compartment == null) {
            throw new RuntimeException("no available compartment of size " + size);
        }

        compartment.open();
        compartment.markOccupied();

        AccessToken accessToken = generateAccessToken(compartment);
        accessTokenMap.put(accessToken.getCode(), accessToken);

        return accessToken.getCode();
    }

    public void pickup(String accessCode) {
        if("".equals(accessCode)) {
            throw new IllegalArgumentException("invalid accessCode!");
        }

        AccessToken accessToken = accessTokenMap.get(accessCode);
        if(accessToken.isExpired()) {
            throw new RuntimeException("access token has expired!");
        }

        Compartment compartment = accessToken.getCompartment();
        compartment.open();

        clearDeposit(accessToken);
    }

    private void clearDeposit(AccessToken accessToken) {
        Compartment compartment = accessToken.getCompartment();
        compartment.markFree();
        accessTokenMap.remove(accessToken.getCode());
    }

    public void openExpiredCompartment() {
        for(AccessToken accessToken : accessTokenMap.values()) {
            if(accessToken.isExpired()) {
                accessToken.getCompartment().open();
            }
        }
    }

    private Compartment findUnOccupiedCompartment(Size size) {
        for(Compartment compartment : compartments) {
            if(!compartment.isOccupied()) {
                return compartment;
            }
        }
        return null;
    }

    private AccessToken generateAccessToken(Compartment compartment) {
        Random random = new Random();
        String accessCode = String.format("%06d", random.nextInt(1000000));
        LocalDateTime expiration = LocalDateTime.now().plusDays(7);
        return new AccessToken(accessCode, expiration, compartment);
    }
}
