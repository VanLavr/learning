#include <stdio.h>
#include "My_deq.h"

int main(int argc, char const *argv[]) {
    
    dequeue dq;
    dq.create = CreateDequeue;
    int* MyNewDequeue;

    MyNewDequeue = dq.create(10);
    for (int i = 0; i < 10; ++i) {
        MyNewDequeue[i] = i + 1;
        printf("my deq[%d] = %d\n", i, MyNewDequeue[i]);
    }
    
    return 0;
}