import factory.GUIFactory;
import factory.MacOSFactory;
import factory.WindowFactory;

public class Driver {
    
    public static void main(String[] args) {
        Application app = null;
        GUIFactory factory = null;

        String osName = "mac";

        if("mac".equals(osName)) {
            factory = new MacOSFactory();
        } else {
            factory = new WindowFactory();
        }

        app = new Application(factory);
        app.paint();
    }
}
