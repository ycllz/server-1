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
	int num = 10;
	display(&num);
	printf("%d\n", num);
	return 0;
}

void display(const int* numPtr)
{
	*numPtr = 11;
}
