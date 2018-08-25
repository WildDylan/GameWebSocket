package scheduler

func SubmitJob(fn Job, args ...interface{}) int {
	nowId := GlobalScheduler.autoId
	GlobalScheduler.autoId++
	newTimer := newTimer(
		nowId,
		fn,
		args...)

	GlobalScheduler.listLocker.Lock()
	GlobalScheduler.timers[nowId] = newTimer
	GlobalScheduler.listLocker.Unlock()

	return nowId
}

func RemoveJob(Id int) {
	GlobalScheduler.StopTimer(Id)
}