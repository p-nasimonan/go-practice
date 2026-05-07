package main

import (
	"fmt"
)

type Server struct {
	Name string
	Load int
}

func printLoads(servers []Server) {
	for _, s := range servers {
		if 50 <= s.Load && s.Load < 80 {
			fmt.Printf("[Warning]%s is %d\n", s.Name, s.Load)
		} else if 80 <= s.Load {
			fmt.Printf("[Critical]%s is %d\n", s.Name, s.Load)
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	loads := make([]Server, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&loads[i].Name, &loads[i].Load)
	}
	printLoads(loads)
}
