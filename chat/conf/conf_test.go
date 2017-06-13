/* 
 * File Name: conf_test.go
 * Descript: 
 * 
 * Version: 1.0 
 * Create Time: 06/12/17 22:05:06
 * Compiler: 
 * Editor: vim 
 * Author: Jimbo
 * Mail: jimboo.lu@gmail.com
 * 
 * Company: 
 */ 
 
package conf

import (
	"fmt"
	"testing"
)

func Test_Debug(t *testing.T) {
	fmt.Println(Debug())
}

func Test_IpAndPort(t *testing.T) {
	fmt.Println(IpAndPort())
}
