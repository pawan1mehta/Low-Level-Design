package factory;

import model.Button.Button;
import model.Checkbox.Checkbox;
import model.Button.MacOSButton;
import model.Checkbox.MacOSCheckbox;

public class MacOSFactory implements GUIFactory {

    public Button createButton() {
        return new MacOSButton();
    }

    public Checkbox createCheckbox() {
        return new MacOSCheckbox();
    }
}
