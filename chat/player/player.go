/* 
 * File Name: player.go
 * Descript: 
 * 
 * Version: 1.0 
 * Create Time: 06/16/17 10:41:40
 * Compiler: 
 * Editor: vim 
 * Author: Jimbo
 * Mail: jimboo.lu@gmail.com
 * 
 * Company: 
 */ 
 
package player

import (
	"chat/conf"
	"net"
	"sync"
	"time"
)

var playerMap = make(map[int]Player)

var mlock = new(sync.RWMutex)

type Player struct {
	id int

	conn net.Conn

	activeTime time.Time
}

func (p *Player)GetPlayerId() int {
	return p.id
}

func (p *Player)GetActiveTime() time.Time {
	return p.activeTime
}

func init() {
	go clean()
}

func clean() {
	for {
		time.Sleep(time.Minute * time.Duration(conf.CleanSpan()))

		for id, v := range playerMap {
			if (time.Now().Sub(v.activeTime) > time.Minute * time.Duration(conf.CleanSpan())) {
				delete(playerMap, id)
			}
		}
	}
}

func AddPlayer(playerId int, c net.Conn) Player {
	mlock.RLock()

	if _, ok := playerMap[playerId]; !ok {
		mlock.RUnlock()

		p := Player {
			id: playerId,
			conn: c,
			activeTime: time.Now(),
		}

		mlock.Lock()

		playerMap[p.GetPlayerId()] = p

		mlock.Unlock()

		return p
	}

	mlock.RUnlock()

	return playerMap[playerId]
}

func GetPlayer(playerId int) *Player {
	mlock.RLock()

	if v, ok := playerMap[playerId]; ok {
		return &v
	}

	mlock.RUnlock()

	return nil
}
