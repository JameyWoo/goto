package main

import "log"

func init() {
	log.SetPrefix("Log: ")
}

func main() {
	log.Println("this is a log")
}