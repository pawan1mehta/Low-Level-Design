import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class UsingExecutorService {
    public static void main(String[] args) throws InterruptedException {
        ExecutorService threadPool = Executors.newFixedThreadPool(3);

        for(int i = 0; i <= 10; i++) {
            final int taskID = i;
            threadPool.submit(()-> {
                System.out.println("Executing Task " + taskID + " via " + Thread.currentThread().getName());
                try {
                    Thread.sleep(500);
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                }
            });
        }

        threadPool.shutdown();
    }
}