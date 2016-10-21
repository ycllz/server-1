/*
 * File Name: threadpool.h
 * Descript: 
 * 
 * Version: 1.0 
 * Created Time: 10/11/16 22:32:22
 * Compiler: 
 * Editor: vim 
 * Author: Jimbo 
 * Mail: jimboo.lu@gmail.com 
 * 
 * Company: 
 */ 
 
#ifndef _THREADPOOL_H_
#define _THREADPOOL_H_

#include <thread>
#include <mutex>
#include <cstdio>

using namespace std;

class threadpool {
	public :
		threadpool();
		threadpool(int max_thread);
		~threadpool();
	public :
		int m_max_thread;
		thread * m_thread_head;
};

#endif
