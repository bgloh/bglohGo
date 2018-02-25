package main

import (
        "time"
        "fmt"
        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/firmata"
)

var buttonCount int = 0

func pinToggle(pin *gpio.DirectPinDriver) (toggle func()){
  var i byte =0
  toggle = func(){
    i^=1
    pin.DigitalWrite(i)
  }
  return
}           

func main() {
        firmataAdaptor := firmata.NewTCPAdaptor("192.168.1.188:3030")
        led := gpio.NewLedDriver(firmataAdaptor, "2")
        pin := gpio.NewDirectPinDriver(firmataAdaptor,"2")
        button := gpio.NewButtonDriver(firmataAdaptor, "0")
        pin2toggle := pinToggle(pin)
        work := func() {
                gobot.Every(1000*time.Millisecond, func() {
                // led.Toggle()
                   pin2toggle()                         
                })
  
                      
               button.On(gpio.ButtonRelease, func(data interface{}) {
                   buttonCount++
                   fmt.Printf("button pressed %d times\n",buttonCount)
                   
                })
           
                 
                        
        }

        robot := gobot.NewRobot("bot",
                []gobot.Connection{firmataAdaptor},
                []gobot.Device{led,pin,button},
                work,
        )

        robot.Start()
}
