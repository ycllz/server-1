VPATH = src:include
objects = main.o threadpool.o
main:$(objects)
$(objects):%.o:%.cpp
	g++ -std=c++qq -c $< -o $@
clean:
	-rm *.o main
