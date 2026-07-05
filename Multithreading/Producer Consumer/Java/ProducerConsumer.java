import java.util.*;

class BoundedBufferQueue<T> {
    private final Queue<T> queue;
    private final int capacity;

    public BoundedBufferQueue(int capacity) {
        this.queue = new LinkedList<>();
        this.capacity = capacity;
    }

    public synchronized void enqueue(T item) throws InterruptedException {
        while(queue.size() == capacity) {
            wait();
        }
        queue.add(item);
        notifyAll();
    }

    public synchronized T dequeue() throws InterruptedException {
        while (queue.isEmpty()) {
            wait();
        }
        T item = queue.poll();
        notifyAll();
        return item;
    }

    public synchronized int size() {
        return queue.size();
    }
}

public class ProducerConsumer {

    public static void main(String[] args) throws InterruptedException {
        BoundedBufferQueue queue = new BoundedBufferQueue(3);
        final int STOP = -1;

        Thread t1 = new Thread(()-> {
            try {
                for(int i = 0; i < 10; i++) {
                    queue.enqueue(i);
                    System.out.println("Produced: " + i);
                }
                queue.enqueue(STOP);
            } catch (InterruptedException exception) {
                Thread.currentThread().interrupt();
            }
        });

        Thread t2 = new Thread(()->{
            try {
                while (true) {
                    int item = (int) queue.dequeue();
                    if(item == STOP) {
                        break;
                    }
                    System.out.println("Consumed: " + item);
                }
            } catch (InterruptedException exception) {
                Thread.currentThread().interrupt();
            }
        });

        t1.start();
        t2.start();

        t1.join();
        t2.join();
    }
}
