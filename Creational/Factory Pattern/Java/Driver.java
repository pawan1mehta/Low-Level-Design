
public class Driver {
    public static void main(String[] args) {
        ShapeFactory shapeFactory;

        shapeFactory = new CircleFactory();
        shapeFactory.render();

        shapeFactory = new RectangleFactory();
        shapeFactory.render();
    }
}