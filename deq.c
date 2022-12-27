#include <stdio.h>
#include <stdlib.h>

void addToLeft(int element, int* DEQ, int* size_of_deq);
void addToRight(int element, int *DEQ, int *size_of_deq);
void delete_from_left(int *DEQ, int *size_of_deq);
void delete_from_right(int *DEQ, int *size_of_deq);
int* create_deq(int size);

int main(int argc, char const *argv[]) {

    int size_of_deq = 10;
    int *deq = malloc(size_of_deq * sizeof(int));

    for (int i = 0; i < size_of_deq; i++) {
        *(deq + i) = i;
    }
    
    for (int i = 0; i < size_of_deq; i++) {
        printf("%d ", *(deq + i));
    }

    printf("\n\n");

    delete_from_right(deq, &size_of_deq);
    delete_from_left(deq, &size_of_deq);
    delete_from_right(deq, &size_of_deq);
    delete_from_left(deq, &size_of_deq);
    addToRight(169, deq, &size_of_deq);
    addToLeft(16, deq, &size_of_deq);
    for (int i = 0; i < size_of_deq; i++) {
        printf("%d ", *(deq + i));
    }

    free(deq);

    return 0;
}

void delete_from_right(int *DEQ, int *size_of_deq) {
    DEQ = realloc(DEQ, sizeof(int) * (*size_of_deq - 1) );
    *size_of_deq = *size_of_deq - 1;
}

void delete_from_left(int *DEQ, int *size_of_deq) {
    int* tmp = (int *)malloc(sizeof(int) * (*size_of_deq - 1));
    for (int i = 1; i < *size_of_deq; i++) {
        tmp[i - 1] = DEQ[i];
    }

    DEQ = realloc(DEQ, sizeof(int) * (*size_of_deq - 1) );
    *size_of_deq = *size_of_deq - 1;

    for (int i = 0; i < *size_of_deq; i++) {
        DEQ[i] = tmp[i];
    }

    free(tmp);
}

void addToRight(int element, int *DEQ, int *size_of_deq) {
    DEQ = realloc(DEQ, sizeof(int) * (*size_of_deq + 1) );
    *size_of_deq = *size_of_deq + 1;
    DEQ[*size_of_deq - 1] = element;
}

void addToLeft(int element, int* DEQ, int* size_of_deq) {
    int* tmp = malloc(sizeof(int) * (*size_of_deq + 1));
    
    for (int i = 0; i < *size_of_deq; i++) {
        tmp[i + 1] = DEQ[i];
    }

    tmp[0] = element;
    
    DEQ = realloc(DEQ, sizeof(int) * (*size_of_deq + 1));
    *size_of_deq = *size_of_deq + 1;
    
    for (int i = 0; i < *size_of_deq; i++) {
        DEQ[i] = tmp[i];
    }

    free(tmp);
}


// до удаления справа:
// [0 0 0 0 0 0 0 0 0 0]
// после удаления справа:
// [0 0 0 0 0 0 0 0 0] 0


// до удаления слева:
// [0 0 0 0 0 0 0 0 0 0]
// после удаления слева:
// 0 [0 0 0 0 0 0 0 0 0]


// *(deq)->  0
//    deq-> [0]
//          [0]
//          [0]
//          [0]
//          [0]
//          [0]