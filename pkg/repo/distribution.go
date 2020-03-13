package repo

import "fmt"

var (
	distributions = []string{
		"raspbian",
	}
)

func ListDistributions() {
	for _, distribution := range distributions {
		fmt.Println(distribution)
	}
}
