import time
from concurrent.futures import ThreadPoolExecutor

def task(taskID: str):
    print(f"\n Doing task '{taskID}'")
    time.sleep(1)
    print(f"\n task '{taskID}' completed")

with ThreadPoolExecutor(max_workers=3) as executor:
    executor.submit(task, "taskID1")
    executor.submit(task, "taskID2")
    executor.submit(task, "taskID3")
    executor.submit(task, "taskID4")
    executor.submit(task, "taskID5")
    executor.submit(task, "taskID6")
    executor.submit(task, "taskID7")