#include <stdio.h>
#include <stdlib.h>

int main(int argc, char const *argv[]) {
    int *arr = NULL;
    printf("addr -> %p\n", arr);
    arr = malloc(sizeof(int) * 10);
    for (int i = 0; i < 10; ++i) {
        arr[i] = i + 1;
        printf("arr[%d] = %d\n", i, arr[i]);
    }
    printf("addr -> %p\n", arr);
    printf("\n");

    arr = realloc(arr, sizeof(int) * 11);
    arr[10] = 11;
    arr = realloc(arr, sizeof(int) * 11);
    arr[11] = 12;
    arr = realloc(arr, sizeof(int) * 13);
    arr[12] = 13;
    arr = realloc(arr, sizeof(int) * 14);
    arr[13] = 14;
    arr = realloc(arr, sizeof(int) * 15);
    arr[14] = 15;
    arr = realloc(arr, sizeof(int) * 16);
    arr[15] = 16;
    for (int i = 0; i < 16; ++i) {
        arr[i] = i + 1;
        printf("arr[%d] = %d\n", i, arr[i]);
    }
    printf("\n");


    printf("MAGIC 1\n");
    int* tmp = NULL;
    tmp = malloc(sizeof(int) * 17);
    tmp[0] = -1;

    printf("tmp:\n");
    for (int i = 0; i < 16; i++) {
        tmp[i + 1] = arr[i];
        printf("tmp[%d] = %d ", i + 1, tmp[i + 1]);
    }
    printf("tmp[0] = %d", tmp[0]);

    arr = realloc(arr, sizeof(int) * 17);
    for (int i = 0; i < 17; i++) {
        arr[i] = tmp[i];
    }

    printf("\n");
    for (int i = 0; i < 17; ++i) {
        printf("arr[%d] = %d\n", i, arr[i]);
    }



    printf("\n\nMAGIC 2\n");
    tmp = NULL;
    tmp = malloc(sizeof(int) * 18);
    tmp[0] = -2;

    printf("tmp:\n");
    for (int i = 0; i < 17; i++) {
        tmp[i + 1] = arr[i];
        printf("tmp[%d] = %d ", i + 1, tmp[i + 1]);
    }
    printf("tmp[0] = %d", tmp[0]);

    arr = realloc(arr, sizeof(int) * 18);
    for (int i = 0; i < 18; i++) {
        arr[i] = tmp[i];
    }

    printf("\n");
    for (int i = 0; i < 18; ++i) {
        printf("arr[%d] = %d\n", i, arr[i]);
    }



    printf("\n\nMAGIC 3\n");
    tmp = NULL;
    tmp = malloc(sizeof(int) * 19);
    tmp[0] = -3;

    printf("tmp:\n");
    for (int i = 0; i < 18; i++) {
        tmp[i + 1] = arr[i];
        printf("tmp[%d] = %d ", i + 1, tmp[i + 1]);
    }
    printf("tmp[0] = %d", tmp[0]);

    arr = realloc(arr, sizeof(int) * 19);
    for (int i = 0; i < 19; i++) {
        arr[i] = tmp[i];
    }

    printf("\n");
    for (int i = 0; i < 19; ++i) {
        printf("arr[%d] = %d\n", i, arr[i]);
    }

    // !!!!!!
    free(arr);
    free(tmp);
    
    return 0;
}