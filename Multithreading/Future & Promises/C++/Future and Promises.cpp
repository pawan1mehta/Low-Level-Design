#include <iostream>
#include <future>
#include <chrono>
#include <thread>

using namespace std;

int async_task() {
    cout << "task running..." << endl;
    this_thread::sleep_for(chrono::seconds(2));
    return 10;
}

int main() {

    future<int> result = async(launch::async, async_task);

    int data = result.get();
    cout<< "result: " << data << endl;

    return 0;
}
