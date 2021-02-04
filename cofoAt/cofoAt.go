package main

import (
	"cf/getData"
	"fmt"
)

func main() {
	getData.GetContest()
	fmt.Println(getData.GetTier("djs100201"))
}
