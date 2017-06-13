/* 
 * File Name: conf.go
 * Descript: 
 * 
 * Version: 1.0 
 * Create Time: 06/12/17 21:17:24
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
	"github.com/larspensjo/config"
)

var debug bool

func init() {
	conf, err := config.ReadDefault("config.cfg")
	if err != nil {
		fmt.Println(err)
	}
	debug, err = conf.Bool("DEFAULT", "debug")
	if err != nil {
		fmt.Println(err)
	}
}

func DEBUG() bool {
	return debug
}
