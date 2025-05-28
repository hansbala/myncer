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

func Warningf(format string, v ...any) {
	log.SetPrefix("WARNING: ")
	log.Printf(format, v...)
	log.SetPrefix("")
}

func Warning(v ...any) {
	log.SetPrefix("WARNING: ")
	log.Print(v...)
	log.SetPrefix("")
}

func Fatalf(v ...any) {
	log.Fatal(v...)
}
