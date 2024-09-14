#include <stdio.h>
#include <stdlib.h>
#include <string.h>


typedef struct PCB {
    // Basic PCB structure
    int pid;
    char state[20];
    struct PCB *next;
} PCB;

PCB* create_pcb(int pid, const char* state){
    PCB* new_pcb = (PCB*) malloc(sizeof(PCB));
    new_pcb->pid = pid;
    new_pcb->next = NULL;

    return new_pcb;
}


// Queues using Linked Lists
typedef struct Queue{
    PCB* front;
    PCB* rear;
} Queue;

void init_queue(Queue* q){
    // A new Queue is initially empty
    q->front = q->rear = NULL;
}

void enqueue(Queue* q, PCB* pcb){
    // Adding a PCB to a queue
    if (q->rear == NULL){
        q->front = q->rear = pcb;
        return;
    }
    q->rear->next = pcb;
    q->rear = pcb;
}

PCB* dequeue(Queue* q){
    // Removing a PCB from queue
    if (q->front == NULL){
        return NULL;
    }
    PCB* tmp = q->front;
    q->front = q->front->next;

    if (q->front == NULL){
        q->rear = NULL;
    }

    return tmp;
}


/*
    PCB scheduling simulation

    Job queue - keeps all processes in the system
    Ready queue - keeps a set of ready and waiting to execute processes.
                - A new process is always put in this queue.
    Device queue - keeps a set of processes blocked due to I/O
    */
void set_ready(Queue* job_queue, Queue* ready_queue){
    PCB* process = dequeue(job_queue);
    if (process != NULL){
        strcpy(process->state, "Ready");
        enqueue(ready_queue, process);
        printf("Process %d moved from Job Queue to Ready Queue.\n", process->pid);
    }
    else{
        printf("No active processes exist in Job Queue. Aborting...");
    }
}

void set_blocked(Queue* ready_queue, Queue* device_queue){
    PCB* process = dequeue(ready_queue);
    if (process != NULL){
        strcpy(process->state, "Blocked");
        enqueue(device_queue, process);
        printf("Process %d moved from Ready Queue to Device Queue (Blocked).\n", process->pid);
    }
}


int main(){
    Queue job_queue, ready_queue, device_queue;

    init_queue(&job_queue);
    init_queue(&ready_queue);
    init_queue(&device_queue);

    // Create some PCBs (processes)
    PCB* p1 = create_pcb(1, "New");
    PCB* p2 = create_pcb(2, "New");
    PCB* p3 = create_pcb(3, "New");

    // Add them to the job queue
    enqueue(&job_queue, p1);
    enqueue(&job_queue, p2);
    enqueue(&job_queue, p3);

    // Simulate process transitions
    set_ready(&job_queue, &ready_queue);
    set_blocked(&ready_queue, &device_queue);

    return 0;
}
