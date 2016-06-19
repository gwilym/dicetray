package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/getlantern/systray"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	systray.Run(onReady)
}

func onReady() {
	go roller(4, systray.AddMenuItem("4", "").ClickedCh)
	go roller(6, systray.AddMenuItem("6", "").ClickedCh)
	go roller(8, systray.AddMenuItem("8", "").ClickedCh)
	go roller(10, systray.AddMenuItem("10", "").ClickedCh)
	go roller(12, systray.AddMenuItem("12", "").ClickedCh)
	go roller(20, systray.AddMenuItem("20", "").ClickedCh)
	go roller(100, systray.AddMenuItem("100", "").ClickedCh)
	go quitter(systray.AddMenuItem("Quit", "").ClickedCh)

	systray.SetTitle("DT")
}

func roller(sides int, ch <-chan interface{}) {
	for range ch {
		systray.SetTitle(fmt.Sprintf("%d", roll(sides)))
	}
}

func roll(sides int) int {
	return rand.Intn(sides-1) + 1
}

func quitter(ch <-chan interface{}) {
	<-ch
	systray.Quit()
	os.Exit(0)
}
