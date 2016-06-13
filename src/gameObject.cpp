/*******************************************************************************
 File Name: gameObject.cpp
 Descript:
 
 Version: 1.0
 Created Time: 
 Compiler: 
 Editor: vim
 Author: Jimbo
 mail: jimboo.lu@gmail.com
 
 Company: 
*******************************************************************************/

#ifndef GAMEOBJECT_H
#define GAMEOBJECT_H
#include "include/gameObject.h"
#include <cstdio>
#endif

#ifndef LIST_H
#define LIST_H
#include <list>
#endif

/// <summary>
/// gameObject constructor
/// </summary>
gameObject::gameObject() {
	this->commponentList = new list<commponent>();
	printf("gameObject constructor\n");
}

/// <summary>
/// gameObject destructor
/// </summary>
gameObject::~gameObject() {
	printf("gameObject destructor\n");
}

/// <summary>
/// add commponent to gameObject
/// </summary>
gameObject::addCommponent() {
	printf("addCommponent\n");
}

/// <summary>
/// get commponent from gameObject
/// </summary>
gameObject::getCommponent() {

}
