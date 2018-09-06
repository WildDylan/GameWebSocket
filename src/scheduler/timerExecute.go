package scheduler

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"sync"
	"sync/atomic"
)

type Job interface{}

type Function struct {
	exeCall         Job
	args            []reflect.Value
	status          int32
	mux             sync.Mutex
}

type Timer struct {
	Id              int
	Finished        int32
	cbFunc          *Function
}

func reflectFunc(fn Job, args ...interface{}) (cb *Function, err error) {
	cb = nil
	err = nil

	t := reflect.TypeOf(fn)

	if t.Kind() != reflect.Func {
		err = errors.New("callback must be a function")
		return
	}

	// Make Params to Array
	argArr := []interface{}(args)

	if len(argArr) < t.NumIn() {
		err = errors.New("not enough arguments")
		return
	}

	var in = make([]reflect.Value, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		argType := t.In(i)
		if argType != reflect.TypeOf(argArr[i]) {
			err = errors.New(fmt.Sprintf("Value not found for type %v", argType))
			return
		}
		in[i] = reflect.ValueOf(argArr[i])
	}

	cb = &Function{
		exeCall:        fn,
		args:           in,
		status:         0,
	}
	return
}

func callFunc(timer *Function) {
	// Lock
	timer.mux.Lock()
	defer timer.mux.Unlock()

	if timer.status >= 1 {
		return
	}

	// Count ++
	timer.status++
	// Call function
	reflect.ValueOf(timer.exeCall).Call(timer.args)
	timer.status--
}

func newTimer(Id int, fn Job, args ...interface{}) *Timer {

	cbo, err := reflectFunc(fn, args...)
	if err != nil {
		log.Printf("error:%v", err)
		return nil
	}

	newValue := &Timer{
		Id:             Id,
		Finished:       0,
		cbFunc:         cbo,
	}

	return newValue
}

func (c *Timer) Execute() {
	if c.IsDestroy() {
		delete(GlobalScheduler.timers, c.Id)
		return
	}

	go callFunc(c.cbFunc)
}

/// Can be destroy
func (c *Timer) IsDestroy() bool {
	return atomic.LoadInt32(&c.Finished) == 1
}

func (c *Timer) Stop() {
	atomic.StoreInt32(&c.Finished, 1)
}