
abstract class ShapeFactory {
    public abstract Shape createShape();

    public void render() {
        Shape shape = createShape();
        shape.draw();
    }
}