package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func reader(inputChan chan byte, timeChan chan int) {
	event_queue := make(chan termbox.Event)
	go func() { // 在其他goroutine中开始监听
		for {
			event_queue <- termbox.PollEvent() //开始监听键盘事件
		}
	}()
	ev := <-event_queue

	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowUp:
			fmt.Println("↑")
		case termbox.KeyArrowDown:
			fmt.Println("↓")
		case termbox.KeyArrowLeft:
			fmt.Println("←")
		case termbox.KeyArrowRight:
			fmt.Println("→")
		default:
			fmt.Println("other")
		}
	}
}

func writer(inputChan chan byte, timeChan chan int) {
	//for {
	//	select {
	//	case input := <-inputChan:
	//		fmt.Println("got:", input, "time", <-timeChan)
	//	}
	//}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

}
