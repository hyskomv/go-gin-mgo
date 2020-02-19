package config

import (
	"bytes"
	"log"
)

type Log struct {
	*log.Logger
}

func init()  {
	logger := log.New(&bytes.Buffer{}, "logger: ", log.Lshortfile)
	logger.Print("hello logger")
}
