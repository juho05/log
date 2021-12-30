package main

import (
	"github.com/Bananenpro/log"
)

func main() {
	log.SetSeverity(log.TRACE)

	log.Trace("Hello", "World")
	log.Info("Dies ist eine tolle", "Info", 3)
	log.Warnf("Dies ist %d tolle %s", 1, "Warnung")
	log.Errorf("%s %s", "Hello", "World")
	log.Fatal("Bye", "World")
}
