package main

import "os"

func main() {
	switch os.Args[1] {
	case "load-modules":
		loadModules()
	case "huge-pages":
		setupHugePages()
	case "setup-nat":
		setupNat()
	case "run-qemu":
		runQemu()
	}
}
