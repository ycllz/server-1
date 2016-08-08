/*******************************************************************************
 File Name: main.cpp
 Descript:
 
 Version: 1.0
 Created Time: 
 Compiler: 
 Editor: vim
 Author: Jimbo
 mail: jimboo.lu@gmail.com
 
 Company: 
*******************************************************************************/

#include <cstdio>
#include <string>

using namespace std;

void display(const int* numPtr);
int main(void) {
<<<<<<< HEAD
	int num = 10;
	display(&num);
	printf("%d\n", num);
=======

	// testing compile environment
	printf("short length is %d\n", sizeof(short));
	printf("int length is %d\n", sizeof(int));
	printf("long length is %d\n", sizeof(long));
	printf("long long is %d\n", sizeof(long long));

>>>>>>> b39acc11f6da0c4e7b2e9cc499a8e4d38f2113fd
	return 0;
}

void display(const int* numPtr)
{
	*numPtr = 11;
}
