package service

func (s ServicesFuncs) CreateWorkers(workerCount int, queueName string) {
	for range workerCount {
		go s.CallUpdateWorker(s.ctx, queueName)
	}
}
