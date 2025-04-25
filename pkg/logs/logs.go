package logs

import (
	"log"
	"os"
)

type Logwriter struct {
	Info  *log.Logger
	Error *log.Logger
}

func New() *Logwriter {
	return &Logwriter{
		Info:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		Error: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
