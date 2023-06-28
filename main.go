package main

import (
	"flag"
	"fmt"
	"github.com/mPandaer/go-utils/utils"
)

func main() {
	suffix := flag.String("rs", "", "--remove suffix")
	dir := flag.String("d", "", "--source dir")
	flag.Parse()

	if *suffix == "" || *dir == "" {
		fmt.Println("suffix or dir is empty")
		return
	}

	err := utils.RemoveSuffixFromFile(*dir, *suffix)
	fmt.Println(*suffix, *dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Done")

}
