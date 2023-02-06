#include <stdio.h>
#include "My_deq.h"

int main(int argc, char const *argv[]) {
    
    // initializing all atributes...
    dequeue dq;
    dq.create = CreateDequeue;
    dq.addRight = AddToRight;
    dq.addLeft = AddToLeft;
    dq.delRight = DeleteFromRight;
    dq.delLeft = DeleteFromLeft;

    // initializing deq and capacity...
    int* MyNewDequeue;
    int cap = 10;
    dq.capacity = cap;


    // tests...
    printf("Dequeue before tests: \n");
    MyNewDequeue = dq.create(dq.capacity);
    for (int i = 0; i < dq.capacity; ++i) {
        MyNewDequeue[i] = i + 1;
        printf("DEQ[%d] = %d\n", i, MyNewDequeue[i]);
    }

    // adding...
    printf("adding tests: \n");
    dq.addRight(13, MyNewDequeue, dq.capacity);
    dq.addRight(11, MyNewDequeue, dq.capacity);
    dq.addRight(454, MyNewDequeue, dq.capacity);
    dq.addLeft(454, MyNewDequeue, dq.capacity);
    dq.addLeft(11, MyNewDequeue, dq.capacity);
    dq.addLeft(13, MyNewDequeue, dq.capacity);
    for (int i = 0; i < dq.capacity; ++i) {
        printf("DEQ[%d] = %d\n", i, MyNewDequeue[i]);
    }

    // deleting...
    printf("deleting tests: \n");
    dq.delRight(MyNewDequeue, dq.capacity);
    dq.delLeft(MyNewDequeue, dq.capacity);
    for (int i = 0; i < dq.capacity; ++i) {
        printf("DEQ[%d] = %d\n", i, MyNewDequeue[i]);
    }

    // adding and deleting...
    printf("mixed adding and deleting tests: \n");
    dq.addRight(454, MyNewDequeue, dq.capacity);
    dq.delRight(MyNewDequeue, dq.capacity);
    dq.delLeft(MyNewDequeue, dq.capacity);
    dq.addLeft(454, MyNewDequeue, dq.capacity);
    
    return 0;
}