#include <stdlib.h>

void AddToRight(int element, int* deq, int* capacity);
void AddToLeft(int element, int* deq, int* capacity);
int  DeleteFromRight(int* deq, int* capacity);
int  DeleteFromLeft(int* deq, int* capacity);
int* CreateDequeue(int capacity);

typedef void (*ptr_AddToRight)(int, int*, int*);
typedef void (*ptr_AddToLeft)(int, int*, int*);
typedef int  (*ptr_DeleteFromRight)(int*, int*);
typedef int  (*ptr_DeleteFromLeft)(int*, int*);
typedef int* (*ptr_CreateDequeue)(int);

typedef struct {
    
    int  capacity;
    //int* deq;

    ptr_AddToRight      addRight;
    ptr_AddToLeft       addLeft;
    ptr_DeleteFromRight delRight;
    ptr_DeleteFromLeft  delLeft;
    ptr_CreateDequeue   create;

} dequeue;

int* CreateDequeue(int capacity) {
    
    int* deq;
    deq = malloc(sizeof(int) * capacity);
    
    return deq;
}

void AddToRight(int element, int* deq, int* capacity) {
    
    deq = realloc( deq, sizeof(int) * (*capacity + 1) );
    *capacity = *capacity + 1;
    deq[*capacity - 1] = element;
    
}

void AddToLeft(int element, int* deq, int* capacity) {

    int* tmp = malloc(sizeof(int) * (*capacity + 1));
    for (int i = 0; i < *capacity; ++i) {
        tmp[i + i] = deq[i];
    }
    tmp[0] = element;

    deq = realloc( deq, sizeof(int) * (*capacity + 1) );
    *capacity = *capacity + 1;
    for (int i = 0; i < *capacity; ++i) {
        deq[i] = tmp[i];
    }

    free(tmp);

}