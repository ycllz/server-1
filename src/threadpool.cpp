/* 
 * File Name: threadpool.cpp
 * Descript: 
 * 
 * Version: 1.0 
 * Created Time: 10/13/16 22:21:53
 * Compiler: 
 * Editor: vim 
 * Author: Jimbo 
 * Mail: jimboo.lu@gmail.com 
 *
 * Company: 
 */
 
#include "../include/threadpool.h"

threadpool::threadpool() {

}

threadpool::threadpool(int max_thread) {
	this->m_max_thread = max_thread;
}

threadpool::~threadpool() {

}
