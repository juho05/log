package main

import (
	"os"

	"github.com/Bananenpro/log"
)

func main() {
	log.SetSeverity(log.INFO)

	log.Trace("Hello", "World")
	log.Info("Dies ist eine tolle", "Info", 3)
	log.Warnf("Dies ist %d tolle %s", 1, "Warnung")
	log.Errorf("%s %s", "Hello", "World")

	log.SetOutput(os.Stdout)
	log.Trace("Hello", "World")
	log.Info("Dies ist eine tolle", "Info", 3)
	log.Warnf("Dies ist %d tolle %s", 1, "Warnung")
	log.Errorf("%s %s", "Hello", "World")

	f, _ := os.Create("test.log")
	log.SetOutput(f)
	log.Trace("Hello", "World")
	log.Info("Dies ist eine tolle", "Info", 3)
	log.Warnf("Dies ist %d tolle %s", 1, "Warnung")
	log.Errorf("%s %s", "Hello", "World")
	log.Fatal("Bye", "World")
}
