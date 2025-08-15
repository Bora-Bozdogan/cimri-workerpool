package handlers

type servicesInterface interface {
	CreateWorkers(workerCount int, queueName string)
	BlockWorkers()
}

type Handler struct {
	service servicesInterface
}

func NewHandler(service servicesInterface) Handler {
	return Handler{service: service}
}

func (h Handler) HandleWorkers() {
	//create workers for each queue
	h.service.CreateWorkers(20, "high")
	h.service.CreateWorkers(10, "med")
	h.service.CreateWorkers(5, "low")

	h.service.BlockWorkers()
}
