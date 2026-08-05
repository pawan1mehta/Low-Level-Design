
class Result {
    private int result;
    private boolean ready = false;

    public synchronized void setResult(int result) {
        if(ready) { return; }

        this.result = result;
        this.ready = true;

        notifyAll();
    }

    public synchronized int get() throws InterruptedException {
        while (!ready) {
            wait();
        }
        return result;
    }
}

public class Custom {
    public static void main(String[] args) throws InterruptedException {
        final Result result = new Result();

        Runnable task = () -> {
            try {
                Thread.sleep(1000);
            } catch (InterruptedException ex) {
                Thread.currentThread().interrupt();
            }
            result.setResult(10);
        };

        new Thread(task).start();

        System.out.println("waiting for result...");
        System.out.println("res " + result.get());
    }
}
