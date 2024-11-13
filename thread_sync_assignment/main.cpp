#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <queue>
#include <chrono>

const int MAX_ROOMS = 1; // Assuming one room for simplicity
const int STUDY_TIME = 2; // Time a student spends in the room (in seconds)

std::mutex room_mutex;
std::condition_variable cv;
std::queue<int> waiting_queue;

void student(int id) {
    // Add student to the queue
    {
        std::unique_lock<std::mutex> lock(room_mutex);
        waiting_queue.push(id);
    }

    // Wait for the room to be available in a FIFO manner
    std::unique_lock<std::mutex> lock(room_mutex);
    cv.wait(lock, [&]() { return waiting_queue.front() == id; });

    // Reserve the room
    std::cout << "Student " << id << " is reserving the room.\n";

    // Simulate studying in the room
    std::this_thread::sleep_for(std::chrono::seconds(STUDY_TIME));
    std::cout << "Student " << id << " has finished using the room.\n";

    // Release the room
    waiting_queue.pop();
    lock.unlock();
    cv.notify_all();
}

int main() {
    int num_students = 5; // Number of students for simulation
    std::vector<std::thread> students;

    // Create student threads
    for (int i = 1; i <= num_students; ++i) {
        students.emplace_back(student, i);
    }

    // Wait for all student threads to finish
    for (auto &t : students) {
        t.join();
    }

    return 0;
}
