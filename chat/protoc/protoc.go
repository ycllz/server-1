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
	"bytes"
	_ "chat/chat"
	"chat/conf"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"reflect"
	"time"
)

var timeOutSec time.Duration

var mapFunc map[int32]func([]byte)

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
		go process(msg[0:end])
	}
}

func process(msg []byte) {
	start := 0

	len := len(msg)

	for start < len {
		b_head := bytes.NewBuffer(msg[start:start+3]

		var head int32

		binary.Read(b_head, binary.BigEndian, &head)

		b_protoId := bytes.NewBuffer(msg[start+4:start+7])

		var protoId int32

		binary.Read(b_protoId, binary.BigEndian, &protoId)

		if v, ok := mapFunc[protoId]; ok {
			v(msg[start+8:head])
		}

		start = head + 1
	}
}
