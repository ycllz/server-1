/*******************************************************************************
 File Name: object.cpp
 Descript:
 
 Version: 1.0
 Created Time: 
 Compiler: 
 Editor: vim
 Author: Jimbo
 mail: jimboo.lu@gmail.com
 
 Company: 
*******************************************************************************/

#ifndef OBJECT_CPP
#define OBJECT_CPP
#include "include/object.h"
#include <cstdio>
#endif

/// <summary>
/// constructor
/// </summary>
object::object() {
	printf("constructor\n");
}

/// <summary>
/// destructor
/// </summary>
object::~object() {
	printf("destructor\n");
}
