VPATH = src:include
objects = main.o threadpool.o
main:$(objects)
$(objects):%.o:%.cpp
	g++ -std=c++11 -c $< -o $@
install:
	mkdir bin
	mv main bin/
clean:
	-rm *.o main
