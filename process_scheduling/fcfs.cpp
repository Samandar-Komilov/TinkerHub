/*
*** FCFS - First Come First Served Algorithm (Non-preemptive)
*/

#include <iostream>
#include <unistd.h>
#include <vector>
#include <queue>


using namespace std;


typedef struct {
    int id;
    int burstTime;
    int arrivalTime;
    int waitingTime;
    int turnaroundTime;
} Process;

void firstComeFirstServed(Process processes[], int n);



int main(){
    int n;
    cout << "Enter number of processes: ";
    cin >> n;

    Process processes[n];

    for (int i=0; i<n; i++){
        processes[i].id = i+1;
        cout << "Enter arrival time for Process " << processes[i].id << ": ";
        cin >> processes[i].arrivalTime;
        cout << "Enter burst time for Process " << processes[i].id << ": ";
        cin >> processes[i].burstTime;
    }

    firstComeFirstServed(processes, n);

    return 0;
}


void firstComeFirstServed(Process processes[], int n){
    int currentTime = 0;
    int totalWaitingTime = 0;
    int totalTurnaroundTime = 0;

    for (int i=0; i<n; i++){
        if (currentTime < processes[i].arrivalTime){
            currentTime = processes[i].arrivalTime;
        }
        processes[i].waitingTime = currentTime - processes[i].arrivalTime;
        processes[i].turnaroundTime = processes[i].waitingTime + processes[i].burstTime;

        currentTime += processes[i].burstTime;

        cout << "Process " << processes[i].id << " | Arrival Time " << processes[i].arrivalTime << " | Burst Time "
            << processes[i].burstTime << " | Waiting time " << processes[i].waitingTime << endl;

        totalWaitingTime += processes[i].waitingTime;
        totalTurnaroundTime += processes[i].turnaroundTime;
    }

    cout << "Average Waiting time: " << (totalWaitingTime / n) << endl;
    cout << "Average Turnaround time: " << (totalTurnaroundTime / n) << endl;
}