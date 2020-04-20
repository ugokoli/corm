/*#####################################
# Ugonna Okoli                       #
# http://ugokoli.com                 #
# Copyright (c) 2019.                #
#####################################*/

package logger

import (
	"errors"
	"fmt"
	"log"
)

var flag = log.Ldate | log.Ltime | log.Lshortfile

func init() {
	log.SetFlags(flag)
}

func Info(s string, v ...interface{}) error {
	log.SetPrefix("CORM INFO: ")
	log.Printf(s, v...)

	return errors.New(fmt.Sprintf(s, v...))
}

func Debug(s string, v ...interface{}) error {
	log.SetPrefix("CORM DEBUG: ")
	log.Printf(s, v...)

	return errors.New(fmt.Sprintf(s, v...))
}

func Error(s string, v ...interface{}) error {
	log.SetPrefix("CORM ERROR: ")
	log.Printf(s, v...)

	return errors.New(fmt.Sprintf(s, v...))
}
