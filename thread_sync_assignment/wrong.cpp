#include <iostream>
#include <thread>
#include <vector>
#include <map>
#include <string>
#include <chrono>
#include <mutex>

#define ROOMS_N 5
#define RESERVE_TIME_IN_SECONDS 5
#define QUEUE_LEN 100


using namespace std;


typedef struct {
    string name;
    int reserveTime;
} Student;

vector<Student> Queue(QUEUE_LEN);
int current_queue_length = QUEUE_LEN;

mutex mtx;


int main(){

    int i, j;
    for (i=0; i<QUEUE_LEN; i++){
        string student_name = "S" + to_string(i);
        Queue[i].name = student_name;
        Queue[i].reserveTime = i+1;
    }

    for (i=0; i<QUEUE_LEN; i++){
        cout << Queue[i].name << " " << Queue[i].reserveTime << endl;
    }

    thread* Rooms[ROOMS_N];
    for (j=0; i<ROOMS_N; j++){
        Rooms[j] = new thread(thread_func, j);
    }

    return 0;
}


void thread_func(int thread_num){
    if (current_queue_length == 0)
        return
    
    mtx.lock();
    Student accepted_student = Queue[--current_queue_length];
    cout << "Room " << thread_num << " accepted Student " << accepted_student.name << " for " << accepted_student.reserveTime << " seconds." << endl;
    mtx.unlock();
}