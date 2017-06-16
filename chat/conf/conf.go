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

const section = "DEFAULT"

var debug bool

var ipAndPort string

var timeOutSec int

var cleanSpan int

func init() {
	conf, err := config.ReadDefault("conf.cfg")
	if err != nil {
		panic(err)
	}

	debug, err = conf.Bool(section, "debug")
	if err != nil {
		panic(err)
	}

	ipAndPort, err = conf.String(section, "ipAndPort")
	if err != nil {
		panic(err)
	}

	timeOutSec, err = conf.Int(section, "timeOutSec")
	if err != nil {
		panic(err)
	}

	cleanSpan, err = conf.Int(section, "cleanSpan") 
	if err != nil {
		panic(err)
	}
}

func Debug() bool {
	return debug
}

func IpAndPort() string {
	return ipAndPort
}

func TimeOutSec() int {
	return timeOutSec
}

func CleanSpan() int {
	return cleanSpan
}
