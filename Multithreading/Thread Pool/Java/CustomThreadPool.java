import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;

class Worker extends Thread {
    private final BlockingQueue<Runnable> queue;

    public Worker(BlockingQueue<Runnable> queue) {
        this.queue = queue;
    }

    @Override
    public void run() {
        try {
            while (!Thread.currentThread().isInterrupted()) {
                Runnable task = queue.take();
                task.run();
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
    }
}

class ThreadPool {
    private final BlockingQueue<Runnable> taskQueue;
    private final Worker[] workers;

    public ThreadPool(int numTasks) {
        this.taskQueue = new LinkedBlockingQueue<>();
        this.workers = new Worker[numTasks];

        for(int i = 0; i < numTasks; i++) {
            workers[i] = new Worker(taskQueue);
            workers[i].setName("Worker - " + (i + 1));
            workers[i].start();
        }
    }

    public void submitTask(Runnable task) {
        try {
            taskQueue.put(task);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
    }

    public void shutDown() {
        for(Worker worker : workers) {
            worker.interrupt();
        }
    }
}

public class CustomThreadPool {
    public static void main(String[] args) throws InterruptedException {
        ThreadPool threadPool = new ThreadPool(3);

        for(int i = 1; i <= 10; i++) {
            final int taskID = i;
            threadPool.submitTask(()-> {
                System.out.println("Executing Task " + taskID + " via " + Thread.currentThread().getName());
                try {
                    Thread.sleep(500);
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                }
            });
        }

        Thread.sleep(3000);

        threadPool.shutDown();
    }
}