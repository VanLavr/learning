// printf("\n");

#include <stdio.h>
#include <stdlib.h>
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
    int* MyNewDequeue = NULL;
    int cap = 10;


    // tests...
    printf("Dequeue before tests: \n");
    MyNewDequeue = dq.create(cap);
    for (int i = 0; i < cap; i++) {
        MyNewDequeue[i] = i + 1;
        printf("DEQ[%d] = %d\n", i, MyNewDequeue[i]);
    }
    printf("\ndeq cap: %d\n", cap);
    printf("\n");

    // adding...
    printf("adding tests: \n");
    printf("added three elements at the end (13, 11, 454) and at the beginning (454, 11, 13)... \n");
    //dq.addRight(13, MyNewDequeue, &cap);
    //dq.addRight(11, MyNewDequeue, &cap);
    //dq.addRight(454, MyNewDequeue, &cap);
    dq.addLeft(454, MyNewDequeue, &cap);
    //dq.addLeft(11, MyNewDequeue, &cap);
    //dq.addLeft(13, MyNewDequeue, &cap);
    for (int i = 0; i < cap; i++) {
        printf("DEQ[%d] = %d\n", i, MyNewDequeue[i]);
    }
    printf("\ndeq cap: %d\n", cap);
    printf("\n");

/*


    // deleting...
    printf("deleting tests: \n");
    printf("deleted two elements, one from the end and one from the begining... \n");
    dq.delRight(MyNewDequeue, &dq.capacity);
    dq.delLeft(MyNewDequeue, &dq.capacity);
    for (int i = 0; i < dq.capacity; i++) {
        printf("DEQ[%d] = %d\n", i, MyNewDequeue[i]);
    }
    printf("\ndeq cap: %d\n", dq.capacity);
    printf("\n");

    // adding and deleting...
    printf("mixed adding and deleting tests: \n");
    dq.addRight(454, MyNewDequeue, &dq.capacity);
    dq.delRight(MyNewDequeue, &dq.capacity);
    dq.delLeft(MyNewDequeue, &dq.capacity);
    dq.addLeft(454, MyNewDequeue, &dq.capacity);
    for (int i = 0; i < dq.capacity; i++) {
        printf("DEQ[%d] = %d\n", i, MyNewDequeue[i]);
    }
    printf("\ndeq cap: %d\n", dq.capacity);
    printf("\n");


*/

    free(MyNewDequeue);
    return 0;
}