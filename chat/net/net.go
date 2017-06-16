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
 
package net

import (
	"bytes"
	"chat/conf"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// tcp timeout second
var timeOutSec time.Duration

var handleFunc = make(map[int]func([]byte))

func init() {
	timeOutSec = time.Second * time.Duration(conf.TimeOutSec())

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

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(timeOutSec))

	for {
		msg := make([]byte, 1024)

		end, err := conn.Read(msg)
		if err != nil {
			fmt.Println(err)
			if e, k := err.(net.Error); k && e.Timeout() {
				break
			} else {
				continue
			}
		}
		go process(0, end, msg[0:end])
	}
}

func process(start, end int, msg []byte) {
	for ;start < end; {
		b_head := bytes.NewBuffer(msg[start:start+3])

		head := 0

		binary.Read(b_head, binary.BigEndian, &head)

		b_protoId := bytes.NewBuffer(msg[start+4:start+7])

		protoId := 0

		binary.Read(b_protoId, binary.BigEndian, &protoId)

		if v, ok :=handleFunc[protoId]; ok {
			v(msg[start+8:head])
		}

		start = head + 1
	}
}
