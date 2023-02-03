#include <stdio.h>
#include <stdlib.h>

int main(int argc, char const *argv[]) {
    int *arr;
    printf("addr -> %p\n", arr);
    arr = malloc(sizeof(int) * 10);
    for (int i = 0; i < 10; ++i) {
        arr[i] = i + 1;
        printf("arr[%d] = %d\n", i, arr[i]);
    }
    printf("addr -> %p\n", arr);

    // !!!!!!
    free(arr);
    
    return 0;
}