package vendas

import (
	"strconv"
	"fmt"
)

type Vendas struct{}


func (Vendas) GetVendas (n string) string {
	i, err := strconv.Atoi(n)
	if err != nil {fmt.Printf("Error converting message.\n")}

	var r = array[i]

	return strconv.Itoa(r)
}