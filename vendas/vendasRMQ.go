package vendas

import (
	"sync"
)

type VendasRMQ struct{}

var (
	wgRMQ = sync.WaitGroup{}// guards
	arrayRMQ  = [20]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
)


func (t *VendasRMQ) GetVendasRMQ(n int, reply *[20]int) error{
	wgRMQ.Add(1)
	go AddVendaRMQ(n)
	wgRMQ.Wait()
	*reply =  arrayRMQ
	return nil
}

func AddVendaRMQ (n int) {
	defer wgRMQ.Done()
	arrayRMQ[n] = arrayRMQ[n] + 1
}