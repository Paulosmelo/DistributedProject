package vendas

import (
	"sync"
)

type VendasTCP struct{}

var (
	wgTCP = sync.WaitGroup{}// guards
	arrayTCP  = [20]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
)


func (VendasTCP) GetVendasTCP(n int) [20]int {
	wgTCP.Add(1)
	go AddVendaTCP(n)
	wgTCP.Wait()
	return arrayTCP	
}

func AddVendaTCP(n int) {
	defer wgTCP.Done()
	arrayTCP[n] = arrayTCP[n] + 1
}