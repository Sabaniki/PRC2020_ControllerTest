package main

import (
	"fmt"
	"log"
	"math"

	"github.com/nobonobo/joycon"
)

func adjustment(n float32) float32 {
	if math.Abs(float64(n)) < 0.2 {
		return 0
	} else {
		return n
	}
}

func makeSign(n float32) float32 {
	if math.Abs(float64(n)) < 0 {
		return -1
	} else {
		return 1
	}
}

func main() {
	devices, err := joycon.Search()
	if err != nil {
		log.Fatalln(err)
	}
	if len(devices) == 0 {
		log.Fatalln("joycon not found")
	}
	jc, err := joycon.NewJoycon(devices[0].Path, false)
	if err != nil {
		log.Fatalln(err)
	}
	for true {
		s := <-jc.State()
		x := adjustment(s.RightAdj.X)
		y := adjustment(s.RightAdj.Y)
		fmt.Printf("L: %#v, ", 100 * (float32(math.Pow(float64(y), 2)) * makeSign(y) + float32(math.Pow(float64(x), 2)) * makeSign(x)))
		fmt.Printf("R: %#v\n", 100 * (float32(math.Pow(float64(y), 2)) * makeSign(y) - float32(math.Pow(float64(x), 2)) * makeSign(x)))
		fmt.Printf("%#v\n", s.Buttons)  // Button bits
		//fmt.Printf("%#v\n", s.RightAdj) // Right Analog Stick State
	}
	//a := <-jc.Sensor()
	//fmt.Printf("%#v\n", a.Accel) // Acceleration Sensor State
	//fmt.Printf("%#v\n", a.Gyro)  // Gyro Sensor State

	jc.Close()
}
