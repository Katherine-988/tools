package tools

import "log"

func Errorln(v ...any) {
	log.Println("[err]", v)
}

func Infoln(v ...any) {
	log.Println("[info]", v)
}

func Debugln(v ...any) {
	log.Println("[debug]", v)
}
