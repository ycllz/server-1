VPATH = src:include
objects = main.o threadpool.o
main:$(objects)
$(objects):%.o:%.cpp
	g++ -std=c++11 -c $< -o $@
clean:
	-rm *.o main
