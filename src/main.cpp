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

#ifndef OBJECT_H
#define OBJECT_H
#include "include/object.h"
#endif

#ifndef GAMEOBJECT_H
#define GAMEOBJECT_H
#include "include/gameObject.h"
#endif

int main(void) {
	object* objPtr = new gameObject();
	delete objPtr;
	return 0;
}
