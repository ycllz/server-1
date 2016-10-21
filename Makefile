VPATH = ./src:./include
main:main.o threadpool.o
	g++ -pthread -o $@ threadpool.o main.o

main.o:main.cpp threadpool.h
	g++ -std=c++11 -c main.cpp

threadpool.o:threadpool.cpp threadpool.h
	g++ -std=c++11 -c threadpool.cpp

clean:
	rm main.o threadpool.o main
