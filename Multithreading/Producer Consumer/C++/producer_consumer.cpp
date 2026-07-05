#include <iostream>
#include <thread>
#include <mutex>
#include <queue>

using namespace std;

mutex mtx;
queue<int> buffer;
condition_variable cv;
bool done = false;

void producer() {
    for (int i = 0; i < 10; i++) {
        {
            lock_guard<mutex> lock(mtx);
            buffer.push(i);
            cout << "Produced: " << i << endl;
        }
        cv.notify_one();
    }
    {
        lock_guard<mutex> lock(mtx);
        done = true;
    }
    cv.notify_all();
}

void consumer() {
    while (true) {
        unique_lock<mutex> lock(mtx);

        cv.wait(lock, [] {
            return !buffer.empty() || done;
        });

        if (buffer.empty() && done) {
            break;
        }

        int value = buffer.front();
        buffer.pop();
        lock.unlock();

        cout << "Consumed: " << value << endl;
    }
}

int main() {
    thread t1(producer);
    thread t2(consumer);

    t1.join();
    t2.join();

    return 0;
}