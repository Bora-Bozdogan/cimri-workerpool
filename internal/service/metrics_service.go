package service

func (s ServicesFuncs) IncrementActiveWorkerCount() {
	s.metrics.IncrementActiveWorkerCount()
}
func (s ServicesFuncs) DecrementActiveWorkerCount() {
	s.metrics.DecrementActiveWorkerCount()
}
