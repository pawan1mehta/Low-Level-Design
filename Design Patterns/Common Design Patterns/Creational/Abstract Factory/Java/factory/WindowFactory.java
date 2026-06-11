package factory;

import model.Button.Button;
import model.Checkbox.Checkbox;
import model.Button.WindowOSButton;
import model.Checkbox.WindowOSCheckbox;

public class WindowFactory implements GUIFactory {

    public Button createButton() {
        return new WindowOSButton();
    }

    public Checkbox createCheckbox() {
        return new WindowOSCheckbox();
    }
}
