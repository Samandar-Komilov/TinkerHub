#include <iostream>

using namespace std;


typedef struct {
    int id;
    int burstTime;
    int arrivalTime;
    int waitingTime;
    int turnaroundTime;
    bool is_completed;
} Process;


void shortest_job_first(Process processes[], int n);
void sort_job_bursts(Process* processes, int n);


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

    shortest_job_first(processes, n);

    return 0;
}


void shortest_job_first(Process processes[], int n){
    int currentTime = 0;
    int totalWaitingTime = 0, totalTurnaroundTime = 0;

    sort_job_bursts(processes, n);

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

    int i=0;
    while (i<n){

        i++;
        currentTime++;
    }

    cout << "Average Waiting time: " << (totalWaitingTime / n) << endl;
    cout << "Average Turnaround time: " << (totalTurnaroundTime / n) << endl;
}


void sort_job_bursts(Process* processes, int n){
    for (int step = 0; step < n - 1; ++step) {
        for (int i = 0; i < n - step - 1; ++i) {
            if (processes[i].burstTime > processes[i + 1].burstTime) {
                Process temp = processes[i];
                processes[i] = processes[i + 1];
                processes[i + 1] = temp;
            }
        }
    }
}