#include "controller.h"
#include <stdlib.h>
#include <stdio.h>

int main() {
    int stop = 0;
    unsigned int tid = 0;
    char user[20];
    double amount = 0.0;
    while (!stop) {
        printf("Enter transaction (tid user amount): ");
        scanf("%u %19s %lf", &tid, user, &amount);
        Transaction t = {tid, user, amount};
        process_transaction(&t);
        
        printf("Stop? (0/1): ");
        scanf("%d", &stop);
    }
}