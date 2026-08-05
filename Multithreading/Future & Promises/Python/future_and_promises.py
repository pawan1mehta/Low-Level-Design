from concurrent.futures.thread import ThreadPoolExecutor

def task():
    sum = 0
    for i in range(10):
        sum += i
    return sum

with ThreadPoolExecutor(max_workers=1) as executor:
    future = executor.submit(task)
    print("Waiting for task to be completed")
    print("result: ", future.result())
