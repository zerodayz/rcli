// +build linux darwin

package main

import (
	"github.com/zerodayz/rcli/cmd"
)

func main() {
	cmd.Execute()
}
//
//func main() {
//	if vars.Debug == true {
//		log.Printf("DEBUG: starting container\n")
//	}
//	switch os.Args[1] {
//	case "run":
//		containers.Run()
//	case "child":
//		containers.Child()
//	default:
//		panic("wat should I do")
//	}
//}