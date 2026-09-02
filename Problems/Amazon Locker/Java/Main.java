import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        List<Compartment> compartmentList = new ArrayList<>();

        for(int i = 0; i < 10; i++) {
            compartmentList.add(new Compartment(Size.SMALL, false));
        }
        for(int i = 0; i < 10; i++) {
            compartmentList.add(new Compartment(Size.MEDIUM, false));
        }
        for(int i = 0; i < 10; i++) {
            compartmentList.add(new Compartment(Size.LARGE, false));
        }

        Locker locker = new Locker(compartmentList);

        String accessCode = locker.depositPackage(Size.LARGE);

        locker.pickup(accessCode);

        locker.openExpiredCompartment();
    }
}
