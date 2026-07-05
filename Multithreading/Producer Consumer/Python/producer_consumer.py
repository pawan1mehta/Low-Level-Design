import threading
import time
from collections import deque

class BoundedBuffer:
    def __init__(self, capacity):
        self.capacity = capacity
        self.buffer = deque()
        self.cv = threading.Condition()
        self.done = False

    def put(self, value):
        with self.cv:
            while len(self.buffer) == self.capacity:
                self.cv.wait()
            self.buffer.append(value)
            self.cv.notify_all()

    def get(self):
        with self.cv:
            while not self.buffer and not self.done:
                self.cv.wait()

            if not self.buffer and self.done:
                return None

            item = self.buffer.popleft()
            self.cv.notify_all()

            return item

    def close(self):
        with self.cv:
            self.done = True
            self.cv.notify_all()

def producer(buffer: BoundedBuffer):
    for i in range(100):
        buffer.put(i)
        print("Produced: \n", i)
        time.sleep(1)
    buffer.close()

def consumer(buffer: BoundedBuffer):
    while True:
        item = buffer.get()
        if item is None:
            break
        print("Consumed: \n", item)
        time.sleep(1)

queue = BoundedBuffer(capacity=10)

producerThread = threading.Thread(target=producer, args=(queue,))
consumerThread = threading.Thread(target=consumer, args=(queue,))

producerThread.start()
consumerThread.start()

producerThread.join()
consumerThread.join()
