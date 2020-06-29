package main

import (
	"androidphone/service"
	"fmt"
	"github.com/go-vgo/robotgo"
)

const (
	//可用下面的AdbShellDumpsysActivityF函数获取包名和activity名
	APPPackageName = "cn.XXX.android"
	APP            = "cn.XXX.android/com.XXX.XXXActivity"
)

func main() {
	//robotgo.ScrollMouse(10, "up")
	//robotgo.MouseClick("left", true)
	//robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
	service.New();
	//todo
	event()
	select{

	}
	////如果手机是休眠状态，则打开电源
	//if AdbShellDumpsysPowerOff() {
	//	AdbShellInputKeyEvent("26") //power
	//}
	////进入手机主屏
	//AdbShellInputKeyEvent("4") //back
	//AdbShellInputKeyEvent("3") //home
	///*如果APP未启动，则启动APP
	//  if !strings.Contains(AdbShellDumpsysActivityF(), APPPackageName) {
	//      AdbShellAmStartN(APP)
	//  }
	//*/
	//Tap("设置", 0)
	//TimeSleepDuration(5)
	//TapOnce(`\d我的`, 0, 3, 573)
	//AdbShellInputKeyEvent("26") //power

}

func event() {
	ok := robotgo.AddEvents("q", "ctrl", "shift")
	if ok {
		fmt.Println("add events...")
	}

	keve := robotgo.AddEvent("k")
	if keve {
		fmt.Println("you press... ", "k")
	}

	mleft := robotgo.AddEvent("mleft")
	if mleft {
		fmt.Println("you press... ", "mouse left button")
	}
}
