import queue
import threading
import time


class Worker(threading.Thread):
    def __init__(self, worker_id: int, taskQueue: queue.Queue):
        super().__init__(daemon=False)
        self.worker_id = worker_id
        self.taskQueue = taskQueue

    def run(self):
        while True:
            func, args, kwargs = self.taskQueue.get()

            if func is None:
                self.taskQueue.task_done()
                break

            func(*args, **kwargs)
            self.taskQueue.task_done()

class ThreadPool:
    def __init__(self, numTasks: int):
        self.taskQueue = queue.Queue()
        self.workers = []

        for i in range(numTasks):
            worker = Worker(worker_id=i, taskQueue=self.taskQueue)
            worker.start()
            self.workers.append(worker)

    def submit(self, func, *args, **kwargs):
        self.taskQueue.put((func, args, kwargs))

    def shut_down(self):
        self.taskQueue.join()

        for _ in self.workers:
            self.taskQueue.put((None, [], {}))

        for worker in self.workers:
            worker.join()

def task(taskID: int):
    print(f"Doing task '{taskID}'")
    time.sleep(1)
    print(f"task '{taskID}' completed")

threadPool = ThreadPool(numTasks=3)

threadPool.submit(task, "taskID1")
threadPool.submit(task, "taskID2")
threadPool.submit(task, "taskID3")
threadPool.submit(task, "taskID4")
threadPool.submit(task, "taskID5")
threadPool.submit(task, "taskID6")
threadPool.submit(task, "taskID7")

threadPool.shut_down()