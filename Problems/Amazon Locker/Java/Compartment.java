
public class Compartment {
    private Size size;
    private boolean occupied;

    public Compartment(Size size, boolean occupied) {
        this.size = size;
        this.occupied = occupied;
    }

    public Size getSize() {
        return this.size;
    }

    public void open() {

    }

    public boolean isOccupied() {
        return this.occupied;
    }

    public void markOccupied() {
        this.occupied = true;
    }

    public void markFree() {
        this.occupied = false;
    }
}
