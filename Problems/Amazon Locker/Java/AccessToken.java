import java.time.LocalDateTime;

public class AccessToken {
    private String code;
    private LocalDateTime expiration;
    private Compartment compartment;

    public AccessToken(String code, LocalDateTime expiration, Compartment compartment) {
        this.code = code;
        this.expiration = expiration;
        this.compartment = compartment;
    }

    public String getCode() {
        return this.code;
    }

    public boolean isExpired() {
        return LocalDateTime.now().isAfter(expiration);
    }

    public Compartment getCompartment() {
        return this.compartment;
    }
}
