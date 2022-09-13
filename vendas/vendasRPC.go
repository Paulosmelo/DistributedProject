package vendas

import (
	"sync"
)

type VendasRPC struct{}

var (
	wgRPC = sync.WaitGroup{}// guards
	arrayRPC  = [20]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
)


func (t *VendasRPC) GetVendasRPC(n int, reply *[20]int) error {
	wgRPC.Add(1)
	go AddVendaRPC(n)
	wgRPC.Wait()
	*reply =  arrayRPC
	return nil
}

func AddVendaRPC (n int) {
	defer wgRPC.Done()
	arrayRPC[n] = arrayRPC[n] + 1
}