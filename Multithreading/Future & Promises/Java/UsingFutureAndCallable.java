import java.util.concurrent.*;

public class UsingFutureAndCallable {
    public static void main(String[] args) throws InterruptedException, ExecutionException {
        ExecutorService threadPool = Executors.newSingleThreadExecutor();

        Callable<Integer> task = () -> {
            Thread.sleep(1000);
            return 10;
        };

        Future<Integer> res = threadPool.submit(task);

        System.out.println("waiting for result...");
        System.out.println("res " + res.get());

        threadPool.shutdown();
    }
}
