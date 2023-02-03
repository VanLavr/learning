TARGET = MyDeq

$(TARGET) : main.o My_deq.o
	gcc main.o My_deq.o -o $(TARGET)

main.o : main.c
	gcc -c main.c -o main.o

My_deq.o : My_deq.c
	gcc -c My_deq.c -o My_deq.o

clean :
	del *.o *.exe