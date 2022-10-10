package main

import (
	"coursecontent/pkg/msg"
	"coursecontent/pkg/display"
)

func main() {
	msg.Hi()
	display.Display("Hello from display")
	msg.Exciting("an exciting message")
}