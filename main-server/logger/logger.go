package logger

import (
	"log"
	"os"
)

const (
	ConsoleMode = iota
	FileMode    = iota
)

func Init(mode int, opts ...string) func() {
	log.SetFlags(log.Ldate | log.Ltime)

	if mode == FileMode {
		if len(opts) < 1 {
			log.Fatalf("wrong amount of arguments while initialising logger")
		}
		logfile, err := os.Create(opts[0])

		if err != nil {
			log.Fatalf("error initializing logger: %v", err)
		}

		log.SetOutput(logfile)

		return func() {
			err := logfile.Close()
			if err != nil {
				log.Fatalf("error while closing logfile: %v", err)
			}
		}
	}

	return func() {}
}
