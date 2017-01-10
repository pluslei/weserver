package writedb

import (
	"fmt"
	"sync"
	"time"
)

var w *WriteData

type WriteData struct {
	lock     *sync.RWMutex
	jsondata chan *socketdata
}

type socketdata struct {
	Username string
	Content  string
}

func init() {
	w = &WriteData{
		lock:     &sync.RWMutex{},
		jsondata: make(chan *socketdata, 2048),
	}
	w.runWriteDb()
}

func addWriteData(v interface{}) {
	w.lock.Lock()
	w.jsondata <- v.(socketdata)
	w.lock.Unlock()
}

func (w *WriteData) runWriteDb() {
	go func() {
		for {
			msg, ok := <-w.jsondata
			if ok {
				addData(msg)
			}
		}

		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
	}()
}

func addData(msg *socketdata) {
	fmt.Println(msg)
}
