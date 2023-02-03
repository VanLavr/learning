#ifndef MY_DEQ_H
#define MY_DEQ_H

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

#endif