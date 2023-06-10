package main

import (
	"fmt"

	sysinfo "github.com/elastic/go-sysinfo"
)

func main() {

	host, err := sysinfo.Host()

	if err != nil {
		panic(err)
	}

	basicInfo := host.Info()
	fmt.Println(basicInfo)

}
