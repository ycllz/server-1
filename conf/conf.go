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
	"github.com/larspensjo/config"
)

var debug bool

func init() {
	conf, err := config.ReadDefault("config.cfg")
	if err != nil {
		panic(err)
	}
	debug, err = conf.Bool("DEFAULT", "debug")
	if err != nil {
		panic(err)
	}
}

func DEBUG() bool {
	return debug
}
