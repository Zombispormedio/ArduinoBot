package arduino

import (

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"server_request"
	"strconv"
	"strings"
	"time"
	"runtime"
)

type fn func(int)

var (
	so						=map[string]string{ "windows":"COM3", "linux": "/dev/ttyACM0",}
	gbot           = gobot.NewGobot()
	firmataAdaptor = firmata.NewFirmataAdaptor("arduino", so[runtime.GOOS])
	led4           = gpio.NewLedDriver(firmataAdaptor, "led", "4")
	led5           = gpio.NewLedDriver(firmataAdaptor, "led", "5")
	led6           = gpio.NewLedDriver(firmataAdaptor, "led", "6")
	led7           = gpio.NewLedDriver(firmataAdaptor, "led", "7")
	led8           = gpio.NewLedDriver(firmataAdaptor, "led", "8")
	led9           = gpio.NewLedDriver(firmataAdaptor, "led", "9")
	led10          = gpio.NewLedDriver(firmataAdaptor, "led", "10")
	led11          = gpio.NewLedDriver(firmataAdaptor, "led", "11")

	leds = []*gpio.LedDriver{led4, led5, led6, led7, led8, led9, led10, led11}
)

func on(i int) {
	leds[i].On()
}

func off(i int) {
	leds[i].Off()
}

func toggle(i int) {
	leds[i].Toggle()
}

func switchSome(cb fn, ls []*gpio.LedDriver) {
	for i, _ := range ls {
		cb(i)
	}

}

func switchSomeIndex(cb fn, ls []string) {
	for _, a := range ls {
		if v, err := strconv.Atoi(a); err == nil {
			if v >= 0 && v < 8 {
				cb(v)
			}else{

				if v>=8{
					cb(v%8)
				}
			}
		}

	}

}
func switchAll(cb fn) {
	switchSome(cb, leds)

}

func twoParameters(parts []string) {
	switch a := parts[1]; a {
	case "on":
		switchAll(on)
	case "off":
		switchAll(off)

	default:
		if v, err := strconv.Atoi(a); err == nil {
			if v >= 0 {
				switchSomeIndex(toggle, parts[1:])

			}
		}

	}

}

func moreParameters(parts []string) {
	switch a := parts[1]; a {
	case "on":
		switchSomeIndex(on, parts[2:])
	case "off":
		switchSomeIndex(off, parts[2:])

	default:
		if v, err := strconv.Atoi(a); err == nil {
			if v >= 0{
				switchSomeIndex(toggle, parts[1:])

			}
		}

	}
}

func choose(parts []string) {
	length := len(parts)
	switch {
	case length == 1:
		switchAll(toggle)

	case length == 2:
		twoParameters(parts)

	case length > 2:
		moreParameters(parts)
	}
}

func Action() {

	work := func() {
		//server_request.Advice()
		gobot.Every(1*time.Second, func() {
			res := server_request.Check()
switchAll(off)
			if res.Data != nil {
				message := res.Data[0]
				
				choose(strings.Split(message, " "))
			}

		})

	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led4, led5, led6, led7, led8, led9, led10, led11},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
