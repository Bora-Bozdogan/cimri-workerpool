package service

func (s ServicesFuncs) BlockWorkers() {
	<-s.ctx.Done()
}