package service

func (s ServicesFuncs) CreateWorkers(workerCount int, queueName string) {
	for range workerCount {
		s.IncrementActiveWorkerCount()
		go s.CallUpdateWorker(s.ctx, queueName)
	}
}
