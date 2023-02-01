/*

void SomePrint(); // void and empty brackets
void (*PrintPointer)(); // void and empty brackets
PrintPointer = Someprint;
// returning value type must match this type -> "void (*PrintPointer)()" and accepted arguments in brackets must match too...

*/



#include <stdio.h>
// headers of functions firstly
void SomePrint();

// than pointers on our functions
typedef void (*PrintPointer)();

// than structs, that uses functions
typedef struct
{
    int a;
    PrintPointer pp;
} SomeStruct;

// than creating an objects of our structs
static SomeStruct s1;

// using stuff
int main(int argc, char const *argv[]) {

    s1.pp = SomePrint;
    s1.pp();

    return 0;
}

// body of our functions
void SomePrint() {
    printf("Hi!\n");
}