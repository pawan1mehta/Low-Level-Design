package main

type Application struct {
	button   Button
	checkbox Checkbox
}

func NewApplication(factory GUIFactory) *Application {
	return &Application{
		factory.createButton(),
		factory.createCheckbox(),
	}
}

func (app *Application) Paint() {
	app.button.Paint()
	app.checkbox.Paint()
}
