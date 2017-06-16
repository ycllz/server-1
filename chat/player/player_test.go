/* 
 * File Name: player_test.go
 * Descript: 
 * 
 * Version: 1.0 
 * Create Time: 06/16/17 12:29:05
 * Compiler: 
 * Editor: vim 
 * Author: Jimbo
 * Mail: jimboo.lu@gmail.com
 * 
 * Company: 
 */ 
 
package player

import (
	"fmt"
	"testing"
)

func Test_Player(t *testing.T) {
	player := GetPlayer(0)
	
	if player == nil {
		p := AddPlayer(0, nil)
		fmt.Println(p)
	}
	fmt.Println(player)
}
