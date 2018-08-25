package scheduler

import "go/types"

// Boost scheduler was a interface

type JobInterface interface {

	Start()
	End()
	Next()

}

const (
	CountDown = 0
	EveryOne_CountDown
)

var Id = 0
var Jobs = make(map[int]*BoostJob)

type BoostJob struct {

	Id              int
 	FireInterval    int
	FireFunction    types.Func
	Type            int
	State           int

}

func NewBoostJob(interval int, function types.Func, Type int, State int) *BoostJob {
	Id ++

	return &BoostJob{
		Id: Id,
		FireInterval:interval,
		FireFunction:function,
		Type:Type,
		State: State,
	}
}

func (job *BoostJob) Start()  {
	Jobs[job.Id] = job
}

func (job *BoostJob) End()  {
	delete(Jobs, job.Id)
}

func (job *BoostJob) Next()  {
	if job.State == 1 {
		// next
	} else {
		// next 
	}
}

func SubmitBoostJob(job JobInterface)  {
	go job.Start()
}