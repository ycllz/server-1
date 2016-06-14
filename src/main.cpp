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

using namespace std;

int main(void) {

	// testing short length
	printf("testing short length\n");
	short shortNum = 1;
	int length = 1;
	while(shortNum > 0) {
		shortNum = shortNum << 1;
		length++;
	}
	printf("length=%d\n", length);

	// testing int length
	int intNum = 1;
	length = 1;
	printf("testing int length\n");
	while(intNum > 0) {
		intNum = intNum << 1;
		length++;
	}
	printf("length=%d\n", length);

	// testing long length
	long longNum = 1;
	length = 1;
	printf("testing long length\n");
	while(longNum > 0) {
		longNum = longNum << 1;
		length++;
	}
	printf("length=%d\n", length);

	// testing long long length
	long long longlongNum = 1;
	length = 1;
	printf("testing long long length\n");
	while(longlongNum > 0) {
		longlongNum = longlongNum << 1;
		length++;
	}
	printf("length=%d\n", length);
	return 0;
}
