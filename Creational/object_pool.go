package patterns

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

//OBJECT POOL

type idbConnection interface {
	getId() int
}

type connection struct {
	id int
}

func (c connection) getId() int {
	return c.id
}

type connectionsPool struct {
	idle     []idbConnection
	active   []idbConnection
	capacity int
}

func (cp *connectionsPool) getConnection() idbConnection {
	mutex.Lock()
	if len(cp.idle) > 0 {
		cp.active = append(cp.active, cp.idle[0])
		dbcon := cp.idle[0]
		cp.idle = cp.idle[1:]
		println("got connection", dbcon.getId())
		mutex.Unlock()
		return dbcon
	} else {
		println("waiting for connection")
		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
		return cp.getConnection()
	}
}

func (cp *connectionsPool) releaseConnection() {
	mutex.Lock()
	if len(cp.active) == 0 {
		mutex.Unlock()
		println("no active connections")
		return
	}
	println("returning connection")
	cp.idle = append(cp.idle, cp.active[0])
	cp.active = cp.active[1:]
	fmt.Println("---", cp.idle)
	mutex.Unlock()
}

func ObjectPoll() {

	cp := connectionsPool{}
	cp.capacity = 5
	for i := 0; i < cp.capacity; i++ {
		con := connection{id: i}
		cp.idle = append(cp.idle, con)
	}

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			go cp.getConnection()
		}
	}()

	time.Sleep(1 * time.Second)
	go func() {
		for {
			time.Sleep(50 * time.Millisecond)
			go cp.releaseConnection()
		}
	}()
	time.Sleep(3 * time.Second)
}
