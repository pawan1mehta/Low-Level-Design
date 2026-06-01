package factory;

import model.Button.Button;
import model.Checkbox.Checkbox;

public interface GUIFactory {

    Button createButton();
    Checkbox createCheckbox();
}
