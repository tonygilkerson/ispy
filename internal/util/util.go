package util

import "log"

func DoOrDie(err error) {
	if err != nil {
		log.Printf("This is not looking good! %v", err.Error())
		log.Panicf("Oops err: %v ", err)
	}
}