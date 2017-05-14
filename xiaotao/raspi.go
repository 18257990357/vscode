package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/load"
)

const (
	BLUE           = 20
	ORANGE         = 21
	CORE_TEMP_PATH = "/sys/class/thermal/thermal_zone0/temp"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {

	fmt.Printf("System initial...")
	if rpio.Open() == nil {
		fmt.Println("[OK]")
	} else {
		fmt.Println("[ERROR]")
	}
	defer rpio.Close()
	orange := rpio.Pin(ORANGE)
	blue := rpio.Pin(BLUE)
	orange.Output()
	blue.Output()
	orange.Low()
	blue.High()

	for {
		stat, err := load.Avg()
		if err != nil {
			fmt.Println(err)
			break
		}
		interval := int(stat.Load1)
		if stat.Load1 < 1 {
			interval = 1
		}
		fmt.Printf("Load1:%.2f Temp:%.2f'C", stat.Load1, loadTemp())
		time.Sleep(time.Millisecond * time.Duration(interval*900))
		blue.Toggle()
		orange.Toggle()
		fmt.Printf("\r")
	}
}

func loadTemp() float64 {
	b, err := ioutil.ReadFile(CORE_TEMP_PATH)
	if err != nil {
		return -1000
	}
	raw, err := strconv.ParseFloat(string(b[:len(b)-2]), 64)
	if err != nil {
		fmt.Println(err)
		return -1001
	}
	return raw / 100
}
