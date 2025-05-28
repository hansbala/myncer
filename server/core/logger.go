package core

import "log"

func Printf(format string, v ...any) {
	log.Printf(format, v...)
}

func Errorf(v ...any) {
	log.SetPrefix("ERROR: ")
	log.Print(v...)
	log.SetPrefix("")
}

func Warningf(v ...any) {
	log.SetPrefix("ERROR: ")
	log.Print(v...)
	log.SetPrefix("")
}

func Fatalf(v ...any) {
	log.Fatal(v...)
}
