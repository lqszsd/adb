package service

import (
	"fmt"
	hook "github.com/robotn/gohook"
)

func direction(){
	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		fmt.Println("anjiana",ev.Keycode)
		getAdbCommand(ev.Keycode)
	}
}
func New(){
	direction()
}
func getAdbCommand(keycode uint16){
	switch keycode {
	case 31:
		AdbShellInputSwipe(302,609,618,667)
		break
	case 17:
		AdbShellInputSwipe(618,667,302,609)
	}
}
