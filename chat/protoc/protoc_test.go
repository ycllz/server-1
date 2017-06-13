/* 
 * File Name: protoc_test.go
 * Descript: 
 * 
 * Version: 1.0 
 * Create Time: 06/14/17 00:38:15
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
	"testing"
	"time"
)

func Test_Protoc(t *testing.T) {
	time.Sleep(time.Second)

	conn, err := net.Dial("tcp", conf.IpAndPort())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("listen...")

	conn.Write([]byte{'h', 'e', 'l', 'l', 'o'})

	data := make([]byte, 1024)

	index, err := conn.Read(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data[0:index]))
}
