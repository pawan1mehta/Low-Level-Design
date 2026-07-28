#include <iostream>
#include <thread>
#include <queue>
#include <mutex>
#include <condition_variable>
#include <functional>
#include <vector>
#include <chrono>
#include <string>

using namespace std;

class ThreadPool {
private:
    vector<thread> workers;
    queue<function<void()>> taskQueue;
    mutex mtx;
    condition_variable cv;
    bool stop = false;

public:
    ThreadPool(int numThreads) {
        for (int i = 0; i < numThreads; i++) {
            workers.emplace_back([this, i] {
                while (true) {
                    function<void()> task;

                    {
                        unique_lock<mutex> lock(mtx);

                        cv.wait(lock, [this] {
                           return !taskQueue.empty() || stop;
                        });

                        if (stop && taskQueue.empty()) {
                            break;
                        }

                        task = taskQueue.front();
                        taskQueue.pop();
                    }

                    task();
                }
                cout << "Worker-" << i + 1 << " shut down"<<  endl;
            });
        }
    }

    void submit(function<void()> task) {
        {
            unique_lock<mutex> lock(mtx);
            taskQueue.push(task);
        }
        cv.notify_one();
    }

    void shut_down() {
        {
            unique_lock<mutex> lock(mtx);
            stop = true;
        }

        cv.notify_all();

        for (thread& worker : workers) {
            worker.join();
        }
    }
};

void task(const string& taskID) {
    cout << "\n Doing task '" << taskID << "'"<< endl;
    this_thread::sleep_for(chrono::seconds(1));
    cout << "\n task '" << taskID << "' completed" << endl;
}

int main() {
    ThreadPool thread_pool(3);

    thread_pool.submit([] { task("taskID1"); });
    thread_pool.submit([] { task("taskID2"); });
    thread_pool.submit([] { task("taskID3"); });
    thread_pool.submit([] { task("taskID4"); });
    thread_pool.submit([] { task("taskID5"); });
    thread_pool.submit([] { task("taskID6"); });
    thread_pool.submit([] { task("taskID7"); });

    thread_pool.shut_down();

    return 0;
}
