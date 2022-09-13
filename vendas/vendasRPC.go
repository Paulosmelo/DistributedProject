package vendas

import (
	"strconv"
	"fmt"
)

type VendasRPC struct{}

var arrayRPC  =  [20]int{1,2,231,3,1213,42,54,509,2101,201,391,30,102,1021,281,21,3,13,312,312}


func (t *VendasRPC) GetVendasRPC(n string, reply *string) error{
	i, err := strconv.Atoi(n)
	if err != nil {fmt.Printf("Error converting message.\n")}
	*reply =  strconv.Itoa(arrayRPC[i])
	return nil
}