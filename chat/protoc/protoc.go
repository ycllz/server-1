/* 
 * File Name: protoc.go
 * Descript: 
 * 
 * Version: 1.0 
 * Create Time: 06/14/17 00:17:45
 * Compiler: 
 * Editor: vim 
 * Author: Jimbo
 * Mail: jimboo.lu@gmail.com
 * 
 * Company: 
 */ 
 
package protoc

import (
	"chat/conf"
	"fmt"
	"net"
	"time"
)

func init() {
	go listen()
}

func listen() {
	listen, err := net.Listen("tcp", conf.IpAndPort())
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println(conf.TimeOutSec())

		timeOutSec := time.Second * time.Duration(conf.TimeOutSec())

		conn.SetReadDeadline(time.Now().Add(timeOutSec))

		go read(conn)
	}
}

func read(conn net.Conn) {
	for {
		data := make([]byte, 1024)

		index, err := conn.Read(data)
		if err != nil {
			fmt.Println(err)
			if e, k := err.(net.Error); k && e.Timeout() {
				break
			} else {
				continue
			}
		}

		fmt.Println(string(data[0:index]))

		conn.Write([]byte{'f', 'i', 'n', 'i', 's', 'h'})
	}
}
