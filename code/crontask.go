package main

import (
	"sync"
	"time"
)

var Crontask *crontask

type crontask struct {
	timers map[string]*time.Timer
	mutex  sync.Mutex
}

func NewCrontask() {
	Crontask = &crontask{
		timers: make(map[string]*time.Timer),
	}

}
