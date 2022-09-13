package vendas

import (
	"sync"
)

type VendasMOM struct{}

var (
	wg = sync.WaitGroup{}// guards
	array  = [20]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
)


func (t *VendasMOM) GetVendasMOM(n int, reply *[20]int) error{
	wg.Add(1)
	go AddVenda(n)
	wg.Wait()
	*reply =  array
	return nil
}

func AddVenda (n int) {
	defer wg.Done()
	array[n] = array[n] + 1
}