#include <stdio.h>

void Greeting();
void AddToRight(int element);
void AddToLeft(int element);
int  DeleteFromRight();
int  DeleteFromLeft();

typedef void (*ptr_Greeting)();
typedef void (*ptr_AddToRight)(int);
typedef void (*ptr_AddToLeft)(int);
typedef int  (*ptr_DeleteFromRight)();
typedef int  (*ptr_DeleteFromLeft)();

typedef struct {
    
    int  capacity;
    int* deq;

    ptr_AddToRight      addRight;
    ptr_AddToLeft       addLeft;
    ptr_DeleteFromRight delRight;
    ptr_DeleteFromLeft  delLeft;
    ptr_Greeting        greet;

} dequeue;

void Greeting() {
    printf("Hi!\n");
}