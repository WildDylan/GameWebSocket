package scheduler

import (
	"runtime"
	"sync"
	"time"
)

type Scheduler struct {
	// was done
	done             chan bool
	schedulerTimer   *time.Ticker
	// timer map
	timers           map[int]*Timer
	// id
	autoId           int
	// lock
	listLocker       sync.Mutex
}

var GlobalScheduler *Scheduler = nil

func InitScheduler(tickTok time.Duration) {
	// MAX PROCESS = CPU NUM
	runtime.GOMAXPROCS(runtime.NumCPU())

	if tickTok <= (time.Millisecond * 100) {
		// SAFE FOR CPU USAGE
		tickTok = time.Duration(time.Millisecond * 100)
	}

	GlobalScheduler = &Scheduler{
		done:               make(chan bool),
		timers:			    make(map[int]*Timer),
		autoId:             1000,
		schedulerTimer:     time.NewTicker(tickTok),
	}

	// START PROCESS
	go GlobalScheduler.processScheduler()
}

// stop timer with id
func (c *Scheduler) StopTimer(Id int) {
	c.listLocker.Lock()
	defer c.listLocker.Unlock()

	timer, has := c.timers[Id]
	if has {
		timer.Stop()
	}
}

func (c *Scheduler) processScheduler() {
	defer func() {
		close(c.done)

		c.schedulerTimer.Stop()
	}()

	for {
		select {
		case <-c.schedulerTimer.C:
			c.listLocker.Lock()
			for Id := range c.timers {
				c.timers[Id].Execute()
			}
			c.listLocker.Unlock()
			break
		case <-c.done:
			return
		}
	}
}
