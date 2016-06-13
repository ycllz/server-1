/*******************************************************************************
 File Name: gameObject.h
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
#include "object.h"
#endif

#ifndef LIST_H
#define LIST_H
#include <list>
#endif

#ifndef COMMPONENT_H
#define COMMPONENT_H
#include "commponent.h"
#endif

/// <summary>
/// gameObject
/// </summary>
class gameObject : public object {
	public:
		gameObject();
		virtual ~gameObject();
		addCommponent();
		getCommponent();

	private: 
		list<commponent>* commponentList;
};
