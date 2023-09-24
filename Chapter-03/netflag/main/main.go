package main

import (
	"fmt"

	"netflag"
)

func main() {
	var v netflag.Flags = netflag.FlagMulticast | netflag.FlagUp
	fmt.Printf("%b %t\n", v, netflag.IsUp(v))
	netflag.TurnDown(&v)
	fmt.Printf("%b %t\n", v, netflag.IsUp(v))
	netflag.SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, netflag.IsUp(v))
	fmt.Printf("%b %t\n", v, netflag.IsCast(v))
}
